package service

import (
	"fmt"
	"testing"
	"time"

	user "github.com/devpablocristo/interviews/bookstore-gin-rest-api/models/users"
	userRepository "github.com/devpablocristo/interviews/bookstore-gin-rest-api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	En VS no sale t.log por la salida estandar
	Correr en en la consola:

	go test -run TestCreateUser

	go test -v user_service_test.go

	Para ver los diferentes tests.
*/

func TestCreateUser(t *testing.T) {

	oid := primitive.NewObjectID()
	l, _ := time.LoadLocation("America/Buenos_Aires") // or the name of your time zone
	tm := time.Now()

	// si no se cambia algo del los datos, no los guarda,
	// no se pq., supungo que tiene que ver con algo sobre seguridad
	user := user.User{
		Id:        oid,
		Username:  "Barnie3333222.Gomadaez",
		Password:  "12345",
		Email:     "barnie@fox.com",
		CreatedAt: tm.In(l),
		UpdatedAt: tm.In(l),
	}

	u, rErr := CreateUser(user)
	if rErr != nil {
		t.Error(rErr.Message)
		t.Error(u)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario creado correctamente:\n", u, "\n\n")
	}
}

func TestGetUsers(t *testing.T) {
	urs, rErr := GetUsers()
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	}

	if len(*urs) < 1 {
		t.Error("No hay documentos en la base de datos.")
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Se obtuvieron todos los documentos correctamente:")
		for _, v := range *urs {
			fmt.Println(v)
		}
		fmt.Println()
	}
}

func TestGetUser(t *testing.T) {
	uId := "60aac57c934f172f083be39a"
	u, rErr := GetUser(uId)
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario obternido correctamente:\n", u, "\n\n")
	}
}

func TestUpdateUser(t *testing.T) {

	lastDocument, rErr := userRepository.GetIdLastInseted()
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	}

	uId := lastDocument["_id"].(primitive.ObjectID).Hex()

	l, _ := time.LoadLocation("America/Buenos_Aires") // or the name of your time zone
	tm := time.Now()

	u := user.User{
		Username:  "Abuelo888.simpson",
		Password:  "5432124234234242",
		Email:     "b.simpson@fox.com",
		UpdatedAt: tm.In(l),
	}

	ur, rErr := UpdateUser(u, uId)

	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario actualizado correctamente:\n", ur, "\n\n")
	}
}

func TestDeleteUser(t *testing.T) {

	lastDocument, rErr := userRepository.GetIdLastInseted()
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	}

	uId := lastDocument["_id"].(primitive.ObjectID).Hex()

	_, rErr = DeleteUser(uId)
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario eliminado:\n", lastDocument, "\n\n")
	}
}
