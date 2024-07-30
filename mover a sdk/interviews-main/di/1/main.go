package main

import "fmt"

/////////////////////////////////////
// Component 1: db
/////////////////////////////////////

type dbHandler interface {
	getUser() string
	getAllUsers() []string
}

type db struct {
	handler dbHandler
}

func newDBHandler(handler dbHandler) *db {
	return &db{handler: handler}
}
func (db db) getUser() string {
	return "Tori"
}

func (db db) getAllUsers() []string {
	return []string{}
}

/////////////////////////////////////

/////////////////////////////////////
// Component 2: greeter
/////////////////////////////////////

type greeterHandler interface {
	helloGreet(userName string)
}

type greeter struct {
	handler greeterHandler
}

func newGreeterHandler(handler greeterHandler) *greeter {
	return &greeter{handler: handler}
}

func (gh greeter) helloGreet(userName string) {
	fmt.Println("Hello,", userName)
}

/////////////////////////////////////

type program struct {
	dbProgram db      // interface NOT struct!!!
	grProgram greeter // interface NOT struct!!!
}

func (p program) execute() {
	user := p.dbProgram.getUser()
	p.grProgram.helloGreet(user)
}

func main() {

	var dbVar db      // struct NOT interface!!!
	var grVar greeter // struct NOT interface!!!

	dbVar2 := db{}      // struct NOT interface!!!
	grVar2 := greeter{} // struct NOT interface!!!

	//x := db{}

	p1 := program{
		dbProgram: dbVar,
		grProgram: grVar,
	}

	p2 := program{
		dbProgram: dbVar2,
		grProgram: grVar2,
	}

	p1.execute()
	p2.execute()

}
