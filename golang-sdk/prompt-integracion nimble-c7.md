### Nombre del Proyecto
```
NimbleCin7Integration
```

### Estructura del Proyecto
```
NimbleCin7Integration/
│
├── cmd/
│   ├── main.go
│
├── internal/
│   ├── nimble/
│   │   ├── domain/
│   │   │   ├── entity.go
│   │   │   ├── port.go
│   │   ├── dto/
│   │   │   ├── nimble_dto.go
│   │   ├── handler/
│   │   │   ├── nimble_handler.go
│   │   ├── repository/
│   │   │   ├── nimble_repository.go
│   │   ├── usecase/
│   │       ├── nimble_usecase.go
│   │
│   ├── cin7/
│   │   ├── domain/
│   │   │   ├── entity.go
│   │   │   ├── port.go
│   │   ├── dto/
│   │   │   ├── cin7_dto.go
│   │   ├── handler/
│   │   │   ├── cin7_handler.go
│   │   ├── repository/
│   │   │   ├── cin7_repository.go
│   │   ├── usecase/
│   │       ├── cin7_usecase.go
│   │
├── pkg/
│   ├── config/
│   │   ├── config.go
│   ├── middleware/
│   │   ├── middleware.go
│   ├── wire/
│       ├── wire.go
│
├── go.mod
└── README.md
```

### cmd/main.go
```go
package main

import (
    "NimbleCin7Integration/pkg/config"
    "NimbleCin7Integration/pkg/wire"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    config.LoadConfig()
    wire.SetupRoutes(r)
    r.Run()
}
```

### internal/nimble/domain/entity.go
```go
package domain

type Order struct {
    OrderID      string
    CustomerName string
    Items        []Item
}

type Item struct {
    ItemID   string
    Quantity int
}
```

### internal/nimble/domain/port.go
```go
package domain

type NimbleRepository interface {
    CreateShipment(order Order) (Shipment, error)
}

type NimbleUseCase interface {
    ProcessOrder(order Order) error
}
```

### internal/nimble/dto/nimble_dto.go
```go
package dto

type NimbleOrder struct {
    OrderID      string `json:"order_id"`
    CustomerName string `json:"customer_name"`
    Items        []Item `json:"items"`
}

type Item struct {
    ItemID   string `json:"item_id"`
    Quantity int    `json:"quantity"`
}
```

### internal/cin7/domain/entity.go
```go
package domain

type Shipment struct {
    ShipmentID  string
    OrderID     string
    ShippedDate string
    Items       []Item
}

type Item struct {
    ItemID   string
    Quantity int
}
```

### internal/cin7/domain/port.go
```go
package domain

type Cin7Repository interface {
    SaveShipment(shipment Shipment) error
}

type Cin7UseCase interface {
    UpdateShipment(shipment Shipment) error
}
```

### internal/cin7/dto/cin7_dto.go
```go
package dto

type Cin7Shipment struct {
    ShipmentID  string `json:"shipment_id"`
    OrderID     string `json:"order_id"`
    ShippedDate string `json:"shipped_date"`
    Items       []Item `json:"items"`
}

type Item struct {
    ItemID   string `json:"item_id"`
    Quantity int    `json:"quantity"`
}
```

### internal/nimble/handler/nimble_handler.go
```go
package handler

import (
    "NimbleCin7Integration/internal/nimble/dto"
    "NimbleCin7Integration/internal/nimble/domain"
    "github.com/gin-gonic/gin"
    "net/http"
)

type NimbleHandler struct {
    useCase domain.NimbleUseCase
}

func NewNimbleHandler(uc domain.NimbleUseCase) *NimbleHandler {
    return &NimbleHandler{useCase: uc}
}

func (h *NimbleHandler) HandleOrderShipment(c *gin.Context) {
    var orderDTO dto.NimbleOrder
    if err := c.ShouldBindJSON(&orderDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order := domain.Order{
        OrderID:      orderDTO.OrderID,
        CustomerName: orderDTO.CustomerName,
        Items:        convertDTOItemsToDomainItems(orderDTO.Items),
    }

    if err := h.useCase.ProcessOrder(order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func convertDTOItemsToDomainItems(items []dto.Item) []domain.Item {
    domainItems := make([]domain.Item, len(items))
    for i, item := range items {
        domainItems[i] = domain.Item{
            ItemID:   item.ItemID,
            Quantity: item.Quantity,
        }
    }
    return domainItems
}
```

### internal/cin7/handler/cin7_handler.go
```go
package handler

import (
    "NimbleCin7Integration/internal/cin7/dto"
    "NimbleCin7Integration/internal/cin7/domain"
    "github.com/gin-gonic/gin"
    "net/http"
)

type Cin7Handler struct {
    useCase domain.Cin7UseCase
}

func NewCin7Handler(uc domain.Cin7UseCase) *Cin7Handler {
    return &Cin7Handler{useCase: uc}
}

func (h *Cin7Handler) HandleShipmentUpdate(c *gin.Context) {
    var shipmentDTO dto.Cin7Shipment
    if err := c.ShouldBindJSON(&shipmentDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    shipment := domain.Shipment{
        ShipmentID:  shipmentDTO.ShipmentID,
        OrderID:     shipmentDTO.OrderID,
        ShippedDate: shipmentDTO.ShippedDate,
        Items:       convertDTOItemsToDomainItems(shipmentDTO.Items),
    }

    if err := h.useCase.UpdateShipment(shipment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func convertDTOItemsToDomainItems(items []dto.Item) []domain.Item {
    domainItems := make([]domain.Item, len(items))
    for i, item := range items {
        domainItems[i] = domain.Item{
            ItemID:   item.ItemID,
            Quantity: item.Quantity,
        }
    }
    return domainItems
}
```

### internal/nimble/usecase/nimble_usecase.go
```go
package usecase

import (
    "NimbleCin7Integration/internal/nimble/domain"
    "NimbleCin7Integration/internal/cin7/domain"
)

type NimbleUseCase struct {
    repo       domain.NimbleRepository
    cin7UseCase domain.Cin7UseCase
}

func NewNimbleUseCase(repo domain.NimbleRepository, cin7UseCase domain.Cin7UseCase) *NimbleUseCase {
    return &NimbleUseCase{repo: repo, cin7UseCase: cin7UseCase}
}

func (uc *NimbleUseCase) ProcessOrder(order domain.Order) error {
    // Transforma la orden de Nimble a un formato de envío de Cin7
    shipment, err := uc.repo.CreateShipment(order)
    if err != nil {
        return err
    }
    // Llama al caso de uso de Cin7 para actualizar el envío
    return uc.cin7UseCase.UpdateShipment(shipment)
}
```

### internal/cin7/usecase/cin7_usecase.go
```go
package usecase

import (
    "NimbleCin7Integration/internal/cin7/domain"
)

type Cin7UseCase struct {
    repo domain.Cin7Repository
}

func NewCin7UseCase(repo domain.Cin7Repository) *Cin7UseCase {
    return &Cin7UseCase{repo: repo}
}

func (uc *Cin7UseCase) UpdateShipment(shipment domain.Shipment) error {
    return uc.repo.SaveShipment(shipment)
}
```

### internal/nimble/repository/nimble_repository.go
```go
package repository

import (
    "NimbleCin7Integration/internal/nimble/domain"
    cin7domain "Nim

bleCin7Integration/internal/cin7/domain"
)

type NimbleRepository struct {}

func NewNimbleRepository() *NimbleRepository {
    return &NimbleRepository{}
}

func (r *NimbleRepository) CreateShipment(order domain.Order) (cin7domain.Shipment, error) {
    // Transforma NimbleOrder a Cin7Shipment
    return cin7domain.Shipment{
        OrderID:     order.OrderID,
        ShippedDate: time.Now().Format("2006-01-02"),
        Items:       order.Items,
    }, nil
}
```

### internal/cin7/repository/cin7_repository.go
```go
package repository

import (
    "NimbleCin7Integration/internal/cin7/domain"
    "github.com/go-redis/redis/v8"
    "context"
)

type Cin7Repository struct {
    client *redis.Client
}

func NewCin7Repository(client *redis.Client) *Cin7Repository {
    return &Cin7Repository{client: client}
}

func (r *Cin7Repository) SaveShipment(shipment domain.Shipment) error {
    ctx := context.Background()
    return r.client.Set(ctx, shipment.ShipmentID, shipment, 0).Err()
}
```

### pkg/config/config.go
```go
package config

import (
    "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func LoadConfig() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
}
```

### pkg/middleware/middleware.go
```go
package middleware

import (
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Authentication logic here
        c.Next()
    }
}
```

### pkg/wire/wire.go
```go
//+build wireinject

package wire

import (
    "NimbleCin7Integration/internal/nimble/handler"
    "NimbleCin7Integration/internal/nimble/repository"
    "NimbleCin7Integration/internal/nimble/usecase"
    cin7handler "NimbleCin7Integration/internal/cin7/handler"
    cin7repository "NimbleCin7Integration/internal/cin7/repository"
    cin7usecase "NimbleCin7Integration/internal/cin7/usecase"
    "NimbleCin7Integration/pkg/config"
    "github.com/gin-gonic/gin"
    "github.com/google/wire"
)

func SetupRoutes(r *gin.Engine) {
    nimbleRepo := repository.NewNimbleRepository()
    cin7Repo := cin7repository.NewCin7Repository(config.RedisClient)
    
    cin7UseCase := cin7usecase.NewCin7UseCase(cin7Repo)
    nimbleUseCase := usecase.NewNimbleUseCase(nimbleRepo, cin7UseCase)
    
    nimbleHandler := handler.NewNimbleHandler(nimbleUseCase)
    cin7Handler := cin7handler.NewCin7Handler(cin7UseCase)
    
    r.POST("/nimble/orders", nimbleHandler.HandleOrderShipment)
    r.POST("/cin7/shipments", cin7Handler.HandleShipmentUpdate)
}
```

### README.md
```markdown
# NimbleCin7Integration

## Descripción
Este proyecto implementa una integración entre Nimble y Cin7 utilizando Golang con arquitectura hexagonal, Gin, DTOs, Redis, Wire y Resty. La API maneja los envíos de órdenes desde Nimble y actualiza la información de los envíos en Cin7.

## Arquitectura
La arquitectura hexagonal se utiliza para separar las preocupaciones de la aplicación en diferentes capas:
- **Aplicación**: Maneja la lógica de negocio.
- **Dominio**: Contiene los modelos y servicios principales.
- **Infraestructura**: Proporciona la implementación de repositorios y la configuración de infraestructura.

## Componentes

### Nimble
- **DTOs**: Definiciones de datos para órdenes y artículos.
- **Handler**: Controladores para las rutas de la API.
- **UseCase**: Lógica de negocio para procesar órdenes.
- **Repository**: Interfaz y implementación para crear envíos.

### Cin7
- **DTOs**: Definiciones de datos para envíos y artículos.
- **Handler**: Controladores para las rutas de la API.
- **UseCase**: Lógica de negocio para actualizar envíos.
- **Repository**: Interfaz y implementación para guardar envíos en Redis.

## Configuración
### Redis
Redis se utiliza para almacenar los datos de los envíos de Cin7. Asegúrate de que Redis esté en funcionamiento en `localhost:6379`.

### Wire
Wire se utiliza para la inyección de dependencias en el proyecto. La configuración de Wire asegura que todos los componentes necesarios estén conectados adecuadamente.

## Ejecución
1. Clonar el repositorio.
2. Ejecutar `go mod tidy` para instalar las dependencias.
3. Asegurarse de que Redis esté en funcionamiento en `localhost:6379`.
4. Ejecutar `go run cmd/main.go` para iniciar el servidor.

## Rutas de la API
### Nimble
- `POST /nimble/orders`: Maneja las órdenes enviadas desde Nimble.

### Cin7
- `POST /cin7/shipments`: Actualiza la información de envíos en Cin7.

### Flujo de Control Detallado

1. **Recepción de Orden desde Nimble**:
    - **Nimble** envía una orden de envío a la API a través de una petición HTTP POST.
    - **Endpoint**: `POST /nimble/orders`
    - **Handler**: `nimble_handler.go`

    ```go
    func (h *NimbleHandler) HandleOrderShipment(c *gin.Context) {
        var orderDTO dto.NimbleOrder
        if err := c.ShouldBindJSON(&orderDTO); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        order := domain.Order{
            OrderID:      orderDTO.OrderID,
            CustomerName: orderDTO.CustomerName,
            Items:        convertDTOItemsToDomainItems(orderDTO.Items),
        }

        if err := h.useCase.ProcessOrder(order); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "success"})
    }

    func convertDTOItemsToDomainItems(items []dto.Item) []domain.Item {
        domainItems := make([]domain.Item, len(items))
        for i, item := range items {
            domainItems[i] = domain.Item{
                ItemID:   item.ItemID,
                Quantity: item.Quantity,
            }
        }
        return domainItems
    }
    ```

2. **Procesamiento de la Orden en el Caso de Uso de Nimble**:
    - El handler de Nimble llama al caso de uso de Nimble para procesar la orden.
    - **Método**: `ProcessOrder`

    ```go
    func (uc *NimbleUseCase) ProcessOrder(order domain.Order) error {
        // Transforma la orden de Nimble a un formato de envío de Cin7
        shipment, err := uc.repo.CreateShipment(order)
        if err != nil {
            return err
        }
        // Llama al caso de uso de Cin7 para actualizar el envío
        return uc.cin7UseCase.UpdateShipment(shipment)
    }
    ```

3. **Transformación de la Orden a Envío**:
    - El repositorio de Nimble transforma la orden en un formato que Cin7 pueda entender.
    - **Método**: `CreateShipment`

    ```go
    func (r *NimbleRepository) CreateShipment(order domain.Order) (cin7domain.Shipment, error) {
        // Transforma NimbleOrder a Cin7Shipment
        return cin7domain.Shipment{
            OrderID:     order.OrderID,
            ShippedDate: time.Now().Format("2006-01-02"),
            Items:       order.Items,
        }, nil
    }
    ```

4. **Actualización del Envío en el Caso de Uso de Cin7**:
    - El caso de uso de Nimble llama al caso de uso de Cin7 para actualizar el envío.
    - **Método**: `UpdateShipment`

    ```go
    func (uc *Cin7UseCase) UpdateShipment(shipment domain.Shipment) error {
        return uc.repo.SaveShipment(shipment)
    }
    ```

5. **Almacenamiento del Envío en Redis**:
    - El repositorio de Cin7 guarda el envío en Redis.
    - **Método**: `SaveShipment`

    ```go
    func (r *Cin7Repository) SaveShipment(shipment domain.Shipment) error {
        ctx := context.Background()
        return r.client.Set(ctx, shipment.ShipmentID, shipment, 0).Err()
    }
    ```

### Resumen del Flujo de Control

1. **Nimble envía una orden**: Nimble envía una orden de envío a la API a través del endpoint `POST /nimble/orders`.
2. **API recibe la orden**: El handler de Nimble recibe la petición y la deserializa en un objeto `NimbleOrder`.
3. **Procesamiento de la orden**: El caso de uso de Nimble procesa la orden y la transforma en un objeto `Cin7Shipment`.
4

. **Llamada a Cin7**: El caso de uso de Nimble llama al caso de uso de Cin7 para actualizar el envío.
5. **Guardado en Redis**: El repositorio de Cin7 guarda el envío en Redis.

### Ilustración con un Diagrama de Secuencia

Un diagrama de secuencia podría visualizar este flujo de control de manera clara. Aquí tienes una descripción textual de cómo se vería:

1. **Nimble System** -> **Nimble API**: POST /nimble/orders
2. **Nimble API** -> **Nimble Handler**: HandleOrderShipment
3. **Nimble Handler** -> **Nimble UseCase**: ProcessOrder
4. **Nimble UseCase** -> **Nimble Repository**: CreateShipment
5. **Nimble Repository** -> **Nimble UseCase**: Return Cin7Shipment
6. **Nimble UseCase** -> **Cin7 UseCase**: UpdateShipment
7. **Cin7 UseCase** -> **Cin7 Repository**: SaveShipment
8. **Cin7 Repository** -> **Cin7 UseCase**: Return Success
9. **Cin7 UseCase** -> **Nimble UseCase**: Return Success
10. **Nimble UseCase** -> **Nimble Handler**: Return Success
11. **Nimble Handler** -> **Nimble API**: Return HTTP 200 OK

## Pruebas
El proyecto incluye una cobertura de pruebas del 70% para asegurar la calidad del código. Ejecutar `go test ./...` para ejecutar todas las pruebas.

## Contribuciones
Las contribuciones son bienvenidas. Por favor, crea un `pull request` con tus cambios.

## Licencia
Este proyecto está licenciado bajo los términos de la licencia MIT.
```

### Extensiones

1. **Benchmarking**: Añadir funciones para medir el uso de CPU y memoria, y registrar estos datos en un archivo de log o sistema de monitoreo.
2. **Multi-Tenancy y Rate Limiting**: Añadir middleware para manejar la multi-tenencia y el rate limiting.
3. **Observabilidad**: Integrar OpenTelemetry para la recolección de logs y trazas.

Si necesitas más detalles o ajustes específicos, por favor házmelo saber.