package repository

import (
	"database/sql"
)

type SQLRepository struct {
	db          *sql.DB
	ytsTrackers []string
}

func NewSQLRepository(db *sql.DB, ytsTrackers []string) SQLRepository {
	if db == nil {
		panic("db is nil")
	}
	if len(ytsTrackers) == 0 {
		panic("ytsTrackers is empty")
	}
	return SQLRepository{
		db:          db,
		ytsTrackers: ytsTrackers,
	}
}
