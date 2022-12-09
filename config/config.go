package config

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(c *gin.Context) (db *sql.DB) {
	config := NewConfiguration()
	config.LoadConfigurationFromFile(getFilePathConfigEnvirontment())
	dbDriver := "mysql"
	// dbHost := config.GetValue(`database.host`) //Uncomment this row if not use docker
	// dbPort := config.GetValue(`database.port`)
	dbContainer := config.GetValue(`database.db-container`)
	dbUser := config.GetValue(`database.user`)
	dbPass := config.GetValue(`database.pass`)
	// dbName := config.GetValue(`database.name`)
	// connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName) //Uncomment this row if not use docker
	connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbContainer + ")/")
	db, err := sql.Open(dbDriver, connection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Errors",
		})
	}
	Migrate(db)
	return db
}

func Migrate(db *sql.DB) {
	query := `CREATE DATABASE IF NOT EXISTS privyTest;`
	table := `CREATE TABLE IF NOT EXISTS privyTest.cakes(id int primary key auto_increment, title varchar(100),  
        description text, rating float, image text, created_at datetime, updated_at datetime)`
	use := `USE privyTest`
	_, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(use)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(context.Background(), table)
	if err != nil {
		panic(err.Error())
	}
}

func getFilePathConfigEnvirontment() string {
	env := "dev"
	switch env {
	case "dev":
		return "config-dev.json"
	case "production":
		return "config-production.json"
	}
	return "config-dev.json"
}
