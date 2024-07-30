package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-rod/rod"

	doctor "github.com/devpablocristo/golang/hex-arch/backend/internal/doctors/domain"

	patient "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/domain"
	ginhandler "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/infrastructure/handlers/gin"
	memkvsrepo "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/infrastructure/repositories/kvs"
	gorodservice "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/infrastructure/scrappers/go-rod"
	patientservice "github.com/devpablocristo/golang/hex-arch/backend/internal/patients/service"
	person "github.com/devpablocristo/golang/hex-arch/backend/internal/persons/domain"
)

func main() {

	patPerson := person.Person{
		UUID:     "1",
		Name:     "Homero",
		Lastname: "Simpson",
		DNI:      12345,
		Gender:   "m",
	}

	docPerson := person.Person{
		UUID:     "2",
		Name:     "Nick",
		Lastname: "Riviera",
		DNI:      63435,
		Gender:   "m",
	}

	doctor := doctor.Doctor{
		Doctor:     docPerson,
		Speciality: "Surgery",
	}

	patient := &patient.Patient{
		Patient:   patPerson,
		Doctor:    doctor,
		Hospital:  "General",
		Diagnosis: "Cancer",
	}

	b, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	kvs := make(map[string][]byte)

	patientRepository := memkvsrepo.NewMemKVS(kvs)
	patientService := patientservice.NewPatientService(patientRepository)
	goRodService := gorodservice.NewGoRodService(browser)
	patientHandler := ginhandler.NewGinHandler(patientService, goRodService)

	runPatientService(patientHandler)

}
