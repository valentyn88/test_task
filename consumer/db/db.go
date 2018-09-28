package db

import (
	"sync"
	"test_task/consumer/user"
)

// DB db for storing users
type DB struct {
	mu sync.Mutex
	rows map[string]user.User
}

// NewDB create new DB instance
func NewDB() DB {
	return DB{rows:make(map[string]user.User)}
}

// AddRow add new or update db row
func (db DB) AddRow(u user.User) {
	db.mu.Lock()
	db.rows[u.ID] = u
	db.mu.Unlock()
}
