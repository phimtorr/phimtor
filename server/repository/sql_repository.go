package repository

import (
	"database/sql"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) SQLRepository {
	if db == nil {
		panic("db is nil")
	}
	return SQLRepository{db: db}
}
