package data

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // add this
)

func CreateConn() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't read .env: " + err.Error())
	}

	config := struct {
		username string
		password string
		address  string
		database string
		table    string
	}{
		username: os.Getenv("username"),
		password: os.Getenv("password"),
		address:  os.Getenv("address"),
		database: os.Getenv("database"),
		table:    os.Getenv("table"),
	}

	connStr := "postgresql://" + config.username + ":" + config.password + "@" + config.address + "/" + config.database + "?sslmode=disable"

	// Connect to database
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Couldn't open database: " + err.Error())
	}

	return db
}
