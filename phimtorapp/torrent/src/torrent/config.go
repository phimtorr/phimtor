package torrent

import "errors"

type Config struct {
	DataDir         string
	DeleteAfterDrop bool
	Debug           bool
}

func (c Config) validate() error {
	if c.DataDir == "" {
		return errors.New("data dir is required")
	}
	return nil
}
