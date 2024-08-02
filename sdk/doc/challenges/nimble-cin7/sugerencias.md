### 1. Realizar pruebas de rendimiento y proporcionar métricas de utilización de CPU y memoria
**Herramientas:**
- **GoBench**: Para realizar pruebas de rendimiento y benchmark de las funciones de Go.
- **pprof**: Paquete integrado en Go para la recolección de perfiles de CPU y memoria.
- **Prometheus**: Sistema de monitoreo y alerta que puede recolectar métricas de CPU y memoria.

### 2. El servicio debe poder ejecutarse en una arquitectura amd64 y ser compatible con instancias de AWS EC2 tipo m7g.large
**Herramientas:**
- **Docker**: Para construir contenedores compatibles con la arquitectura amd64.
- **Kubernetes**: Para la orquestación de contenedores en la infraestructura de AWS.

**Patrones de Diseño:**
- **Arquitectura de Microservicios**: Permite la escalabilidad y despliegue independiente de cada componente del sistema.

### 3. Implementar multi-tenancy, asegurando que cada cliente de Nimble tenga diferentes credenciales de aplicación para sus integraciones con Cin7
**Herramientas:**
- **JWT (JSON Web Tokens)**: Para la gestión de autenticación y autorización por inquilino.
- **Viper**: Para la gestión de configuraciones de múltiples inquilinos.

**Patrones de Diseño:**
- **Tenant Isolation**: Asegura que los datos y configuraciones de cada inquilino estén aislados.
- **Factory Pattern**: Para crear instancias de servicios configurados con credenciales específicas de cada inquilino.

### 4. Considerar el rate limiting (limitación de tasa) y la estrategia de reintento exponencial (exponential backoff retry)
**Herramientas:**
- **Rate** (paquete de Go): Implementación de rate limiting.
- **Backoff** (paquete de Go): Para la implementación de estrategias de reintento exponencial.

**Patrones de Diseño:**
- **Token Bucket Algorithm**: Para la implementación de rate limiting.
- **Retry Pattern**: Para manejar fallos temporales en las solicitudes a Cin7.

### 5. Incluir logging y trazabilidad detallada para facilitar la depuración, preferiblemente usando OpenTelemetry
**Herramientas:**
- **OpenTelemetry**: Para la recolección de logs y trazas.
- **Zap** o **Logrus**: Paquetes de logging para Go que integran bien con OpenTelemetry.

**Patrones de Diseño:**
- **Centralized Logging**: Para consolidar logs de diferentes servicios.
- **Correlation ID**: Para rastrear solicitudes individuales a través de múltiples servicios.

### 6. Procesamiento en tiempo casi real sin necesidad de cronjobs o polling
**Herramientas:**
- **Goroutines y Channels**: Para el manejo de concurrencia en Go.
- **Webhooks**: Para recibir notificaciones en tiempo real en lugar de polling.

**Patrones de Diseño:**
- **Event-Driven Architecture**: Para manejar eventos en tiempo real de manera eficiente.
- **Observer Pattern**: Para notificar a los servicios cuando ocurren eventos específicos.

### Ejemplo de Arquitectura y Componentes
#### Estructura de Código
- **cmd/**: Punto de entrada para el servidor.
- **pkg/**
  - **api/**: Manejo de endpoints HTTP.
  - **service/**: Lógica de negocio y manejo de multi-tenancy.
  - **repository/**: Conexión y operaciones con Cin7.
  - **config/**: Gestión de configuración.
  - **middleware/**: Rate limiting y logging.
  - **models/**: Definición de estructuras de datos.

#### Dockerfile
```dockerfile
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
```

#### Configuración de Kubernetes
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cin7-integration
spec:
  replicas: 2
  selector:
    matchLabels:
      app: cin7-integration
  template:
    metadata:
      labels:
        app: cin7-integration
    spec:
      containers:
        - name: cin7-integration
          image: cin7-integration:latest
          ports:
            - containerPort: 8080
          resources:
            requests:
              memory: "256Mi"
              cpu: "500m"
            limits:
              memory: "512Mi"
              cpu: "1000m"
```
