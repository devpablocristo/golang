package service

import (
	"context"
	"log"

	domain "crudmongodb/internal/domain"
	port "crudmongodb/internal/service/ports"
)

type Service struct {
	repo port.Repo
	ctx  context.Context
}

func NewService(r port.Repo, c context.Context) port.Service {
	return &Service{
		repo: r,
		ctx:  c,
	}
}

func (srv *Service) Create(newDocument domain.Listing) error {
	err := srv.repo.Create(newDocument)
	if err != nil {
		log.Fatal("Error al crear el documento:", err)
	}
	log.Println("Documento creado con éxito")
	return nil
}

func (srv *Service) ReadAll() error {
	srv.repo.ReadAll()

	//fmt.Println(all)
	return nil
}

func (srv *Service) ReadByID() error {
	// filter := bson.M{}
	// results, err := mongoService.ReadDocuments(filter)
	// if err != nil {
	// 	log.Fatal("Error al leer documentos:", err)
	// }
	// log.Println("Documentos leídos:")
	// for _, result := range results {
	// 	log.Println(result)
	// }
	return nil
}

func (ms *Service) Update() error {
	// updateFilter := bson.M{"name": "Nuevo alojamiento"}
	// update := bson.M{"$set": bson.M{"name": "Alojamiento actualizado"}}
	// err = mongoService.UpdateDocument(updateFilter, update)
	// if err != nil {
	// 	log.Fatal("Error al actualizar el documento:", err)
	// }
	// log.Println("Documento actualizado con éxito")
	return nil
}

func (ms *Service) Delete() error {
	// deleteFilter := bson.M{"name": "Alojamiento actualizado"}
	// err = mongoService.DeleteDocument(deleteFilter)
	// if err != nil {
	// 	log.Fatal("Error al eliminar el documento:", err)
	// }
	// log.Println("Documento eliminado con éxito")
	return nil
}
