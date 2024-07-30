package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func MysqlConn() (*sql.DB, error) {
	dsn := openConn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	log.Println("MySQL connection established")
	return db, nil
}

func openConn() string {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", USER, PASS, HOST, PORT, DBNAME)
}
