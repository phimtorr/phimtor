package repository

import "database/sql"

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	if db == nil {
		panic("db is nil")
	}
	return AdminRepository{db: db}
}
