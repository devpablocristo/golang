package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database struct
type Database struct {
	DB *gorm.DB
}

// NewDatabase : intializes and returns mysql db
func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}


func dbConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
}




package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

const (
	DB_HOST = "mysql-meli-items" // nombre del container en docker
	DB_PORT = 3306
	DB_NAME = "meli_items"
	DB_USER = "root"
	DB_PASS = "secret"
)

var db *sqlx.DB //nolint:gochecknoglobals

func GetConnectionDB() (*sqlx.DB, error) {
	var err error

	if db == nil {
		db, err = sqlx.Connect("mysql", dbConnectionURL())
		if err != nil {
			fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
			return nil, fmt.Errorf("### DB ERROR: %w", err)
		}
	}

	// if err := autoMigrate(db); err != nil {
	// 	return nil, err
	// }

	return db, nil
}

// func autoMigrate(db *sqlx.DB) error {
// 	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
// 	if err != nil {
// 		fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
// 		return fmt.Errorf("error instantiating migration: %w", err)
// 	}

// 	dbMigration, err := migrate.NewWithDatabaseInstance(
// 		"file://../../../db",
// 		"mysql",
// 		driver,
// 	)

// 	if err != nil {
// 		fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
// 		return fmt.Errorf("error instantiating migration: %w", err)
// 	}

// 	if err := dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
// 		return fmt.Errorf("error executing migration: %w", err)
// 	}

// 	return nil
// }

func dbConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
}
