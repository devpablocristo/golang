package api

import (
	patients "github.com/devpablocristo/interviews/hospital2/patients/src/application"
	"github.com/go-chi/chi"
)

func routes(services *patients.PatientHTTPService) *chi.Mux {
	r := chi.NewMux()

	r.Get("/patients", services.GetPatientsHandler)
	r.Post("/patients", services.CreatePatientsHandler)
	r.Get("/patients/{patientID}", services.GetPatientsByIDHandler)

	return r
}
