### Data Transfer Object (DTO)

El DTO (Objeto de Transferencia de Datos) es un patrón de diseño utilizado para transferir datos entre sistemas o capas de una aplicación de manera simplificada. Los DTOs encapsulan datos y se utilizan comúnmente para transferir información relevante entre cliente y servidor, o entre distintas partes de un sistema, sin exponer detalles internos de cómo se almacenan o manejan esos datos. Esto promueve una mayor flexibilidad y desacoplamiento entre las capas de una aplicación, siendo especialmente útiles en el desarrollo de APIs para modelar tanto las solicitudes entrantes como las respuestas salientes, permitiendo controlar exactamente qué datos se exponen a través de la API.

#### Ejemplo de DTO en Go para una aplicación de gestión de libros:

```go
package dto

// LibroCreateDTO se utiliza para la creación de libros, definiendo lo necesario para añadir un nuevo libro.
type LibroCreateDTO struct {
    Titulo  string `json:"titulo"`
    Autor   string `json:"autor"`
    Resumen string `json:"resumen,omitempty"` // Opcional
    Anio    int    `json:"anio"`
}

// LibroResponseDTO se utiliza para enviar datos del libro al cliente, definiendo cómo se presenta un libro en las respuestas.
type LibroResponseDTO struct {
    ID      int64  `json:"id"`
    Titulo  string `json:"titulo"`
    Autor   string `json:"autor"`
    Resumen string `json:"resumen,omitempty"` // Opcional
    Anio    int    `json:"anio"`
}
```

### Presenter

El Presenter forma parte del patrón Modelo-Vista-Presentador (MVP), actuando como intermediario entre la vista (UI) y el modelo (datos de negocio). Su objetivo principal es preparar los datos del modelo para su presentación en la vista, conteniendo lógica específica de presentación que decide cómo se deben mostrar los datos al usuario.

#### Ejemplo de Presenter en Go:

```go
package presenter

import (
    "fmt"
    "strings"
    "tuProyecto/dto"
)

// LibroPresenter es responsable de preparar los datos del libro para la vista.
type LibroPresenter struct {
    Libro dto.LibroDTO
}

// FormatearTitulo convierte el título del libro a mayúsculas.
func (lp *LibroPresenter) FormatearTitulo() string {
    return strings.ToUpper(lp.Libro.Titulo)
}

// DetallesDelLibro compone una cadena con los detalles formateados del libro para la presentación.
func (lp *LibroPresenter) DetallesDelLibro() string {
    return fmt.Sprintf("Título: %s, Autor: %s, Año de Publicación: %d",
        lp.FormatearTitulo(), lp.Libro.Autor, lp.Libro.Publicacion)
}
```

### Data Access Object (DAO)

El DAO (Objeto de Acceso a Datos) es un patrón de diseño que proporciona una interfaz abstracta para acceder a datos almacenados en una base de datos, archivo, o cualquier otro medio de persistencia. Su propósito es separar la lógica de acceso a datos de la lógica de negocio de la aplicación, permitiendo que esta última sea independiente del mecanismo de almacenamiento subyacente.

#### Ejemplo de DAO en Go para la entidad Libro:

```go
package dao

import (
    "errors"
    "tuProyecto/models"
)

// LibroDAO define las operaciones de acceso a datos para los libros.
type LibroDAO interface {
    Crear(libro models.Libro) error
    ObtenerPorID(id string) (models.Libro, error)
    Actualizar(libro models.Libro) error
    Eliminar(id string) error
}

// LibroDAOImpl implementa LibroDAO usando un mapa como almacenamiento en memoria.
type LibroDAOImpl struct {
    libros map[string]models.Libro
}

func NuevoLibroDAOImpl() *LibroDAOImpl {
    return &LibroDAOImpl{libros: make(map[string]models.Libro)}
}

// Implementación de los métodos Crear, ObtenerPorID, Actualizar y Eliminar...
```

Estos ejemplos ilustran cómo el DTO se utiliza para la transferencia eficiente de datos entre capas o servicios, el Presenter para adaptar y formatear datos para la presentación, y el DAO para encapsular el acceso y manipulación de datos, demostrando la importancia de cada patrón en el diseño y arquitectura de software.