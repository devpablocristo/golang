package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// user struct
type User struct {
	ID       int64
	userName string
}

// database connection
var db *sql.DB

// init connection with the database
func initDB() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(localhost:3306)/test", "root", "root")
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	//Check that the database is available and accessible
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection OK")
}

// close de database connection
func shutdownDB() {
	if db != nil {
		db.Close()
	}
	fmt.Println("Database closed OK")
}

// list of all user
func list() []User {
	// query
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	var users []User
	for rows.Next() {
		var rId int64
		var rUserName string
		err = rows.Scan(&rId, &rUserName)
		u := User{
			ID:       rId,
			userName: rUserName,
		}
		users = append(users, u)
	}
	return users
}

// inserr  new user
func insert(user User) User {
	// insert
	stmt, err := db.Prepare("INSERT user SET username=?")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(user.userName)

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.ID = id
	return user
}

// update one user
func update(user User) {
	stmt, err := db.Prepare("UPDATE user SET username=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.userName, user.ID)

	rowCnt, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Rows affected", rowCnt)
	}

	if err != nil {
		panic(err)
	}
}

func main() {
	initDB()

	// insert list
	fmt.Println("Inserting users")
	userG := insert(User{userName: "gustavo"})
	userS := insert(User{userName: "samuel"})
	userJ := insert(User{userName: "jhon"})

	// list list in DB
	fmt.Println("List of all list in DB")
	userList := list()
	for _, u := range userList {
		fmt.Println(u)
	}

	// update
	userG.userName = "gus"
	userS.userName = "samu"
	userJ.userName = "jhony"

	fmt.Println("Updating list")
	update(userG)
	update(userS)
	update(userJ)

	fmt.Println("List of all list updated")
	usersUpdated := list()
	for _, u := range usersUpdated {
		fmt.Println(u)
	}

	shutdownDB()
}
