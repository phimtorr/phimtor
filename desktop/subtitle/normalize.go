package subtitle

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/asticode/go-astisub"
	"github.com/friendsofgo/errors"
	"github.com/rs/zerolog/log"
)

func GetFileFromLink(link string) (string, []byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", nil, errors.Wrap(err, "get link")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error().Err(err).Msg("close response body")
		}
	}()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, errors.Wrap(err, "read content")
	}

	tokens := strings.Split(link, "/")
	fileName := html.UnescapeString(tokens[len(tokens)-1])

	return fileName, content, nil
}

func Normalize(fileName string, content []byte, addDuration time.Duration) ([]byte, error) {
	var (
		sub *astisub.Subtitles
		err error
	)

	reader := bytes.NewReader(content)

	fileExt := path.Ext(fileName)
	switch strings.Trim(fileExt, ".") {
	case "srt":
		sub, err = astisub.ReadFromSRT(reader)
	case "vtt":
		sub, err = astisub.ReadFromWebVTT(reader)
	case "ssa", "ass":
		sub, err = astisub.ReadFromSSA(reader)
	default:
		return nil, fmt.Errorf("not supported file extension %s", fileExt)
	}

	if err != nil {
		return nil, fmt.Errorf("read subtitle: %w", err)
	}

	if addDuration != 0 {
		sub.Add(addDuration)
	}

	buf := &bytes.Buffer{}
	if err := sub.WriteToWebVTT(buf); err != nil {
		return nil, fmt.Errorf("write subtitle: %w", err)
	}

	return []byte(html.UnescapeString(buf.String())), nil
}
