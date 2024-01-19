package presenter

import (
	"time"

	domain "items/internal/domain"
)

/*
Presenters

En la arquitectura limpia, el Presenter es una capa de la aplicación que se encarga de mostrar la información al usuario de la manera en que se requiere, a menudo en forma de HTML, JSON, XML, etc. En otras palabras, el Presenter se encarga de la presentación de los datos.

El Presenter es responsable de tomar los datos que se generan a través de las interacciones entre la capa de negocios (Interactor) y la capa de acceso a datos (Repositorio) y transformarlos en un formato adecuado para su visualización o envío a otras aplicaciones o servicios.
*/

/*
Diferencias entre Presenters y DTSs
Los DTOs y los presenters son conceptos diferentes en Golang y se utilizan para propósitos distintos.

Un DTO (Data Transfer Object) es una estructura de datos que se utiliza para transferir datos entre diferentes componentes del sistema. Su objetivo principal es minimizar el acoplamiento entre las diferentes partes del sistema y permitir que los datos se transfieran de forma sencilla y clara. Los DTOs suelen ser estructuras simples que solo contienen los datos necesarios para realizar la transferencia.

Por otro lado, un presenter es una capa de la arquitectura del software que se utiliza para formatear y presentar los datos al usuario. Su objetivo principal es separar la lógica de presentación del resto del sistema y proporcionar una interfaz de usuario clara y fácil de usar. Los presenters se utilizan para convertir los datos del sistema en un formato adecuado para ser presentado al usuario.

En resumen, mientras que los DTOs se utilizan para transferir datos entre diferentes partes del sistema, los presenters se utilizan para presentar los datos al usuario en un formato adecuado. Ambos conceptos son importantes en la arquitectura de software y se utilizan en combinación con otros patrones de diseño para crear sistemas modulares y escalables.
*/

/*
En la arquitectura limpia, los DTO (Data Transfer Objects) se utilizan para transferir datos entre las diferentes capas de la aplicación, ya que cada capa tiene su propio modelo de datos. Los DTO permiten que los datos se muevan de manera transparente a través de las capas de la aplicación, sin que la capa de presentación tenga que conocer los detalles de implementación de las capas subyacentes.

En Go, los DTO se pueden definir como estructuras que representan los datos que se van a transferir entre las diferentes capas de la aplicación. Por ejemplo, en una aplicación de gestión de usuarios, se podría definir un DTO para representar la información del usuario:

```

	type UserDTO struct {
		ID    int
		Name  string
		Email string
	}

```

Los DTO se utilizan en la capa de presentación para recibir y enviar datos a la capa de Interactor y viceversa. En la capa de Interactor, los DTO se pueden utilizar para transferir datos entre las capas de negocio y la capa de Infraestructura, donde se realizan las operaciones de acceso a datos.

Es importante tener en cuenta que los DTO deben ser simples y contener solo la información necesaria para transferir los datos entre las capas. No deben contener ninguna lógica de negocio, ya que esto debe ser manejado por la capa de Interactor.

En resumen, los DTO se utilizan en Go y en la arquitectura limpia para transferir datos entre las diferentes capas de la aplicación, permitiendo que los datos se muevan de manera transparente a través de la aplicación sin que la capa de presentación tenga que conocer los detalles de implementación de las capas subyacentes.
*/
type jsonItem struct {
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Item(i *domain.Item) *jsonItem {
	var itemResponse jsonItem

	itemResponse.Code = i.Code
	itemResponse.Title = i.Title
	itemResponse.Description = i.Description
	itemResponse.Price = i.Price
	itemResponse.Stock = i.Stock
	itemResponse.Status = i.Status
	itemResponse.CreatedAt = i.CreatedAt
	itemResponse.UpdatedAt = i.UpdatedAt

	return &itemResponse
}

func Items(items domain.MapRepo) map[domain.ID]*jsonItem {
	mJson := make(map[domain.ID]*jsonItem)
	for id, i := range items {
		mJson[id] = Item(i)
	}

	return mJson
}

// func Items(items []domain.Item) []jsonItem {
// 	var itemResponse []jsonItem

// 	for _, val := range items {
// 		itemResponse = append(itemResponse, Item(val))
// 	}

// 	return itemResponse
// }
