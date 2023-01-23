package db

import (
	"database/sql"
)

//Store provide all function to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

//NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
