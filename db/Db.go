package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
	"time"
)

var (
	user     = os.Getenv("userdb")
	password = os.Getenv("passworddb")
	ip       = os.Getenv("ipdb")
	port     = os.Getenv("portdb")
)

var (
	once  sync.Once
	db    *gorm.DB
	dbErr error
)

func ConnDB() (*gorm.DB, error) {
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

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ip, user, password, "api-rest-orm", port)
		db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if dbErr != nil {
			log.Fatal(dbErr)
		}

		log.Println("Connected!")

		configPool()

	})

	return db, dbErr
}

func ConfigShema[T interface{}](class T) {
	// Migrate the schema
	db.AutoMigrate(class)
}

func configPool() {
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Error configurando el pool de conexiones..... %v", err)
	}

}
