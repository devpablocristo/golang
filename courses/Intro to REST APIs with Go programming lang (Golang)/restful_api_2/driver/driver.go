package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

/*
no entiendo pq si no uso esta funcion y llamo desde main, cuando uso db en una fucion da error
*/
func ConectarDb() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logErrores(err)

	db, err := sql.Open("postgres", pgUrl)
	logErrores(err)

	err = db.Ping()
	logErrores(err)

	return db
}

func logErrores(err error) {
	if err != nil {
		log.Fatal("ERROR!!!", err)
	}
}
