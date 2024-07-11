# Crawler

Este proyecto implementa un web crawler en Go diseñado para rastrear y recolectar URLs en una página web específica. Aquí encontrarás instrucciones para compilar, ejecutar y probar el crawler.

## Comandos

### 1. Compilar el Proyecto

```bash
go build -o ./crawler -v ./cmd/crawler/
```

Este comando compila el proyecto y genera un ejecutable llamado `crawler` en el directorio raíz.

### 2. Ejecutar el Crawler

```bash
./crawler crawl http://quotes.toscrape.com
```

Este comando ejecuta el crawler y comienza el rastreo desde la URL proporcionada (`http://quotes.toscrape.com` en este ejemplo).


### 3. Ejecutar el Proyecto con `go run`

```bash
go run ./cmd/crawler/crawler.go ./cmd/crawler/launcher.go crawl http://quotes.toscrape.com
```

Este comando también ejecuta el crawler, pero usando el comando `go run` para ejecutar el código fuente directamente.

**NOTA:** La bandera `crawl` debe utilizarse para poder correr el programa correctamente, tanto con el binario, como con el comando.


## Comandos Make

Se proporcionan también comandos a través de un `makefile` para una fácil ejecución.

### 1. **build:**
   - **Descripción:** Compila el proyecto y genera el ejecutable `crawler` en el directorio raíz.
   - **Uso:**
     ```bash
     make build
     ```

### 2. **runbin:**
   - **Descripción:** Ejecuta el binario `crawler` con el comando `crawl` y una URL proporcionada como parámetro.
   - **Uso:**
     ```bash
     make runbin url="http://example.com"
     ```

### 3. **runcmd:**
   - **Descripción:** Ejecuta el proyecto directamente con el comando `go run`, proporcionando la URL como parámetro.
   - **Uso:**
     ```bash
     make runcmd url="http://example.com"
     ```

### 4. **test:**
   - **Descripción:** Ejecuta las pruebas del proyecto.
   - **Uso:**
     ```bash
     make test
     ```

### 5. **up:**
   - **Descripción:** Levanta los contenedores de Docker definidos en el archivo `docker-compose.yml`.
   - **Nota:** El archivo `docker-compose.yml` está configurado para correr la URL: `"http://quotes.toscrape.com"`.
   
   Si se desea cambiar la URL, modificar la línea:
   `command: ["./crawler", "crawl", "http://quotes.toscrape.com"]`

   - **Uso:**
     ```bash
     make up
     ```

### 6. **down:**
   - **Descripción:** Detiene y elimina los contenedores de Docker.
   - **Uso:**
     ```bash
     make down
     ```

### 7. **logs:**
   - **Descripción:** Muestra los registros (logs) del contenedor Docker llamado `crawler`.
   - **Consejo:** Este comando permite visualizar las URLs encontradas y visitadas durante el rastreo.
   - **Uso:**
     ```bash
     make logs
     ```

## Requisitos Previos

- Asegúrate de tener [Go](https://golang.org/doc/install) instalado en tu sistema para ejecutar los comandos que usan Docker.
- Docker y Docker Compose deben estar instalados para utilizar los comandos `up` y `down`.