### 1. Usar Swagger con Comentarios en Go

#### Paso 1: Instalar `swaggo/swag`

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

#### Paso 2: Anotar tu API

Agrega comentarios de Swagger a tus controladores de API. Aquí tienes un ejemplo simple:

```go
// @Summary Get user by ID
// @Description get user by ID
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User   "user data"
// @Router /user/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

#### Paso 3: Generar Documentación

Ejecuta `swag init` en la raíz de tu proyecto para generar la documentación.

#### Paso 4: Servir la Documentación

Usa `swaggo/http-swagger` para servir tu documentación desde tu aplicación.

```go
import "github.com/swaggo/http-swagger"

// Asume que ya tienes un *gin.Engine llamado router
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

### 2. Crear YAMLs de OpenAPI Manualmente

#### Paso 1: Escribe el YAML

Crea un archivo `api-spec.yaml` y define tu API. Aquí tienes un ejemplo básico:

```yaml
openapi: 3.0.0
info:
  title: Sample API
  version: 0.1.0
paths:
  /users/{userId}:
    get:
      summary: Gets a user by ID.
      parameters:
        - name: userId
          in: path
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: A user object.
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/User'
components:
  schemas:
    User:
      type: object
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
```

#### Paso 2: Usar Swagger UI o Editor

Para visualizar tu documentación, puedes usar Swagger UI localmente o en línea a través de Swagger Editor.

#### Paso 3: Actualizaciones Manuales

Cada vez que tu API cambie, actualiza tu archivo `api-spec.yaml` manualmente y vuelve a cargarlo en Swagger UI o Editor para ver los cambios.

### 3. Enfoque Menos Laborioso: Uso de Frameworks

#### Paso 1: Elegir un Framework

Para este ejemplo, usaremos Gin, que es popular y tiene integración con Swagger a través de `swaggo/gin-swagger`.

#### Paso 2: Instalar Dependencias

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger
```

#### Paso 3: Definir Rutas con Gin

Crea tu aplicación con Gin y define las rutas.

```go
package main

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func main() {
    router := gin.Default()

    // Define tus rutas aquí

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router.Run(":8080")
}
```

#### Paso 4: Generar Swagger Automáticamente

Añade comentarios básicos a tus handlers en Gin y utiliza `swag init` para generar la documentación automáticamente. Gin y `gin-swagger` facilitan la integración, reduciendo la necesidad de anotaciones detalladas.

### Conclusión

Estas tres aproximaciones ofrecen diferentes niveles de control y facilidad para trabajar con Swagger en Go. Para un desarrollador junior, empezar con el enfoque de comentarios puede ser una buena forma de entender cómo se estructura la documentación de una API. A medida que te familiarices con Swagger y tus necesidades evolucionen, podrías optar por crear archivos YAML manual