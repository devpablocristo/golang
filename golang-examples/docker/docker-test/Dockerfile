# Usar la imagen base oficial de Go
FROM golang:1.22.3-alpine3.20

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de la aplicación al contenedor
COPY . .

# Construir la aplicación
RUN go build -o /app/bin/myapp ./cmd/rest/main.go

# Exponer el puerto en el que la aplicación estará escuchando
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/app/bin/myapp"]
