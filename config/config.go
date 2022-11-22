package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@db_apps:3305/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
