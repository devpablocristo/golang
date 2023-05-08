package api

import (
	"github.com/go-chi/chi"

	patients "github.com/devpablocristo/patients-api/patients/application"
)

func routes(services *patients.PatientHTTPService) *chi.Mux {
	r := chi.NewMux()

	r.Get("/patients", services.GetPatientsHandler)
	r.Post("/patients", services.CreatePatientsHandler)
	r.Get("/patients/{patientID}", services.GetPatientsByIDHandler)

	return r
}
