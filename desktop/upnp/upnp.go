package upnp

import (
	stdErrors "errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/phimtorr/phimtor/common/logs"
	"github.com/rs/cors"

	"github.com/phimtorr/phimtor/desktop/subtitle"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/desktop/net"
	"github.com/phimtorr/phimtor/desktop/torrent"
	"github.com/rs/zerolog/log"
)

type UPnP struct {
	torManager *torrent.Manager
	listenPort int
	closeFns   []func() error

	mu    sync.RWMutex
	state State
}

func NewUPnP(torManager *torrent.Manager) *UPnP {
	if torManager == nil {
		panic("torManager is nil")
	}
	return &UPnP{
		torManager: torManager,
	}
}

func (u *UPnP) Run() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(logs.NewHTTPStructuredLogger(log.Logger))
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.NoCache)

	r.Get("/torrents/{infoHash}/{fileIndex}/{fileName}", func(w http.ResponseWriter, r *http.Request) {
		infoHash, err := parseInfoHash(chi.URLParam(r, "infoHash"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fileIndex, err := parseFileIndex(chi.URLParam(r, "fileIndex"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := u.torManager.StreamVideoFile(w, r, infoHash, fileIndex); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/subtitles/{fileName}", func(w http.ResponseWriter, r *http.Request) {
		fileName, err := url.QueryUnescape(chi.URLParam(r, "fileName"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u.mu.RLock()
		defer u.mu.RUnlock()

		if u.state.SubtitleFileName != fileName {
			http.Error(w, "subtitle not found", http.StatusNotFound)
			return
		}

		content, err := subtitle.NormalizeToSRT(u.state.SubtitleFileName, u.state.SubtitleContent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/srt")
		_, _ = w.Write(content)
	})

	ln, listenPort, cleanUp, err := net.CreateListener()
	if err != nil {
		return errors.Wrap(err, "create listener")
	}
	u.closeFns = append(u.closeFns, cleanUp)
	u.listenPort = listenPort

	server := &http.Server{
		Handler: r,
	}
	u.closeFns = append(u.closeFns, server.Close)

	go func() {
		log.Info().Int("port", u.listenPort).Msg("Starting UPnP HTTP server")
		if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start UPnP HTTP server")
		}
	}()

	return nil
}

func (u *UPnP) Close() error {
	slices.Reverse(u.closeFns)
	var errs []error
	for _, fn := range u.closeFns {
		if err := fn(); err != nil {
			errs = append(errs, err)
		}
	}
	return stdErrors.Join(errs...)
}

var (
	ErrInvalidInfoHash = commonErrors.NewIncorrectInputError("invalid-info-hash", "invalid info hash")
)

func parseInfoHash(infoHashRaw string) (torrent.InfoHash, error) {
	infoHash, err := torrent.InfoHashFromString(infoHashRaw)
	if err != nil {
		return torrent.InfoHash{}, errors.Wrap(ErrInvalidInfoHash, fmt.Sprintf("parse info_hash=%s, err=%v", infoHashRaw, err))
	}
	return infoHash, nil
}

var (
	ErrInvalidFileIndex = commonErrors.NewIncorrectInputError("invalid-file-index", "invalid file index")
)

func parseFileIndex(fileIndexRaw string) (int, error) {
	fileIndex, err := strconv.Atoi(fileIndexRaw)
	if err != nil {
		return 0, errors.Wrap(ErrInvalidFileIndex, fmt.Sprintf("parse file_index=%s, err=%v", fileIndexRaw, err))
	}
	return fileIndex, nil
}
