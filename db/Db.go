package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // add this
	"log"
	"os"
	"sync"
)

var (
	user     = os.Getenv("userdb")
	password = os.Getenv("passworddb")
	ip       = os.Getenv("ipdb")
	port     = os.Getenv("portdb")
)

var (
	once  sync.Once
	db    *sql.DB
	dbErr error
)

// Connect to database
func ConnDB() (*sql.DB, error) {
	once.Do(func() {
		if user == "" {
			user = "postgres"
		}
		if password == "" {
			password = "postgres"
		}
		if ip == "" {
			ip = "localhost"
		}
		if port == "" {
			port = "25432"
		}
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/golang?sslmode=disable", user, password, ip, port)
		db, dbErr = sql.Open("postgres", connStr)
		if dbErr != nil {
			log.Fatal(dbErr)
		}
		pingErr := db.Ping()
		if pingErr != nil {
			log.Fatal(pingErr)
		}
		log.Println("Connected!")
	})
	return db, dbErr
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return rows, err
}
