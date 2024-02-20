package repository

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	if db == nil {
		panic("db is nil")
	}
	return Repository{db: db}
}
