
# **Documentación de la Arquitectura de OAuth2**

Este repositorio divide la funcionalidad OAuth2 en tres capas/paquetes principales:

1. **`pkgoauth2` (paquete base)**  
   - Define la **interfaz genérica** de OAuth2 (`Service`, `Config`).
   - Contiene una **implementación mínima** (`BaseConfig`) que gestiona validaciones y campos básicos (clientID, clientSecret, URLs, scopes, etc.).
   - Expone estructuras genéricas como `OAuth2Token` y `TokenClaims`.

2. **`pkgxaouth2` (implementación con `golang.org/x/oauth2`)**  
   - Implementa la interfaz `pkgoauth2.Service` usando la librería estándar de Go para OAuth2.
   - Soporta flujos como **Authorization Code** (GetAuthCodeURL, ExchangeCode) y **Refresh Token**.
   - Provee un método `BootstrapStd(...)` para configurar y crear el servicio de manera sencilla (leyendo variables de entorno o pasando parámetros).

3. **`pkgauth0` (implementación para Auth0)**  
   - Implementa la interfaz `pkgoauth2.Service`, enfocándose en **Auth0**.
   - Utiliza la librería `github.com/auth0-community/go-auth0` para validar tokens con JWK.
   - Soporta el **Client Credentials Flow** (GetClientCredentialsToken) y un mecanismo sencillo de validación de tokens (`ValidateToken`).

El objetivo es mantener un diseño **modular**, donde cada proveedor OAuth2 (Auth0, Google, GitHub, etc.) viva en un paquete propio, pero se adhiera a la **misma interfaz base** definida en `pkgoauth2`. Esto facilita la expansión o sustitución de proveedores en el futuro.

---

## **1. Paquete Base: `pkgoauth2`**

### **Estructura clave**

1. **Interface `Config`**  
   Define los métodos que toda configuración OAuth2 debe implementar:
   ```go
   type Config interface {
       Validate() error
       GetClientID() string
       GetClientSecret() string
       GetAuthURL() string
       GetTokenURL() string
       GetRedirectURL() string
       GetScopes() []string
       GetTimeout() time.Duration
   }
   ```

2. **Interface `Service`**  
   Define los métodos que todo “servicio” OAuth2 debe tener:
   ```go
   type Service interface {
       GetAuthCodeURL(state string) string
       ExchangeCode(ctx context.Context, code string) (*OAuth2Token, error)
       RefreshToken(ctx context.Context, refreshToken string) (*OAuth2Token, error)
       ValidateToken(ctx context.Context, tokenStr string) (*TokenClaims, error)
   }
   ```

3. **Struct `BaseConfig`**  
   Implementa `Config` de forma **genérica**, gestionando validaciones mínimas y campos base:
   ```go
   type BaseConfig struct {
       ClientID     string
       ClientSecret string
       AuthURL      string
       TokenURL     string
       RedirectURL  string
       Scopes       []string
       TimeoutSec   int
   }
   ```

4. **Estructuras `OAuth2Token` y `TokenClaims`**  
   Representan el token y sus claims. Se usan como base para cualquier implementación:
   ```go
   type OAuth2Token struct {
       AccessToken  string
       RefreshToken string
       TokenType    string
       Expiry       time.Time
   }

   type TokenClaims struct {
       Subject string `json:"sub,omitempty"`
       // Otros campos opcionales
   }
   ```

### **Ejemplo de uso directo (teórico)**
En la práctica, **no** usas `pkgoauth2` directamente para obtener tokens, sino que recurres a una implementación concreta (e.g. `pkgxaouth2` o `pkgauth0`). Pero, si quisieras, podrías crear tu propia implementación:

```go
// Ejemplo teórico: MiProveedorConfig y MiProveedorService
type MiProveedorConfig struct {
    pkgoauth2.BaseConfig
    CustomField string
}

// Validate, etc...

type MiProveedorService struct { /* ... */ }

// Implementar métodos de pkgoauth2.Service
```

---

## **2. Paquete `pkgxaouth2`: Implementación con `golang.org/x/oauth2`**

Este paquete está diseñado para trabajar con cualquier **proveedor OAuth2 estándar** (Google, GitHub, etc.) que cumpla el protocolo. Usa internamente la librería oficial `golang.org/x/oauth2`.

### **Archivos principales**

1. **`config.go`**  
   - Define `Config` que embebe `BaseConfig`.

2. **`service.go`**  
   - Contiene la `struct service` que implementa `pkgoauth2.Service`.
   - Usa `oauth2.Config` para generar URLs de autorización (`GetAuthCodeURL`), intercambiar códigos (`ExchangeCode`) y refrescar tokens (`RefreshToken`).
   - `ValidateToken` no está implementado (retorna un error) porque la validación local depende del proveedor o una introspección adicional.

3. **`bootstrap.go`** (o similar)  
   - Expone la función `BootstrapStd(...)`, que:
     1. **Lee variables de entorno** (o usa parámetros) para poblar la configuración (`clientID`, `tokenURL`, etc.).
     2. Crea la `Config`.
     3. Llama a `NewService(cfg)` para construir el `Service`.

### **Ejemplo de uso**

#### **1. Configuración**  
En tu `main.go` (o donde configures el servicio):
```go
import (
    "fmt"
    "log"
    "context"

    pkgoauth2std "github.com/devpablocristo/monorepo/pkgxaouth2"
)

func main() {
    svc, err := pkgoauth2std.BootstrapStd(
        "google-client-id", 
        "google-client-secret", 
        "https://accounts.google.com/o/oauth2/auth",
        "https://oauth2.googleapis.com/token",
        "http://localhost:8080/callback",
        []string{"profile", "email"},
        15, // timeout en segundos
    )
    if err != nil {
        log.Fatal("Error al inicializar OAuth2 Std: ", err)
    }

    // 2. Obtener la URL para iniciar el flujo Authorization Code
    url := svc.GetAuthCodeURL("random-state")
    fmt.Println("Visita la siguiente URL para autorizar:", url)

    // 3. Supón que recibes `code` en tu endpoint /callback...
    code := "el_code_que_devuelve_el_proveedor"
    token, err := svc.ExchangeCode(context.Background(), code)
    if err != nil {
        log.Fatal("No se pudo intercambiar el código:", err)
    }
    fmt.Println("Token obtenido:", token)
}
```

Con esto, tu aplicación realiza el **Authorization Code Flow** usando la **lib estándar** de Go para OAuth2, pero a través de la **interfaz** `pkgoauth2.Service`.

---

## **3. Paquete `pkgauth0`: Implementación para Auth0**

Este paquete se especializa en **Auth0**, usando la librería `github.com/auth0-community/go-auth0`. Maneja principalmente:

- **Client Credentials Flow**: `GetClientCredentialsToken`.
- **Validación de tokens** (`ValidateToken`) via `ValidateRequest`, decodificando JWT con las llaves de Auth0.

### **Archivos principales**

1. **`config.go`**  
   - `Config` que embebe `BaseConfig` y añade campos específicos de Auth0 (`Domain`, `Audience`).
   - Valida que `Domain` y `Audience` no estén vacíos.

2. **`service.go`**  
   - Implementa `pkgoauth2.Service`.  
   - Usa la librería `auth0-community/go-auth0` para **validar tokens** con JWK.  
   - `GetClientCredentialsToken`: genera un token de acceso vía `clientcredentials.Config`, añadiendo la `audience`.

3. **`bootstrap.go`**  
   - `Bootstrap(...)` para inicializar la configuración leyendo variables como `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID`, etc.
   - Crea el `Config` y finalmente el `service`.

### **Ejemplo de uso**

```go
import (
    "context"
    "fmt"
    "log"

    "github.com/devpablocristo/monorepo/pkg/pkgauth0"
)

func main() {
    // 1. Configurar Auth0
    svc, err := pkgauth0.Bootstrap(
        "dev-xxxx.us.auth0.com", 
        "myClientID",
        "mySecret",
        "myAudience",
        10, // timeout
    )
    if err != nil {
        log.Fatal("Error al inicializar Auth0:", err)
    }

    // 2. Obtener un token usando Client Credentials
    token, err := svc.(*pkgauth0.Service).GetClientCredentialsToken(context.Background())
    if err != nil {
        log.Fatal("Error al obtener token CC:", err)
    }
    fmt.Println("Token de Auth0 (Client Credentials):", token)

    // 3. Validar un token
    claims, err := svc.ValidateToken(context.Background(), token)
    if err != nil {
        log.Fatal("Token inválido:", err)
    }
    fmt.Printf("Claims del token: %#v\n", claims)
}
```

En este ejemplo:
- **`GetClientCredentialsToken`** obtiene un token de acceso para consumir APIs en nombre de la aplicación (no involucra usuarios).
- **`ValidateToken`** construye una petición ficticia y usa `auth0.NewJWKClient` para validar la firma JWT.

---

# **Flujos soportados**

| **Paquete**      | **Flujos principales**                                                | **Notas**                                     |
|------------------|----------------------------------------------------------------------|-----------------------------------------------|
| `pkgxaouth2`     | Authorization Code, Refresh Token                                    | `ValidateToken` pendiente de implementación   |
| `pkgauth0`       | Client Credentials Flow, Validate Token (JWT)                        | `GetAuthCodeURL`, `ExchangeCode` no implement |
| `pkgoauth2` (base) | N/A (definición de interfaz y structs, sin implementación real)       | Sirve como cimiento para proveedores.         |

---

# **Conclusión y Recomendaciones**

1. **Arquitectura Modular**:  
   - `pkgoauth2` actúa como capa base definiendo **interfaces** y **estructuras** comunes.  
   - `pkgxaouth2` y `pkgauth0` se adhieren a la misma interfaz `Service`, lo que permite cambiar de proveedor sin alterar la lógica de negocio.

2. **Elección de Paquete**:  
   - Usa **`pkgxaouth2`** si trabajas con un **proveedor OAuth2 estándar** (Google, GitHub, Okta en modo genérico, etc.).  
   - Usa **`pkgauth0`** si trabajas específicamente con **Auth0** y necesitas su lógica de validación con JWK y Client Credentials Flow.

3. **Personalizaciones**:  
   - Si más adelante necesitas **Validar tokens** en `pkgxaouth2`, deberás implementar `ValidateToken`, por ejemplo, usando un endpoint de introspección o decodificación JWT manual.  
   - Si necesitas **Authorization Code Flow** con Auth0, puedes extender `pkgauth0` implementando `ExchangeCode`, etc.

4. **Mantenimiento**:  
   - Documenta y mantén actualizadas las variables de entorno en cada “bootstrap” (`OAUTH2_CLIENT_ID`, `AUTH0_DOMAIN`, etc.).  
   - Considera agregar tests unitarios para cada paquete, validando flujos como `ExchangeCode`, `RefreshToken`, y `ValidateToken`.

Con esto, tu arquitectura OAuth2 queda **completa**, escalable y mantenible. ¡Listo para usar en tus proyectos!