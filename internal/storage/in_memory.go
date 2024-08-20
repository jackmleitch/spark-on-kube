package storage

import (
	"github.com/google/uuid"
	"sync"
)

type InMemoryStorage[T any] struct {
	mu    sync.RWMutex
	cache map[uuid.UUID]T
}

func NewInMemoryStorage[T any]() *InMemoryStorage[T] {
	return &InMemoryStorage[T]{
		cache: make(map[uuid.UUID]T),
	}
}

func (s *InMemoryStorage[T]) Get(key uuid.UUID) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	record, found := s.cache[key]
	return record, found
}

func (s *InMemoryStorage[T]) GetAll() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	values := make([]T, len(s.cache))
	for _, value := range s.cache {
		values = append(values, value)
	}
	return values
}

func (s *InMemoryStorage[T]) Put(key uuid.UUID, record T) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache[key] = record
	return record
}

func (s *InMemoryStorage[T]) Update(key uuid.UUID, record T) T {
	s.Put(key, record)
	return record
}

func (s *InMemoryStorage[T]) Delete(key uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.cache, key)
}
