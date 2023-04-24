package presenter

import (
	"time"

	entity "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/entity"
)

/*
Presenters

En Golang, un presenter es una capa de la arquitectura del software que se utiliza para formatear y presentar los datos al usuario. El objetivo principal de un presenter es separar la lógica de presentación del resto del sistema y proporcionar una interfaz de usuario clara y fácil de usar.

En la arquitectura limpia, el presenter se encuentra en la capa de aplicación del sistema, que se encarga de gestionar la lógica de negocio y actúa como una capa intermedia entre la interfaz de usuario y la capa de persistencia de datos. En esta arquitectura, el presenter se utiliza para convertir los datos de la capa de aplicación en un formato adecuado para ser presentado al usuario.

Por ejemplo, si se está construyendo una aplicación web, un presenter podría tomar los datos de un modelo de usuario y presentarlos en una página HTML. El presenter sería responsable de dar formato a los datos para que sean fácilmente legibles y entendibles por el usuario.

Los presenters pueden ser implementados de diferentes maneras, pero generalmente se utilizan para tomar los datos del sistema y convertirlos en un formato adecuado para ser mostrados al usuario. Los presenters son una forma útil de separar la lógica de presentación del resto del sistema y hacer que el sistema sea más modular y fácil de mantener.
*/

/*
DTOs

En Golang, un DTO (Data Transfer Object) es una estructura de datos que se utiliza para transferir datos entre diferentes componentes de un sistema. Por lo general, los DTOs son estructuras simples que solo contienen los datos necesarios para realizar la transferencia, y suelen ser utilizados para minimizar el acoplamiento entre las diferentes partes del sistema.

Un DTO puede ser definido como una estructura que contiene un conjunto de campos que representan los datos que se van a transferir. Por ejemplo, si se está construyendo un sistema de gestión de usuarios, se podría definir un DTO de usuario que contenga los campos como nombre, correo electrónico, contraseña, etc. Luego, este DTO de usuario se utilizaría para transferir datos entre diferentes componentes del sistema, como la capa de presentación y la capa de persistencia de datos.

Se pueden utilizar en combinación con otros patrones de diseño, como el patrón de repositorio, para transferir datos de la capa de persistencia de datos a la capa de presentación, sin exponer los detalles de implementación de la capa de persistencia de datos. En general, los DTOs son una forma útil de organizar los datos de un sistema de software y facilitar su transferencia entre diferentes componentes del mismo.
*/

/*
Diferencias entre Presenters y DTSs
Los DTOs y los presenters son conceptos diferentes en Golang y se utilizan para propósitos distintos.

Un DTO (Data Transfer Object) es una estructura de datos que se utiliza para transferir datos entre diferentes componentes del sistema. Su objetivo principal es minimizar el acoplamiento entre las diferentes partes del sistema y permitir que los datos se transfieran de forma sencilla y clara. Los DTOs suelen ser estructuras simples que solo contienen los datos necesarios para realizar la transferencia.

Por otro lado, un presenter es una capa de la arquitectura del software que se utiliza para formatear y presentar los datos al usuario. Su objetivo principal es separar la lógica de presentación del resto del sistema y proporcionar una interfaz de usuario clara y fácil de usar. Los presenters se utilizan para convertir los datos del sistema en un formato adecuado para ser presentado al usuario.

En resumen, mientras que los DTOs se utilizan para transferir datos entre diferentes partes del sistema, los presenters se utilizan para presentar los datos al usuario en un formato adecuado. Ambos conceptos son importantes en la arquitectura de software y se utilizan en combinación con otros patrones de diseño para crear sistemas modulares y escalables.
*/

type itemDTO struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Item(i entity.Item, id uint) itemDTO {
	var itemResponse itemDTO

	itemResponse.ID = id
	itemResponse.Code = i.Code
	itemResponse.Title = i.Title
	itemResponse.Description = i.Description
	itemResponse.Price = i.Price
	itemResponse.Stock = i.Stock
	itemResponse.CreatedAt = i.CreatedAt
	itemResponse.UpdatedAt = i.UpdatedAt

	return itemResponse
}

// func Items(items []entity.Item) []itemDTO {
// 	var itemResponse []itemDTO

// 	for _, val := range items {
// 		itemResponse = append(itemResponse, Item(val))
// 	}

// 	return itemResponse
// }
