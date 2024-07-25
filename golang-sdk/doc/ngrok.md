La doc pertenece a otra API pero se uso esa misma api para crear la config de esta, igual revisar y ajustar.

### Configuración de Ngrok, Docker, Docker-Compose y Golang, utilizando un dominio específico

#### Estructura del Proyecto

```
/project-root
  - Dockerfile
  - docker-compose.yml
  - ngrok.yml
  - main.go
```

### Archivo Dockerfile

```dockerfile
# Utiliza una imagen base oficial de Golang para construir la aplicación
FROM golang:1.22.3 as builder

WORKDIR /app

# Copia todos los archivos al contenedor
COPY . .

# Construye la aplicación
RUN go build -o main .

# Utiliza una imagen base más ligera para ejecutar la aplicación
FROM alpine:latest

WORKDIR /app

# Copia el binario construido desde la fase anterior
COPY --from=builder /app/main .

# Asegúrate de que el binario tenga permisos de ejecución
RUN chmod +x ./main

# Expone el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
```

### Archivo docker-compose.yml

```yaml
version: "3.8"

services:
  rest:
    container_name: nimcin7
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    networks:
      - app-network
    restart: on-failure

  ngrok:
    image: ngrok/ngrok:latest
    container_name: ngrok
    command: ["start", "--all", "--config", "/etc/ngrok.yml"]
    volumes:
      - ./ngrok.yml:/etc/ngrok.yml
    ports:
      - "5050:5050"
    networks:
      - app-network
    restart: unless-stopped

networks:
  app-network:
    driver: bridge
```

### Archivo ngrok.yml

```yaml
version: "2"
authtoken: YOUR_NGROK_AUTHTOKEN
web_addr: "127.0.0.1:5050"
tunnels:
  my-tunnel:
    proto: http
    addr: rest:8080
    domain: brave-dane-forcibly.ngrok-free.app
```

### Archivo main.go

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	// Iniciar el servidor HTTP de forma concurrente
	go func() {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Iniciar ngrok para exponer el puerto 8080
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(config.WithForwardsTo(":8080")),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Println("Ingress established at:", listener.URL())

	return http.Serve(listener, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ngrok-go!")
}
```

### Pasos para Ejecutar

1. **Construir y Levantar los Servicios**:
   Navega al directorio del proyecto y ejecuta:

   ```sh
   docker-compose up --build
   ```

2. **Verificar los Logs**:
   Verifica los logs del contenedor `ngrok` para asegurarte de que se está utilizando el dominio especificado:

   ```sh
   docker logs ngrok
   ```

### Detalles Adicionales

- Asegúrate de reemplazar `YOUR_NGROK_AUTHTOKEN` en `ngrok.yml` con tu token de autenticación de Ngrok.
- La configuración de `web_addr` en `ngrok.yml` permite que la interfaz web de Ngrok esté disponible en el puerto 5050.
- Puedes acceder a la interfaz web de Ngrok para inspeccionar el tráfico HTTP en `http://localhost:5050`.
- El dominio especificado (`brave-dane-forcibly.ngrok-free.app`) se utilizará para exponer tu servicio.
- La URL pública también puede obtenerse iniciando sesión en tu cuenta de Ngrok y revisando los túneles activos.

### Verificación

Para asegurarte de que todo está funcionando correctamente, sigue estos pasos:

1. Abre un navegador web y navega a la URL proporcionada por ngrok. Esta URL se mostrará en los logs del contenedor `ngrok`.
2. Asegúrate de que la aplicación responda correctamente mostrando "Hello from ngrok-go!".

### Conclusión

Esta configuración te permite exponer tu aplicación Golang a través de Ngrok utilizando un dominio específico, todo orquestado con Docker y Docker-Compose. Esto facilita el desarrollo y las pruebas en un entorno local con acceso a una URL pública personalizada.