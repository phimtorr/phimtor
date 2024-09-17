package repository

import (
	"database/sql"
)

type SQLRepo2 struct {
	db *sql.DB
}

func NewSQLRepo2(db *sql.DB) SQLRepo2 {
	if db == nil {
		panic("db is nil")
	}
	return SQLRepo2{db: db}
}
