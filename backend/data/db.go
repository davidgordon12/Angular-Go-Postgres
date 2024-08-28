package data

import (
	"database/sql"
	"log"
	"os"
	   _ "github.com/lib/pq" // add this
	"github.com/joho/godotenv"
)

func CreateConn() *sql.DB {
	err := godotenv.Load() 
	if err != nil {
		log.Fatal("Couldn't read .env: " + err.Error())
	}

	config := struct {
		username string
		password string
		address string
		database string
		table string
	}{
		username: os.Getenv("username"),
		password: os.Getenv("password"),
		address: os.Getenv("address"),
		database: os.Getenv("database"),
		table : os.Getenv("table"),
	}

	connStr := "postgresql://" + config.username + ":" + config.password + "@" + config.address + "/" + config.database + "?sslmode=disable"

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Couldn't open database: " + err.Error())
	}

	return db
}
