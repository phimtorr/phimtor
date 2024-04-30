package upnp

import (
	"bytes"
	"embed"
	"fmt"
	"html"
	"strings"
	"text/template"

	"github.com/friendsofgo/errors"
)

const (
	// dlnaOrgFlagSenderPaced = 1 << 31
	// dlnaOrgFlagTimeBasedSeek = 1 << 30
	// dlnaOrgFlagByteBasedSeek = 1 << 29
	// dlnaOrgFlagPlayContainer = 1 << 28
	// dlnaOrgFlagS0Increase = 1 << 27
	// dlnaOrgFlagSnIncrease = 1 << 26
	// dlnaOrgFlagRtspPause = 1 << 25
	dlnaOrgFlagStreamingTransferMode = 1 << 24
	// dlnaOrgFlagInteractiveTransfertMode = 1 << 23
	dlnaOrgFlagBackgroundTransfertMode = 1 << 22
	dlnaOrgFlagConnectionStall         = 1 << 21
	dlnaOrgFlagDlnaV15                 = 1 << 20
)

//go:embed didl-lite.xml
var didlTemplate embed.FS

var tmpl *template.Template

type templateData struct {
	Title           string
	MimeType        string
	ContentFeatures string
	VideoURL        string
	SubURL          string
}

func init() {
	tmpl = template.Must(template.ParseFS(didlTemplate, "didl-lite.xml"))
}

var (
	dlnaprofiles = map[string]string{
		"video/x-mkv":             "DLNA.ORG_PN=MATROSKA",
		"video/x-matroska":        "DLNA.ORG_PN=MATROSKA",
		"video/x-msvideo":         "DLNA.ORG_PN=AVI",
		"video/mpeg":              "DLNA.ORG_PN=MPEG1",
		"video/vnd.dlna.mpeg-tts": "DLNA.ORG_PN=MPEG1",
		"video/mp4":               "DLNA.ORG_PN=AVC_MP4_MP_SD_AAC_MULT5",
		"video/quicktime":         "DLNA.ORG_PN=AVC_MP4_MP_SD_AAC_MULT5",
		"video/x-m4v":             "DLNA.ORG_PN=AVC_MP4_MP_SD_AAC_MULT5",
		"video/3gpp":              "DLNA.ORG_PN=AVC_MP4_MP_SD_AAC_MULT5",
		"video/x-flv":             "DLNA.ORG_PN=AVC_MP4_MP_SD_AAC_MULT5",
		"video/x-ms-wmv":          "DLNA.ORG_PN=WMVHIGH_FULL",
		"audio/mpeg":              "DLNA.ORG_PN=MP3",
		"image/jpeg":              "JPEG_LRG",
		"image/png":               "PNG_LRG",
	}
)

func defaultStreamingFlags() string {
	return fmt.Sprintf("%.8x%.24x", dlnaOrgFlagStreamingTransferMode|
		dlnaOrgFlagBackgroundTransfertMode|
		dlnaOrgFlagConnectionStall|
		dlnaOrgFlagDlnaV15, 0)
}

func buildContentFeatures(mimeType string) string {
	var cf strings.Builder

	if profile, ok := dlnaprofiles[mimeType]; ok {
		cf.WriteString(profile)
	} else {
		cf.WriteString("DLNA.ORG_PN=UNKNOWN")
	}
	cf.WriteString(";DLNA.ORG_OP=01;DLNA.ORG_CI=0;")

	cf.WriteString("DLNA.ORG_FLAGS=")
	cf.WriteString(defaultStreamingFlags())

	return cf.String()
}

func buildCurrentURIMetaData(videoFileName, mineType, videoURL, subURL string) (string, error) {
	contentFeatures := buildContentFeatures(mineType)

	var b bytes.Buffer
	if err := tmpl.Execute(&b, templateData{
		Title:           videoFileName,
		MimeType:        mineType,
		ContentFeatures: contentFeatures,
		VideoURL:        videoURL,
		SubURL:          subURL,
	}); err != nil {
		return "", errors.Wrap(err, "execute template")
	}

	return strings.TrimSpace(html.EscapeString(b.String())), nil
}
