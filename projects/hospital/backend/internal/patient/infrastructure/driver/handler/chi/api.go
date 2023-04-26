package api

import (
	// "github.com/devpablocristo/interviews/hospital2/patients/internal/storage"
	// patients "github.com/devpablocristo/interviews/hospital2/patients/src/application"

	postgres "github.com/devpablocristo/patients-api/patients/infrastructure/driven/repository/postgres"
)

func Start(port string) {
	db := postgres.ConnectToDB()
	defer db.Close()

	r := routes(postgres.NewPatientHTTPService(db))
	server := newServer(port, r)

	server.Start()
}
