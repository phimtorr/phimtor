package updater

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/fynelabs/selfupdate"
	"github.com/phimtorr/phimtor/desktop/build"

	"github.com/rs/zerolog/log"
)

type VersionFetcher interface {
	GetVersion(ctx context.Context) (string, error)
}

type Updater struct {
	currentVersion   string
	intervalDuration time.Duration
	versionFetcher   VersionFetcher

	done    chan struct{}
	stopped chan struct{}
}

func NewUpdater(currentVersion string, intervalDuration time.Duration, versionFetcher VersionFetcher) *Updater {
	if versionFetcher == nil {
		panic("versionFetcher is required")

	}
	if intervalDuration <= 0 {
		panic("intervalDuration must be positive")

	}
	if currentVersion == "" {
		panic("currentVersion is required")
	}

	return &Updater{
		currentVersion:   currentVersion,
		intervalDuration: intervalDuration,
		versionFetcher:   versionFetcher,

		done:    make(chan struct{}),
		stopped: make(chan struct{}),
	}
}

func (u *Updater) Start() {
	defer close(u.stopped)

	u.doUpdateIfNeed()

	ticker := time.NewTicker(u.intervalDuration)
	defer ticker.Stop()

	for {
		select {
		case <-u.done:
			return
		case <-ticker.C:
			u.doUpdateIfNeed()
		}
	}
}

func (u *Updater) doUpdateIfNeed() {
	ctx := context.Background()

	logger := log.Ctx(ctx).With().Str("component", "updater").Logger()

	version, err := u.versionFetcher.GetVersion(ctx)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get version")
		return
	}

	if version == u.currentVersion {
		logger.Debug().Msg("No updates available")
		return
	}

	logger.Debug().Str("currentVersion", u.currentVersion).Str("newVersion", version).Msg("Update available")
	logger.Debug().Msg("Updating...")

	url := generateUpdateURL()
	err = doUpdate(url)
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to update")
		return
	}

	// It avoids infinite loop of updates
	u.currentVersion = version

	logger.Info().Msg("Updated successfully")
}

func generateUpdateURL() string {
	goos := runtime.GOOS
	url := build.ServerAddr + "/public/desktop-binaries/" + goos + "/" + build.AppName
	if goos == "windows" {
		url += ".exe"
	}
	return url
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "get update")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		return errors.Wrap(err, "apply update")
	}
	return nil
}

func (u *Updater) Stop() error {
	close(u.done)
	<-u.stopped
	return nil
}
