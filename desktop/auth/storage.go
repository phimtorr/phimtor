package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"

	"github.com/adrg/xdg"
)

type FileStorage struct {
	configFilePath string
	mu             sync.RWMutex
	credentials    Credentials
}

func NewFileStorage(appName string) *FileStorage {
	if appName == "" {
		panic("app name is empty")
	}
	configFilePath, err := xdg.ConfigFile(filepath.Join(appName, "auth.yaml"))
	if err != nil {
		panic(fmt.Errorf("get config file path: %w", err))
	}

	credentials, err := loadCredentialsFromFile(configFilePath)
	if err != nil {
		panic(fmt.Errorf("load auth file: %w", err))
	}

	return &FileStorage{
		configFilePath: configFilePath,
		credentials:    credentials,
	}
}

func loadCredentialsFromFile(configFilePath string) (Credentials, error) {
	var creds Credentials
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return creds, nil
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return creds, fmt.Errorf("read file: %w", err)
	}

	if err := yaml.Unmarshal(data, &creds); err != nil {
		return creds, fmt.Errorf("unmarshal: %w", err)
	}

	return creds, nil
}

func (s *FileStorage) GetCredentials() Credentials {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.credentials
}

func (s *FileStorage) SetCredentials(creds Credentials) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.credentials = creds
	return saveCredentialsToFile(s.configFilePath, creds)
}

func saveCredentialsToFile(filePath string, creds Credentials) error {
	data, err := yaml.Marshal(creds)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
