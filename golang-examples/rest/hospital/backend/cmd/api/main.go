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

// package main

// import (
// 	"log"

// 	"github.com/mercadolibre/fury_go-platform/pkg/fury"
// 	"github.com/osalomon89/test-crud-api/backend/internal/patients/services"
// 	"github.com/osalomon89/test-crud-api/internal/infrastructure/repositories/mysql"
// 	server "github.com/osalomon89/test-crud-api/internal/infrastructure/server"
// 	"github.com/osalomon89/test-crud-api/internal/infrastructure/server/handler"
// )

// func main() {
// 	if err := run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func run() error {
// 	app, err := fury.NewWebApplication()
// 	if err != nil {
// 		return err
// 	}

// 	furyHandler := server.NewHTTPServer(app, newHandlers())
// 	furyHandler.SetupRouter()

// 	return furyHandler.Run()
// }

// func newHandlers() handler.ItemHandler {
// 	conn, err := mysql.GetConnectionDB()
// 	if err != nil {
// 		panic("error connecting to DB: " + err.Error())
// 	}

// 	itemRepository, err := mysql.NewItemRepository(conn)
// 	if err != nil {
// 		panic("error creating item repository: " + err.Error())
// 	}

// 	itemService, err := services.NewItemService(itemRepository)
// 	if err != nil {
// 		panic("error creating item service: " + err.Error())
// 	}

// 	itemHandler, err := handler.NewItemHandler(itemService)
// 	if err != nil {
// 		panic("error creating item handler: " + err.Error())
// 	}

// 	return itemHandler
// }

// package main

// import (
// 	"sync"

// 	application "github.com/devpablocristo/go-concepts/hex-arch/persons/application"
// 	domain "github.com/devpablocristo/go-concepts/hex-arch/persons/domain"
// 	inmemorydb "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driven/repository/inmemory"
// 	ginAdapter "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driving/http/gin"
// 	goriAdapter "github.com/devpablocristo/go-concepts/hex-arch/persons/infrastructure/driving/http/gorilla-mux"
// )

// func main() {

// 	wg := sync.WaitGroup{}
// 	wg.Add(2)

// 	storage := inmemorydb.NewInmemoryDB(make(map[string]domain.Person))
// 	//mysql := mysql.NewPersonMySQL(nil)
// 	personService := application.NewPersonaApplication(storage)

// 	muxHandler := goriAdapter.NewGorillaHandler(personService)
// 	muxHandler.SetupRoutes()
// 	go func() {
// 		muxHandler.RunServer("default")
// 		wg.Done()
// 	}()

// 	ginHandler := ginAdapter.NewGinHandler(personService)
// 	ginHandler.SetupRoutes()
// 	go func() {
// 		ginHandler.Run("default")
// 		wg.Done()
// 	}()

// 	wg.Wait()

// }
