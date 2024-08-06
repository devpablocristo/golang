## Consul

**Consul** es una herramienta de código abierto desarrollada por HashiCorp que proporciona una solución completa para el descubrimiento de servicios, la gestión de configuraciones y la segmentación de red en un entorno distribuido. Consul está diseñado para ser altamente disponible, escalable y fácil de integrar con diversas aplicaciones y entornos de microservicios.

### Características Clave de Consul

1. **Descubrimiento de Servicios**:
   - Consul permite que los servicios se registren y se descubran entre sí. Los servicios pueden encontrar otros servicios sin necesidad de conocer sus ubicaciones exactas (IP y puertos).

2. **Gestión de Configuraciones**:
   - Consul proporciona una base de datos KV (key-value) distribuida que se puede usar para almacenar configuraciones dinámicas que los servicios pueden leer y reaccionar en tiempo real.

3. **Segmentación de Red (Service Segmentation)**:
   - Consul puede gestionar políticas de red y aplicar malla de servicios (service mesh) para controlar y asegurar la comunicación entre servicios.

4. **Supervisión y Salud de Servicios**:
   - Consul realiza comprobaciones de salud para asegurarse de que los servicios estén funcionando correctamente. Los servicios pueden registrarse con comprobaciones de salud que Consul ejecutará periódicamente.

5. **Multi-Datacenter**:
   - Consul es capaz de operar en múltiples centros de datos, proporcionando descubrimiento de servicios y configuración de clave-valor a través de diferentes ubicaciones geográficas.

### Componentes Principales de Consul

1. **Agentes**:
   - **Agentes de Cliente**: Se ejecutan en cada nodo donde los servicios están registrados. Se encargan de registrar servicios locales y ejecutar comprobaciones de salud.
   - **Agentes de Servidor**: Mantienen el estado del clúster, gestionan la base de datos KV y coordinan el descubrimiento de servicios y las operaciones de segmentación de red.

2. **Catálogo de Servicios**:
   - Un registro de todos los servicios registrados y sus instancias, junto con el estado de sus comprobaciones de salud.

3. **Base de Datos KV**:
   - Almacena configuraciones y otros datos que los servicios pueden necesitar para su funcionamiento.

4. **Comprobaciones de Salud**:
   - Verifican periódicamente el estado de los servicios para asegurarse de que están operativos.

5. **Interfaz de Usuario y API**:
   - Consul proporciona una interfaz web para visualizar el estado del clúster y una API HTTP para interactuar programáticamente con el sistema.

### Ejemplo de Uso de Consul en una Arquitectura de Microservicios

Supongamos que tenemos una aplicación de comercio electrónico con varios microservicios:

- **User Service**
- **Product Service**
- **Order Service**
- **Payment Service**

Estos servicios necesitan descubrirse entre sí y gestionar sus configuraciones dinámicamente.

#### 1. Despliegue de Consul en Kubernetes

Podemos usar Helm para desplegar Consul en un clúster de Kubernetes.

```sh
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install consul hashicorp/consul --set global.name=consul
```

#### 2. Configuración de Consul para Descubrimiento de Servicios

##### Archivo `consul-agent-config.json`:

```json
{
  "datacenter": "dc1",
  "data_dir": "/opt/consul",
  "log_level": "INFO",
  "node_name": "consul-server",
  "server": true,
  "bootstrap_expect": 1,
  "ui": true,
  "bind_addr": "0.0.0.0",
  "client_addr": "0.0.0.0",
  "advertise_addr": "<node-ip>",
  "retry_join": ["<ip-of-another-consul-server>"]
}
```

#### 3. Configuración de un Servicio para Registrar en Consul

Supongamos que tenemos un servicio de usuarios escrito en Go que necesita registrarse en Consul.

**Archivo `main.go` del User Service**:

```go
package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/hashicorp/consul/api"
)

func main() {
    // Configurar Consul
    config := api.DefaultConfig()
    config.Address = "consul-server.consul:8500"
    client, err := api.NewClient(config)
    if err != nil {
        log.Fatalf("Failed to create Consul client: %v", err)
    }

    // Registrar el servicio en Consul
    registration := &api.AgentServiceRegistration{
        ID:      "user-service",
        Name:    "user-service",
        Address: "localhost",
        Port:    8081,
        Check: &api.AgentServiceCheck{
            HTTP:     "http://localhost:8081/health",
            Interval: "10s",
        },
    }

    err = client.Agent().ServiceRegister(registration)
    if err != nil {
        log.Fatalf("Failed to register service with Consul: %v", err)
    }

    // Configurar el router Gin
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "healthy"})
    })

    r.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Get Users"})
    })

    r.Run(":8081")
}
```

#### 4. Despliegue del User Service en Kubernetes

##### Archivo `k8s/user-service-deployment.yaml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: user-service
  template:
    metadata:
      labels:
        app: user-service
    spec:
      containers:
      - name: user-service
        image: your-docker-repo/user-service:latest
        ports:
        - containerPort: 8081
        env:
        - name: CONSUL_HTTP_ADDR
          value: "consul-server.consul:8500" # Dirección del servidor Consul
---
apiVersion: v1
kind: Service
metadata:
  name: user-service
spec:
  selector:
    app: user-service
  ports:
    - protocol: TCP
      port: 8081
      targetPort: 8081
```

### Despliegue de Prometheus, Grafana e Istio

#### Despliegue de Prometheus y Grafana

Usa Helm para desplegar Prometheus y Grafana:

```sh
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

# Desplegar Prometheus
helm install prometheus prometheus-community/prometheus

# Desplegar Grafana
helm install grafana grafana/grafana --set adminPassword='YourPassword' --set service.type=NodePort
```

#### Despliegue de Istio

Instala Istio CLI y despliega Istio en tu clúster de Kubernetes:

```sh
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.9.0
export PATH=$PWD/bin:$PATH
istioctl install --set profile=demo -y
```

Etiquetar el namespace para la inyección automática de Istio:

```sh
kubectl label namespace default istio-injection=enabled
```

### Despliegue de Consul en Kubernetes

Despliega Consul en Kubernetes usando Helm:

```sh
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install consul hashicorp/consul --set global.name=consul
```

### Desplegar los Recursos en Kubernetes

Aplica los manifiestos de Kubernetes para desplegar tu aplicación:

```sh
kubectl apply -f k8s/user-service-deployment.yaml
```

### Construir y Publicar la Imagen Docker

Construye y publica la imagen Docker de tu aplicación:

```sh
docker build -t your-docker-repo/user-service:latest .
docker push your-docker-repo/user-service:latest
```

### Registrar y Descubrir Servicios

Necesitas tener un servidor de Consul en funcionamiento para que los servicios puedan registrarse y descubrirse. Consul puede ejecutarse localmente en tu máquina de desarrollo o en un servidor dedicado en tu entorno de producción. Aquí hay una guía para configurar y levantar Consul:

### Paso 1: Instalar Consul

#### En macOS:

Puedes instalar Consul usando Homebrew:

```sh
brew install consul
```

#### En Linux:

Descarga el binario desde el sitio oficial de Consul y sigue las instrucciones de instalación:

```sh
wget https://releases.hashicorp.com/consul/1.10.4/consul_1.10.4_linux_amd64.zip
unzip consul_1.10.4_linux_amd64.zip
sudo mv consul /usr/local/bin/
```

##### Debian/Ubuntu/etc
```sh
sudo apt install consul 
```

#### En Windows:

Descarga el binario desde el sitio oficial de Consul y sigue las instrucciones de instalación.


#### docker-compose:

Si esta definido en docker-compose, no hace falta instalarlo localmente.

### Paso 2: Iniciar Consul

Puedes iniciar un servidor de Consul en modo desarrollo con el siguiente comando:

```sh
consul agent -dev
```

Este comando inicia un agente de Consul en modo desarrollo en tu máquina local, escuchando en `127.0.0.1:8500`.

### Paso 3: Verificar Consul

Una vez que Consul esté en funcionamiento, puedes verificar que está corriendo accediendo a la interfaz web en tu navegador:

```
http://localhost:8500
```

Esto te llevará a la UI de Consul, donde puedes ver los servicios registrados, los nodos y otros detalles.

### Paso 4: Configurar y Ejecutar los Servicios

Con Consul en funcionamiento, ahora puedes ejecutar los servicios `EventService` y `NotificationService` para que se registren en Consul y puedan ser descubiertos.

#### EventService

##### Archivo `main.go`

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/consul"
	"github.com/micro/go-micro/v2/web"
	"path/to/your/project/handler"
	"path/to/your/project/internal/core"
)

func main() {
	// Configurar el registro de servicios con Consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:8500"}
	})

	// Crear un nuevo servicio web con go-micro
	service := web.NewService(
		web.Name("event.service"),
		web.Version("latest"),
		web.Registry(reg),
		web.Address(":8081"), // Especificar el puerto en el que este servicio escuchará
	)

	// Inicializar el servicio
	if err := service.Init(); err != nil {
		logger.Fatal(err)
	}

	// Crear el UseCasePort
	useCasePort := core.NewUseCasePort()

	// Crear el handler REST
	restHandler := handler.NewRestHandler(useCasePort)

	// Configurar el router Gin
	r := gin.Default()
	r.POST("/event", restHandler.CreateEvent)

	// Registrar el handler REST con go-micro
	service.Handle("/", r)

	// Ejecutar el servicio
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
```

#### NotificationService

##### Archivo `main.go`

```go
package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/consul"
	"path/to/your/project/proto"
)

type NotificationService struct{}

func (s *NotificationService) SendNotification(ctx context.Context, req *proto.NotificationRequest, res *proto.NotificationResponse) error {
	// Implementar la lógica para enviar una notificación
	fmt.Printf("Sending notification: %s\n", req.Message)
	res.Status = "Notification sent successfully"
	return nil
}

func main() {
	// Configurar el registro de servicios con Consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:8500"}
	})

	// Crear un nuevo servicio en un puerto específico
	service := micro.NewService(
		micro.Name("notification.service"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Address(":8082"), // Especificar el puerto en el que este servicio escuchará
	)

	// Inicializar el servicio
	service.Init()

	// Registrar el handler del servicio
	proto.RegisterNotificationServiceHandler(service.Server(), new(NotificationService))

	// Ejecutar el servicio
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
```

### Paso 5: Ejecutar el Cliente

Finalmente, puedes ejecutar el cliente que consumirá ambos servicios:

##### Archivo `client/main.go`

```go
package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/consul"
	"path/to/your/project/proto"
)

func main() {
	// Configurar el registro de servicios con Consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:8500"}
	})

	// Crear un nuevo servicio cliente
	service := micro.NewService(
		micro.Name("client.service"),
		micro.Registry(reg),
	)
	service.Init()

	// Crear clientes para los servicios Event y Notification
	eventClient := proto.NewEventService("event.service", service.Client())
	notificationClient := proto.NewNotificationService("notification.service", service.Client())

	// Crear un evento
	eventReq := &proto.EventRequest{
		Id:          "1",
		Title:       "Sample Event",
		Description: "This is a sample event",
	}

	eventRes, err := eventClient.CreateEvent(context.Background(), eventReq)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println("Event Response:", eventRes.Message)

	// Enviar una notificación
	notificationReq := &proto.NotificationRequest{
		Message: "New event created: " + eventReq.Title,
	}

	notificationRes, err := notificationClient.SendNotification(context.Background(), notificationReq)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println("Notification Response:", notificationRes.Status)
}
```

### Conclusión

Al configurar y levantar un servidor de Consul, tus microservicios pueden registrarse y ser descubiertos dinámicamente, facilitando la comunicación y escalabilidad en una arquitectura de microservicios. Cada servicio escucha en su propio puerto, y Consul se encarga de mantener la información sobre estos servicios y sus ubicaciones, permitiendo a los clientes encontrar y comunicarse con los servicios necesarios sin configuraciones estáticas.
