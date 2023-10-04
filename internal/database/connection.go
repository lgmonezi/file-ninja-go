package database

import (
	"database/sql"
	"github.com/lgmonezi/file-ninja-go/config"
	_ "modernc.org/sqlite"
)

var dbConn *sql.DB

func init() {
	var err error
	dbConn, err = getConnection()
	if err != nil {
		panic(err)
	}
}
func getConnection() (*sql.DB, error) {
	return sql.Open(config.DatabaseDriver, config.DatabaseUrl)
}
