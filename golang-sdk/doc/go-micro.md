## go-micro

`go-micro` es un framework de microservicios en Go (Golang) que te ayuda a construir aplicaciones distribuidas y escalables. Está diseñado para ser sencillo de usar y proporciona las herramientas necesarias para desarrollar servicios robustos en un sistema distribuido, enfocándose en las mejores prácticas y patrones de microservicios.

Clave de `go-micro`:

1. **Pluggable**: `go-micro` ofrece una arquitectura basada en plugins, lo que permite cambiar y extender sus componentes fácilmente. Puedes elegir entre diferentes opciones para el transporte de mensajes, el descubrimiento de servicios, y el balanceo de carga, entre otros.

2. **Descubrimiento de servicios**: Automáticamente gestiona el registro y el descubrimiento de servicios utilizando sistemas como Consul, etcd, o incluso multicast DNS para entornos de desarrollo.

3. **Balanceo de carga**: Proporciona balanceo de carga de cliente y mecanismos de resilencia integrados, como reintentos y circuit breakers.

4. **RPC y Eventos**: Facilita la comunicación entre servicios usando RPC y también soporta la comunicación asíncrona mediante mensajes y eventos.

5. **API Gateway**: Incluye un gateway que permite a los servicios exponer sus funciones a través de HTTP, facilitando la integración con aplicaciones que no son parte del sistema de microservicios.

### Características y Operaciones:

### Características y Funcionalidades Adicionales de `go-micro`

1. **Transporte y Codificación**:
   - `go-micro` permite utilizar diferentes protocolos de transporte (HTTP, gRPC, NATS, etc.) y mecanismos de codificación (JSON, Protobuf, etc.).
   - Se puede configurar mediante opciones adicionales en la creación del servicio.

2. **Middleware**:
   - Soporte para middleware y wrappers que permiten agregar funcionalidades transversales como autenticación, autorización, logging, métricas, etc.
   - Ejemplo: `service.Client().Use(...)` para configurar middleware en el cliente.

3. **Configuración**:
   - `go-micro` soporta la configuración dinámica y estática mediante múltiples fuentes (archivos, entornos, servicios de configuración, etc.).
   - Ejemplo: Usar `config` para cargar configuraciones.

4. **Métricas y Monitorización**:
   - Integración con sistemas de monitorización y métricas como Prometheus para recopilar y visualizar datos operacionales.
   - Ejemplo: `service.Options().Metrics` para configurar métricas.

5. **Tracing**:
   - Soporte para tracing distribuido con herramientas como Jaeger o Zipkin, permitiendo rastrear las solicitudes a través de múltiples servicios.
   - Ejemplo: Integración con `opentracing`.

6. **Pub/Sub (Publicación/Suscripción)**:
   - Soporte para patrones de publicación y suscripción que permiten arquitecturas basadas en eventos.
   - Ejemplo: Usar `service.Server().Subscribe(...)` para suscribirse a eventos.

7. **Balanceo de Carga**:
   - Integrado con el registro de servicios para realizar balanceo de carga entre múltiples instancias de un servicio.
   - Ejemplo: `micro.RoundRobin` para configurar el balanceo de carga.

8. **Fallback y Retry**:
   - Mecanismos de retry y fallbacks para manejar fallos de comunicación de manera robusta.
   - Ejemplo: `service.Client().Retries(...)` para configurar reintentos.

9. **Autenticación y Autorización**:
   - Soporte para agregar autenticación y autorización a los servicios utilizando middleware.
   - Ejemplo: Integración con `auth` para manejar tokens JWT.

10. **Hooks y Lifecycle Management**:
    - Permite definir hooks para el ciclo de vida del servicio, como `OnStart` y `OnStop`.
    - Ejemplo: `service.Init(micro.AfterStart(func() {...}))` para ejecutar código después de que el servicio inicie.

### Ejemplo Avanzado de Uso de `go-micro`

Aquí tienes un ejemplo más completo que incluye algunas de estas características avanzadas:

```go
package main

import (
    "context"
    "fmt"

    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/logger"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/registry/consul"
    "github.com/micro/go-micro/v2/web"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "path/to/your/project/handler"
    "path/to/your/project/internal/core"
    "path/to/your/project/proto"
)

func main() {
    // Configurar el registro de servicios con Consul
    reg := consul.NewRegistry(func(op *registry.Options) {
        op.Addrs = []string{"127.0.0.1:8500"}
    })

    // Crear un nuevo servicio con go-micro
    service := micro.NewService(
        micro.Name("example.service"),
        micro.Registry(reg),
        micro.WrapHandler(loggingWrapper),  // Añadir un middleware de logging
    )

    // Inicializar el servicio
    service.Init(
        micro.AfterStart(func() error {
            fmt.Println("Service started")
            return nil
        }),
        micro.BeforeStop(func() error {
            fmt.Println("Service stopping")
            return nil
        }),
    )

    // Crear el cliente para NotificationService
    notificationClient := proto.NewNotificationService("notification.service", service.Client())

    // Crear el handler del EventService
    eventHandler := &EventServiceHandler{
        notificationClient: notificationClient,
    }

    // Registrar el handler del EventService
    proto.RegisterEventServiceHandler(service.Server(), eventHandler)

    // Crear el servicio web para exponer el REST API
    webService := web.NewService(
        web.Name("event.web"),
        web.Registry(reg),
        web.Address(":8081"),
    )

    // Inicializar el servicio web
    webService.Init()

    // Crear el UseCasePort
    useCasePort := core.NewUseCasePort()

    // Crear el handler REST
    restHandler := handler.NewRestHandler(useCasePort)

    // Configurar el router Gin
    r := gin.Default()
    r.POST("/event", restHandler.CreateEvent)
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))  // Endpoint para métricas

    // Registrar el handler REST con go-micro
    webService.Handle("/", r)

    // Ejecutar ambos servicios
    go func() {
        if err := service.Run(); err != nil {
            logger.Fatal(err)
        }
    }()

    if err := webService.Run(); err != nil {
        logger.Fatal(err)
    }
}

// Ejemplo de un middleware de logging
func loggingWrapper(fn micro.HandlerFunc) micro.HandlerFunc {
    return func(ctx context.Context, req micro.Request, rsp interface{}) error {
        logger.Info("Request received")
        return fn(ctx, req, rsp)
    }
}

// Implementación del EventServiceHandler
type EventServiceHandler struct {
    notificationClient proto.NotificationService
}

func (h *EventServiceHandler) CreateEvent(ctx context.Context, req *proto.EventRequest, res *proto.EventResponse) error {
    // Implementar la lógica para crear un evento
    fmt.Printf("Creating event: %s\n", req.Title)

    // Simular la creación del evento
    res.Message = "Event created successfully"

    // Enviar notificación después de crear el evento
    notificationReq := &proto.NotificationRequest{
        Message: "New event created: " + req.Title,
    }

    notificationRes, err := h.notificationClient.SendNotification(ctx, notificationReq)
    if err != nil {
        res.Err = err.Error()
        return err
    }
    fmt.Println("Notification Response:", notificationRes.Status)

    return nil
}
```

### Explicación del Ejemplo

- **Middleware**: Se utiliza `micro.WrapHandler(loggingWrapper)` para agregar un middleware que registra cada solicitud recibida.
- **Hooks de Ciclo de Vida**: Se utilizan `micro.AfterStart` y `micro.BeforeStop` para ejecutar código personalizado cuando el servicio se inicia o se detiene.
- **Monitorización**: Se expone un endpoint `/metrics` para Prometheus, que permite recopilar métricas del servicio.
- **Servicio Web**: Se utiliza `web.NewService` para crear un servicio web que expone endpoints RESTful.

### Pasos para usar go-micro con Consul

Estos cuatro pasos son fundamentales para configurar, inicializar y ejecutar un microservicio utilizando `go-micro` y Consul. Este flujo garantiza que el servicio esté correctamente configurado, registrado y operativo, permitiendo una comunicación y descubrimiento efectivos en una arquitectura de microservicios.


### Paso 1: Configurar el Registro de Servicios con Consul

```go
reg := consul.NewRegistry(func(op *registry.Options) {
    op.Addrs = []string{"127.0.0.1:8500"}
})
```

- **Propósito**: Configurar el registro de servicios utilizando Consul como el backend de registro.
- **Proceso**:
  - `consul.NewRegistry` crea una nueva instancia de registro que utiliza Consul.
  - La función anónima `func(op *registry.Options)` se utiliza para especificar opciones adicionales, como la dirección del servidor Consul (`127.0.0.1:8500`).

### Paso 2: Crear un Nuevo Servicio con `go-micro`

```go
service := micro.NewService(
    micro.Name("example.service"),
    micro.Registry(reg),
)
```

- **Propósito**: Crear una instancia del servicio con `go-micro`, especificando su nombre y el registro de servicios que utilizará.
- **Proceso**:
  - `micro.NewService` inicializa una nueva instancia del servicio.
  - `micro.Name("example.service")` establece el nombre del servicio como `"example.service"`.
  - `micro.Registry(reg)` especifica que el servicio utilizará Consul para el registro de servicios.

### Paso 3: Inicializar el Servicio

```go
service.Init()
```

- **Propósito**: Inicializar el servicio para prepararlo para su ejecución.
- **Proceso**:
  - `service.Init()` realiza cualquier configuración adicional necesaria y prepara el servicio para ser ejecutado.
  - Esta etapa puede incluir la configuración de dependencias, la lectura de configuraciones, y la inicialización de conexiones a bases de datos u otros servicios.

### Paso 4: Ejecutar el Servicio

```go
if err := service.Run(); err != nil {
    logger.Fatal(err)
}
```

- **Propósito**: Ejecutar el servicio, registrarlo en Consul y comenzar a escuchar solicitudes.
- **Proceso**:
  - `service.Run()` arranca el servicio, lo registra en Consul y lo pone en marcha para que comience a escuchar solicitudes entrantes en su puerto especificado.
  - Si ocurre algún error durante la ejecución, se registra el error y el programa se detiene.

### Flujo Completo

Para ilustrar cómo estos pasos se integran en un flujo completo, aquí tienes el código de ejemplo:

```go
package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/logger"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/registry/consul"
)

func main() {
    // Paso 1: Configurar el registro de servicios con Consul
    reg := consul.NewRegistry(func(op *registry.Options) {
        op.Addrs = []string{"127.0.0.1:8500"}
    })

    // Paso 2: Crear un nuevo servicio con go-micro
    service := micro.NewService(
        micro.Name("example.service"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Address(":8082"), // Especificar el puerto en el que este servicio escuchará
	)
    // Paso 3: Inicializar el servicio
    service.Init()

    // Paso 4: Ejecutar el servicio
    if err := service.Run(); err != nil {
        logger.Fatal(err)
    }
}
```

### Explicación de los Pasos

1. **Configuración del Registro con Consul**:
   - Configura Consul para que actúe como el registro de servicios. Esto permite que los servicios se registren y sean descubiertos dinámicamente.
  
2. **Creación del Servicio**:
   - Configura y crea una nueva instancia del servicio, especificando su nombre y el registro de servicios a utilizar.

3. **Inicialización del Servicio**:
   - Prepara el servicio para su ejecución, realizando configuraciones adicionales y preparando dependencias.

4. **Ejecución del Servicio**:
   - Arranca el servicio, lo registra en Consul y comienza a escuchar solicitudes entrantes. Este paso es crucial para que el servicio esté activo y disponible en el ecosistema de microservicios.


### Conclusión

`go-micro` proporciona una amplia gama de características avanzadas que facilitan la creación, gestión y operación de microservicios. Además de los pasos básicos de configuración, inicialización y ejecución, puedes aprovechar middleware, hooks de ciclo de vida, monitorización, balanceo de carga y muchas otras funcionalidades para construir sistemas robustos y escalables.
















