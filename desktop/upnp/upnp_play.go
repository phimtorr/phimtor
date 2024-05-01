package upnp

import (
	"context"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/huin/goupnp/soap"

	"github.com/rs/zerolog/log"

	"github.com/friendsofgo/errors"
	"github.com/gabriel-vasile/mimetype"
	"github.com/huin/goupnp/dcps/av1"
	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/torrent"
)

func (u *UPnP) Play(ctx context.Context, torrentLink api.TorrentLink, subFileName string, subContent []byte) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	client, err := u.state.GetSelectedClient()
	if err != nil {
		return errors.Wrap(err, "get selected client")
	}

	infoHash, err := u.torManager.AddFromLink(torrentLink.Link)
	if err != nil {
		return errors.Wrap(err, "add torrent from link")
	}

	// set state
	u.state.InfoHash = infoHash
	u.state.FileIndex = torrentLink.FileIndex

	videoFile, err := u.torManager.GetFile(infoHash, torrentLink.FileIndex)
	if err != nil {
		return errors.Wrap(err, "get video file")
	}
	videoFileName := filepath.Base(videoFile.DisplayPath())
	videoMimeType, err := mimetype.DetectReader(videoFile.NewReader())
	if err != nil {
		return errors.Wrap(err, "detect video mime type")
	}

	videoURL := u.buildVideoURL(client, infoHash, torrentLink.FileIndex, videoFileName)

	var subURL string
	if subFileName != "" && len(subContent) > 0 {
		if err := u.state.SetSubtitle(subFileName, subContent); err != nil {
			return errors.Wrap(err, "set subtitle")
		}
		subURL = u.buildSubtitleURL(client, u.state.SubtitleFileName)
	}

	if err := client.StopCtx(ctx, 0); err != nil {
		log.Ctx(ctx).Warn().Err(err).Msg("Stop failed")
	}

	currentURIMetaData, err := buildCurrentURIMetaData(videoFileName, videoMimeType.String(), videoURL, subURL)
	if err != nil {
		return errors.Wrap(err, "build current uri meta data")
	}

	err = client.SetAVTransportURICtx(ctx, 0, videoURL, currentURIMetaData)
	if err != nil {
		return errors.Wrap(err, "set av transport uri")
	}

	err = client.PlayCtx(ctx, 0, "1")
	var soapError *soap.SOAPFaultError
	if errors.As(err, &soapError) {
		// When call play, the TV still play but return Action Failed error. I don't know why.
		// So, I just ignore this error.
		if soapError.Detail.UPnPError.ErrorDescription == "Action Failed" {
			return nil
		}
	}

	if err != nil {
		return errors.Wrap(err, "play")
	}

	return nil
}

func (u *UPnP) buildVideoURL(client *av1.AVTransport1, hash torrent.InfoHash, fileIndex int, fileName string) string {
	fileName = url.QueryEscape(fileName)
	return fmt.Sprintf("http://%s/torrents/%s/%d/%s", u.listenAddress(client), hash, fileIndex, fileName)
}

func (u *UPnP) buildSubtitleURL(client *av1.AVTransport1, fileName string) string {
	fileName = url.QueryEscape(fileName)
	return fmt.Sprintf("http://%s/subtitles/%s", u.listenAddress(client), fileName)
}

func (u *UPnP) listenAddress(client *av1.AVTransport1) string {
	return fmt.Sprintf("%s:%d", client.LocalAddr().String(), u.listenPort)
}
