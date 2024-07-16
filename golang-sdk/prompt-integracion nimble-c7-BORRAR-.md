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

O sea, de nimble a cin7