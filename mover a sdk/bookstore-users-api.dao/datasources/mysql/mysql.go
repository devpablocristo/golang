package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	UsersDB *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"myroot",
		"pass",
		"127.0.0.1:3306",
		"users",
	)
	var err error
	UsersDB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = UsersDB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database successfully configurated")
}
