package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	sqldblogger "github.com/simukti/sqldb-logger"
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
		MultiStatements:      true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(fmt.Sprintf("Error create database connection: %v", err))
	}
	db = sqldblogger.OpenDriver(config.FormatDSN(), db.Driver(), newSQLLogAdapter())
	// TODO: move to config
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Second)
	return db
}
