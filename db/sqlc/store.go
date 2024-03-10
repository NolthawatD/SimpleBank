package db

import "database/sql"

// Store provides all function to execute db queries and transactions
// *queries have already by sqlc in db.go
// extend * Queries to composition (like inheritance)
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
