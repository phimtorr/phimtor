package torrent

import (
	"errors"
	"fmt"

	"github.com/anacrolix/torrent"
)

type Manager struct {
	client *torrent.Client
}

func NewManager(dataDir string) *Manager {
	if dataDir == "" {
		panic("data dir is required")
	}

	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = dataDir
	cfg.Debug = false

	client, err := torrent.NewClient(cfg)
	if err != nil {
		panic(fmt.Sprintf("create torrent client: %v", err))
	}

	return &Manager{
		client: client,
	}
}

func (m *Manager) Close() error {
	if m.client != nil {
		errs := m.client.Close()
		return errors.Join(errs...)
	}
	return nil
}
