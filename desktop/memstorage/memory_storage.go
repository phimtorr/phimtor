package memstorage

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Storage struct {
	maxSize int64

	mu          sync.RWMutex
	data        map[uuid.UUID]Item
	itemIDs     []uuid.UUID
	currentSize int64
}

type Item struct {
	FileName     string
	Content      []byte
	DateModified time.Time
}

// New create new storage.
func New(maxSize int64) *Storage {
	return &Storage{
		maxSize: maxSize,
		data:    make(map[uuid.UUID]Item),
		itemIDs: make([]uuid.UUID, 0),
	}
}

func (s *Storage) Save(fileName string, content []byte) (uuid.UUID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fileSize := int64(len(content))
	for s.currentSize+fileSize > s.maxSize && s.currentSize > 0 {
		s.removeLastItem()
	}

	id := uuid.New()
	s.data[id] = Item{
		FileName:     fileName,
		Content:      content,
		DateModified: time.Now(),
	}
	s.itemIDs = append(s.itemIDs, id)
	s.currentSize += fileSize

	return id, nil
}

func (s *Storage) removeLastItem() {
	lastID := s.itemIDs[0]
	s.delete(lastID)
}

func (s *Storage) Delete(id uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.delete(id)
}

func (s *Storage) delete(id uuid.UUID) {
	item, ok := s.data[id]
	if !ok {
		return
	}

	delete(s.data, id)
	for i, itemID := range s.itemIDs {
		if itemID == id {
			s.itemIDs = append(s.itemIDs[:i], s.itemIDs[i+1:]...)
			break
		}
	}
	s.currentSize -= int64(len(item.Content))
}

func (s *Storage) Get(id uuid.UUID) (Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.data[id]
	return item, ok
}
