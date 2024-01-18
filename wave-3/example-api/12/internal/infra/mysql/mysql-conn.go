package mysql

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

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
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
}
