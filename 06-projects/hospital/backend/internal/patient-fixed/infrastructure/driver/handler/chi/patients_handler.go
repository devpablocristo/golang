package patients

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/devpablocristo/patients-api/internal/web"
	models "github.com/devpablocristo/patients-api/patients/domain"

	//patients "github.com/devpablocristo/patients-api/patients/infrastructure/http"
	"github.com/go-chi/chi"
)

type PatientHTTPService struct {
	gtw patients.PatientGateway
}

func NewPatientHTTPService(db *sql.DB) *PatientHTTPService {
	return &PatientHTTPService{
		patients.NewPatientGateway(db),
	}
}

func (s *PatientHTTPService) GetPatientsHandler(w http.ResponseWriter, r *http.Request) {
	p := s.gtw.GetPatients()
	if p == nil || len(p) == 0 {
		p = []*models.Patient{}
	}
	web.Success(&p, http.StatusOK).Send(w)
}

func (s *PatientHTTPService) GetPatientsByIDHandler(w http.ResponseWriter, r *http.Request) {
	patientID := chi.URLParam(r, "patientID")
	id, _ := strconv.ParseInt(patientID, 10, 64)
	patient, err := s.gtw.GetPatientByID(id)

	if err != nil {
		web.ErrBadRequest.Send(w)
		return
	}

	web.Success(&patient, http.StatusOK).Send(w)
}

func (s *PatientHTTPService) CreatePatientsHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	var cmd models.CreatePatientCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		web.ErrInvalidJSON.Send(w)
		return
	}

	patient, err := s.gtw.CreatePatient(&cmd)

	if err != nil {
		web.ErrBadRequest.Send(w)
		return
	}

	log.Println(patient)
	web.Success(&patient, http.StatusOK).Send(w)
}
