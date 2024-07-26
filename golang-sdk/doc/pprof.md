## Pprof con Gin

Para integrar `pprof` con el framework Gin en Go, necesitas realizar algunos ajustes. A continuación, te muestro cómo hacerlo:

1. **Instalar el paquete `gin-contrib/pprof`**:

   Este paquete proporciona una forma fácil de integrar `pprof` con Gin.

   ```bash
   go get github.com/gin-contrib/pprof
   ```

2. **Importar los paquetes necesarios**:

   Necesitas importar Gin y el paquete `gin-contrib/pprof`.

   ```go
   import (
       "github.com/gin-gonic/gin"
       "github.com/gin-contrib/pprof"
   )
   ```

3. **Configurar `pprof` en tu router de Gin**:

   Utiliza el paquete `gin-contrib/pprof` para registrar las rutas de `pprof` en tu router de Gin.

   ```go
   func main() {
       // Crear un router de Gin
       router := gin.Default()

       // Registrar las rutas de pprof
       pprof.Register(router)

       // Definir tus rutas normales aquí
       router.GET("/ping", func(c *gin.Context) {
           c.JSON(200, gin.H{
               "message": "pong",
           })
       })

       // Iniciar el servidor en el puerto 8080
       router.Run(":8080")
   }
   ```

4. **Compilar y ejecutar tu aplicación**:

   Una vez que hayas agregado el código anterior, compila y ejecuta tu aplicación. Luego, puedes acceder a las herramientas de `pprof` a través de un navegador web en `http://localhost:8080/debug/pprof/`.

   ```bash
   go build -o myapp
   ./myapp
   ```

Con esto, `pprof` estará integrado con tu aplicación Gin, y podrás acceder a los perfiles de rendimiento a través de las rutas generadas. Por ejemplo:
- **Perfil de CPU**: `http://localhost:8080/debug/pprof/profile?seconds=30`
- **Goroutines activas**: `http://localhost:8080/debug/pprof/goroutine`
- **Uso de memoria**: `http://localhost:8080/debug/pprof/heap`

Esta configuración te permite usar `pprof` junto con Gin de manera eficiente, proporcionándote herramientas para analizar y mejorar el rendimiento de tu aplicación.