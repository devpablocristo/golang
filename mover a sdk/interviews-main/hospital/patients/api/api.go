package api

import (
	"github.com/devpablocristo/interviews/hospital2/patients/internal/storage"
	patients "github.com/devpablocristo/interviews/hospital2/patients/src/application"
)

func Start(port string) {
	db := storage.ConnectToDB()
	defer db.Close()

	r := routes(patients.NewPatientHTTPService(db))
	server := newServer(port, r)

	server.Start()
}
