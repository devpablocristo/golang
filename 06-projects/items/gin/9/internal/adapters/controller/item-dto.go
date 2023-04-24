package controller

import "github.com/devpablocristo/golang/06-projects/items/gin/9/internal/entity"

/*
DTOs

En Golang, un DTO (Data Transfer Object) es una estructura de datos que se utiliza para transferir datos entre diferentes componentes de un sistema. Por lo general, los DTOs son estructuras simples que solo contienen los datos necesarios para realizar la transferencia, y suelen ser utilizados para minimizar el acoplamiento entre las diferentes partes del sistema.

Un DTO puede ser definido como una estructura que contiene un conjunto de campos que representan los datos que se van a transferir. Por ejemplo, si se está construyendo un sistema de gestión de usuarios, se podría definir un DTO de usuario que contenga los campos como nombre, correo electrónico, contraseña, etc. Luego, este DTO de usuario se utilizaría para transferir datos entre diferentes componentes del sistema, como la capa de presentación y la capa de persistencia de datos.

Se pueden utilizar en combinación con otros patrones de diseño, como el patrón de repositorio, para transferir datos de la capa de persistencia de datos a la capa de presentación, sin exponer los detalles de implementación de la capa de persistencia de datos. En general, los DTOs son una forma útil de organizar los datos de un sistema de software y facilitar su transferencia entre diferentes componentes del mismo.
*/

type ItemDTO struct {
	Code        string  `json:"code"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
}

func Dto2Item(dto *ItemDTO) *entity.Item {
	var i *entity.Item

	i.Code = dto.Code
	i.Title = dto.Title
	i.Description = dto.Description
	i.Price = dto.Price
	i.Stock = dto.Stock
	i.Status = dto.Status

	return i
}
