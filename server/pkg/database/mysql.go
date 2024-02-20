package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	defaultDialTimeout = 10 * time.Second
)

func NewMySqlDB() *sql.DB {
	config := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		Loc:                  time.Now().Location(),
		Timeout:              defaultDialTimeout,
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(fmt.Sprintf("Error create database connection: %v", err))
	}
	return db
}
