package config

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"log"

)

func InitDB(c *gin.Context) (db *sql.DB) {
	// Get environment setting for database
	dbDriver := "mysql"
	dbHost := goDotEnvVariable("MYSQL_HOST")
	dbPort := goDotEnvVariable("MYSQL_PORT")
	dbUser := goDotEnvVariable("MYSQL_USER")
	dbPass := goDotEnvVariable("MYSQL_PASSWORD")
	// dbName := goDotEnvVariable("MYSQL_DBNAME")

	// Create connection database
	connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" +"?parseTime=true")
	db, err := sql.Open(dbDriver, connection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Errors",
		})
	}
	// Migrate to create database name and tables
	Migrate(db)
	return db
}

// Function migrate
func Migrate(db *sql.DB) {
	query := `CREATE SCHEMA IF NOT EXISTS todo4`
	table1 := `CREATE TABLE IF NOT EXISTS todo4.activities(id int primary key auto_increment, email varchar(100), title varchar(100),  
        created_at datetime, updated_at datetime)`
	table2 := `CREATE TABLE IF NOT EXISTS todo4.todos(id int primary key auto_increment, activity_group_id int , title varchar(100), 
		is_active boolean, priority varchar(100) DEFAULT "very-high", created_at datetime, updated_at datetime)`
	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.ExecContext(context.Background(), table1)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.ExecContext(context.Background(), table2)
	if err != nil {
		panic(err.Error())
	}
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }