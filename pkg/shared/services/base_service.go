package services

import (
    "errors"
    "sync"
)

// BaseService реализует общую логику для работы с сущностями в памяти.
type BaseService[T any] struct {
    items map[string]T
    mu    sync.RWMutex
}

// NewBaseService создает новый BaseService.
func NewBaseService[T any]() *BaseService[T] {
    return &BaseService[T]{
        items: make(map[string]T),
    }
}

func (s *BaseService[T]) Create(id string, item T) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.items[id]; exists {
        return errors.New("item already exists")
    }
    s.items[id] = item
    return nil
}

func (s *BaseService[T]) Get(id string) (T, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    item, exists := s.items[id]
    if !exists {
        return *new(T), errors.New("item not found")
    }
    return item, nil
}

func (s *BaseService[T]) Update(id string, item T) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.items[id]; !exists {
        return errors.New("item not found")
    }
    s.items[id] = item
    return nil
}

func (s *BaseService[T]) Delete(id string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.items[id]; !exists {
        return errors.New("item not found")
    }
    delete(s.items, id)
    return nil
}
