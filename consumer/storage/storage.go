package storage

import (
	"sync"
	"test_task/user"
)

// Storager an interface for storage
type Storager interface {
	AddRow(u user.User) user.User
}

// MapStorage map for storing users
type MapStorage struct {
	mu   sync.Mutex
	rows map[int64]user.User
}

// New create new MapStorage
func New() MapStorage {
	return MapStorage{rows: make(map[int64]user.User)}
}

// AddRow add new or update exists row
func (m MapStorage) AddRow(u user.User) user.User {
	m.mu.Lock()
	m.rows[u.Id] = u
	m.mu.Unlock()

	return m.rows[u.Id]
}
