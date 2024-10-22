### 1. Nombre del Microservicio: Singular o Plural

En cuanto a la convención de nombres de microservicios, ambos enfoques (singular y plural) son válidos y generalmente dependen de las preferencias del equipo y la consistencia con otros servicios. Sin embargo, la convención más común y recomendada es usar **plural**, ya que normalmente un microservicio gestiona múltiples instancias de un recurso o entidad.

**Razón para usar plural:**
- Un microservicio generalmente maneja un conjunto de recursos o entidades. En este caso, el microservicio gestionará **múltiples solicitudes administrativas** (administrative requests), por lo que el nombre en plural es más natural y descriptivo.
  
Por lo tanto, el microservicio debería llamarse **"administrative-requests"**.

---

### 2. Propuesta para la Entidad `AdministrativeRequest`

Aquí está la propuesta de la entidad renombrada a `AdministrativeRequest` y la estructura del sistema con este nuevo enfoque.

#### Estructura de la Entidad `AdministrativeRequest`:

```go
package domain

import (
    "time"
)

type AdministrativeRequest struct {
    ID          string     // Unique identifier of the request (UUID)
    UserUUID    string     // Identifier of the user who submitted the request
    Type        string     // Type of request (e.g., "Permit", "Registration")
    Description string     // Description of the request
    Status      string     // Current status of the request (e.g., "Pending", "Approved", "Rejected")
    CreatedAt   time.Time  // Date when the request was created
    UpdatedAt   time.Time  // Date when the request was last updated
    DeletedAt   *time.Time // Date when the request was deleted (optional)
}
```

### 3. Relación entre `User` y `AdministrativeRequest`

La relación entre los usuarios y las solicitudes administrativas seguirá estando definida por el campo `UserUUID` en la entidad `AdministrativeRequest`. El microservicio de **usuarios** no almacenará directamente las solicitudes, sino que realizará consultas al microservicio de **solicitudes administrativas** para obtener la información requerida.

### 4. Estrategia de Comunicación entre Microservicios

#### Opción 1: gRPC

Ejemplo de definición en el archivo `.proto` para el microservicio de **administrative-requests**:

```proto
service AdministrativeRequestService {
    rpc GetRequestsByUser (GetRequestsByUserRequest) returns (GetRequestsByUserResponse);
}

message GetRequestsByUserRequest {
    string user_uuid = 1;
}

message GetRequestsByUserResponse {
    repeated AdministrativeRequest requests = 1;
}

message AdministrativeRequest {
    string id = 1;
    string user_uuid = 2;
    string type = 3;
    string description = 4;
    string status = 5;
    string created_at = 6;
    string updated_at = 7;
    string deleted_at = 8;
}
```

#### Opción 2: REST API

Puedes exponer un endpoint en el microservicio **administrative-requests** para que los usuarios puedan consultar las solicitudes.

Ejemplo de endpoint REST:

```http
GET /administrative-requests?user_uuid={userUUID}
```

Este endpoint devolvería una lista de **Administrative Requests** asociadas al `UserUUID`.

#### Opción 3: Mensajería (AMQP, Kafka)

Otra opción sería implementar un sistema basado en eventos para comunicar los cambios en las solicitudes administrativas. Cuando una solicitud es creada o modificada, el microservicio podría publicar eventos que el microservicio de **usuarios** consuma.

### 5. Implementación en el Microservicio de **Usuarios**

El microservicio de **usuarios** puede utilizar un cliente gRPC o REST para obtener las solicitudes administrativas de un usuario. A continuación, te muestro cómo podría lucir el servicio de usuarios al interactuar con el microservicio de solicitudes administrativas.

#### Ejemplo de Servicio en el Microservicio de Usuarios

```go
package services

import (
    "errors"
    "myapp/internal/core/domain"
    "myapp/internal/core/ports"
)

type UserService struct {
    userRepo  ports.UserRepository
    reqClient ports.AdministrativeRequestClient // Client to communicate with the administrative requests microservice
}

func (s *UserService) GetUserWithRequests(userUUID string) (*domain.UserWithRequests, error) {
    user, err := s.userRepo.FindByUUID(userUUID)
    if err != nil {
        return nil, err
    }

    requests, err := s.reqClient.GetRequestsByUserUUID(userUUID)
    if err != nil {
        return nil, errors.New("failed to fetch administrative requests for the user")
    }

    return &domain.UserWithRequests{
        User:         user,
        Requests:     requests,
    }, nil
}
```

### 6. Conclusión

1. **Nombre del Microservicio**: Usaremos **"administrative-requests"** en plural, siguiendo la convención más común para nombres de microservicios.
   
2. **Entidades y Relación**: La relación entre usuarios y solicitudes administrativas se manejará a través de `UserUUID`, permitiendo que el microservicio de usuarios consulte el microservicio de solicitudes administrativas cuando sea necesario.

3. **Opciones de Comunicación**: La comunicación entre los microservicios puede realizarse mediante gRPC, REST, o un sistema de mensajería basado en eventos.

¿Te gustaría proceder con esta estructura o ajustar algún aspecto en particular?


Dado el marco en el que estás trabajando (con microservicios separados para **usuarios** y **solicitudes administrativas**), la relación entre usuarios y solicitudes es una típica relación **uno a muchos** (1:N), donde:

- **Un usuario** puede tener **múltiples solicitudes**.
- **Una solicitud** solo pertenece a **un usuario**.

Dado que los microservicios están separados, no vas a establecer la relación directamente en la base de datos (como un `JOIN` en una base de datos monolítica). En su lugar, la relación debe manejarse a través de los **identificadores únicos** (UUID) y **consultas a los microservicios**. Aquí te dejo una descripción de cómo podrías gestionar esta relación en tu contexto.

### Estrategia para Establecer la Relación

La mejor manera de establecer esta relación en un entorno de microservicios es:

1. **Usar `UserUUID` en las solicitudes**:
   - Cada solicitud tiene un `UserUUID` que actúa como **clave foránea** en la base de datos de solicitudes.
   - Esto permite identificar a qué usuario pertenece cada solicitud.
   
2. **Consulta del microservicio de solicitudes desde el microservicio de usuarios**:
   - Cuando el microservicio de usuarios necesite obtener las solicitudes relacionadas con un usuario, enviará una **consulta al microservicio de solicitudes** utilizando el `UserUUID` como parámetro.

### Relación basada en UUID

En la entidad `AdministrativeRequest`, el campo `UserUUID` es el que establece la relación con el usuario. No necesitas agregar más información del usuario a la solicitud (como nombre o correo electrónico), ya que el microservicio de solicitudes debe consultar al microservicio de usuarios si se necesita información adicional del usuario.

#### Ejemplo de la Entidad `AdministrativeRequest`:

```go
package domain

import (
    "time"
)

type AdministrativeRequest struct {
    ID          string     // Unique identifier of the request (UUID)
    UserUUID    string     // UUID of the user who submitted the request (Foreign Key)
    Type        string     // Type of request (e.g., "Permit", "Registration")
    Description string     // Description of the request
    Status      string     // Current status of the request (e.g., "Pending", "Approved", "Rejected")
    CreatedAt   time.Time  // Date when the request was created
    UpdatedAt   time.Time  // Date when the request was last updated
    DeletedAt   *time.Time // Date when the request was deleted (optional)
}
```

### Ejemplo de la Relación en la Base de Datos (Microservicio de Solicitudes)

En el microservicio de solicitudes, el campo `UserUUID` es el que define la relación con el usuario. El esquema de la tabla de solicitudes en PostgreSQL podría ser algo como esto:

```sql
CREATE TABLE administrative_requests (
    id UUID PRIMARY KEY,
    user_uuid UUID NOT NULL,  -- Foreign Key referencing the user
    type VARCHAR(100),
    description TEXT,
    status VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

### Flujo de Consulta entre Microservicios

Dado que los usuarios y las solicitudes están en microservicios separados, el flujo típico sería:

1. **Usuario hace una solicitud**: El usuario realiza una solicitud a través del microservicio de **usuarios**.
   
2. **El microservicio de usuarios envía la solicitud al microservicio de solicitudes**: El microservicio de usuarios se comunica con el microservicio de **solicitudes** para crear una nueva solicitud. Incluye el `UserUUID` para asociar la solicitud con el usuario correspondiente.

3. **Obtener todas las solicitudes de un usuario**: Cuando el microservicio de **usuarios** necesite obtener las solicitudes de un usuario, realizará una consulta al microservicio de **solicitudes**, pasando el `UserUUID` como parámetro.

   - Por ejemplo, usando una API REST:
     ```http
     GET /administrative-requests?user_uuid={userUUID}
     ```

   - O mediante un método gRPC:
     ```proto
     rpc GetRequestsByUser (GetRequestsByUserRequest) returns (GetRequestsByUserResponse);
     ```

4. **Respuesta de las solicitudes**: El microservicio de **solicitudes** devuelve una lista de las solicitudes relacionadas con el usuario consultado.

### Ejemplo de Código para la Relación

#### Microservicio de **Solicitudes**:
En el microservicio de **solicitudes** puedes tener un repositorio que gestione la relación entre los usuarios y sus solicitudes.

```go
package repository

import (
    "database/sql"
    "myapp/internal/core/domain"
)

type AdministrativeRequestRepository struct {
    db *sql.DB
}

func (r *AdministrativeRequestRepository) FindByUserUUID(userUUID string) ([]domain.AdministrativeRequest, error) {
    query := `SELECT id, user_uuid, type, description, status, created_at, updated_at
              FROM administrative_requests WHERE user_uuid = $1`
    
    rows, err := r.db.Query(query, userUUID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var requests []domain.AdministrativeRequest
    for rows.Next() {
        var req domain.AdministrativeRequest
        err := rows.Scan(&req.ID, &req.UserUUID, &req.Type, &req.Description, &req.Status, &req.CreatedAt, &req.UpdatedAt)
        if err != nil {
            return nil, err
        }
        requests = append(requests, req)
    }

    return requests, nil
}
```

#### Microservicio de **Usuarios**:
En el microservicio de **usuarios**, para obtener las solicitudes de un usuario, puedes tener un método como el siguiente:

```go
package services

import (
    "myapp/internal/core/domain"
    "myapp/internal/core/ports"
)

type UserService struct {
    userRepo ports.UserRepository
    reqClient ports.AdministrativeRequestClient // Client to communicate with the administrative requests microservice
}

func (s *UserService) GetUserWithRequests(userUUID string) (*domain.UserWithRequests, error) {
    user, err := s.userRepo.FindByUUID(userUUID)
    if err != nil {
        return nil, err
    }

    requests, err := s.reqClient.GetRequestsByUserUUID(userUUID)
    if err != nil {
        return nil, err
    }

    return &domain.UserWithRequests{
        User:         user,
        Requests:     requests,
    }, nil
}
```

### Conclusión:

- La relación **1:N** entre **usuarios** y **solicitudes administrativas** se establece a través del campo `UserUUID` en el microservicio de **solicitudes**.
- No es necesario mantener una relación directa en la base de datos de **usuarios**. La relación se maneja mediante consultas entre los microservicios.
- Puedes usar `UserUUID` como clave para identificar las solicitudes de cada usuario, asegurando una separación clara entre microservicios.

Esta es la mejor forma de establecer y manejar la relación en un entorno de microservicios desacoplados. ¿Te gustaría que exploremos más detalles en algún área?

Para validar si un usuario ya tiene una solicitud guardada en la base de datos, puedes hacer una consulta que verifique si existe al menos una solicitud relacionada con el `UserUUID`. Aquí te muestro cómo implementar esta validación tanto a nivel de repositorio (para consultar la base de datos) como a nivel de servicio (para utilizar esta información en la lógica de negocio).

### 1. **Consulta en el Repositorio**

En el **microservicio de solicitudes**, puedes crear una función en el repositorio que verifique si existe al menos una solicitud para un `UserUUID` dado. Aquí te dejo un ejemplo:

#### Implementación en el Repositorio (PostgreSQL)

```go
package repository

import (
    "database/sql"
)

type AdministrativeRequestRepository struct {
    db *sql.DB
}

// Check if there is at least one request for the given UserUUID
func (r *AdministrativeRequestRepository) ExistsByUserUUID(userUUID string) (bool, error) {
    query := `SELECT EXISTS(SELECT 1 FROM administrative_requests WHERE user_uuid = $1)`
    
    var exists bool
    err := r.db.QueryRow(query, userUUID).Scan(&exists)
    if err != nil {
        return false, err
    }

    return exists, nil
}
```

### 2. **Lógica de Negocio en el Servicio**

Una vez que tienes la función `ExistsByUserUUID` en el repositorio, puedes utilizarla en la capa de servicio para validar si el usuario ya tiene una solicitud registrada.

#### Implementación del Servicio

```go
package services

import (
    "errors"
    "myapp/internal/core/ports"
)

type AdministrativeRequestService struct {
    repo ports.AdministrativeRequestRepository
}

// Validate if a user has a pending request
func (s *AdministrativeRequestService) ValidateUserHasRequest(userUUID string) (bool, error) {
    exists, err := s.repo.ExistsByUserUUID(userUUID)
    if err != nil {
        return false, err
    }

    if exists {
        return true, nil // The user already has at least one request in the system
    }

    return false, nil // No request found for this user
}
```

### 3. **Uso en Controlador o Cliente de Microservicios**

Si estás haciendo esta validación en el microservicio de **usuarios**, puedes llamar a este servicio a través de una API o cliente gRPC para obtener la respuesta de si el usuario ya tiene una solicitud.

#### Ejemplo de Uso en el Controlador HTTP

Si utilizas un controlador en el microservicio de **solicitudes** para exponer esta validación:

```go
package http

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/core/services"
)

func NewAdministrativeRequestHandler(r *gin.Engine, service *services.AdministrativeRequestService) {
    r.GET("/check-user-request/:userUUID", func(c *gin.Context) {
        userUUID := c.Param("userUUID")

        exists, err := service.ValidateUserHasRequest(userUUID)
        if err != nil {
            c.JSON(500, gin.H{"error": "Failed to check user requests"})
            return
        }

        if exists {
            c.JSON(200, gin.H{"message": "User already has a request"})
        } else {
            c.JSON(200, gin.H{"message": "No request found for this user"})
        }
    })
}
```

### 4. **Ejemplo de SQL para Validar Existencia**

La consulta SQL que se usa aquí es muy eficiente porque usa `SELECT EXISTS`, que devuelve un valor booleano tan pronto como encuentra una coincidencia, sin necesidad de recuperar todos los datos de la solicitud. Esto es importante para optimizar la consulta, especialmente si la tabla tiene muchas filas.

```sql
SELECT EXISTS(SELECT 1 FROM administrative_requests WHERE user_uuid = 'user-uuid-example');
```

Este tipo de consulta es preferible porque simplemente verifica la existencia de un registro, en lugar de cargar todas las solicitudes del usuario, lo que hace que sea mucho más rápida.

### 5. **Conclusión**

- **Validación eficiente**: Utilizar `SELECT EXISTS` permite validar de manera rápida si un usuario ya tiene una solicitud en la base de datos.
- **Desacoplamiento**: La validación se maneja completamente dentro del microservicio de **solicitudes**, manteniendo la separación de responsabilidades.
- **Uso en la capa de servicio**: Puedes fácilmente integrar esta validación dentro de cualquier flujo de negocio que requiera saber si un usuario ya tiene una solicitud.

¿Te gustaría que ajustemos algo más o ver detalles adicionales sobre la implementación?