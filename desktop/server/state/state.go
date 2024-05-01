package state

import "sync"

type State struct {
	mu   sync.RWMutex
	data map[string]any
}

func New() *State {
	return &State{
		data: make(map[string]any),
	}
}

func (s *State) Set(key string, value any) {
	s.mu.Lock()
	s.mu.Unlock()

	s.data[key] = value
}

func (s *State) Delete(key string) {
	s.mu.Lock()
	s.mu.Unlock()

	delete(s.data, key)
}

func (s *State) Get(key string) (any, bool) {
	s.mu.RLock()
	s.mu.RUnlock()

	v, ok := s.data[key]
	return v, ok
}
