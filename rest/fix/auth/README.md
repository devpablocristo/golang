# Challenge

## Descripción del problema

Nesecitamos tener rutas protegidas (que sean accesibles solo por usuarios autenticados). El objetivo es definir rutas publicas para el login y registro de usuarios y tener una ruta protegida que sea accesible solo si el usuario envía un token de acceso.

### Usuarios

Tenemos usuarios que guardan como favoritos algunos de los items. Para propósito de este ejercicio, la única información con la que vamos a contar sobre el usuario es su email. Y los usuarios pueden tener algunos items marcados como favoritos.

Solo si el usuario está autenticado puede marcar como favoritos sus artículos, y listar sus artículos favoritas.

Queda a su criterio escoger una forma para autenticar el usuario.

#### Desafío

Usando la estructura `User` permita las funcionalidades registro e ingreso de usuarios.

1. **Crear un usuario**

_Request_:

```
POST v1/users/
```

_Body_:

```json
{
  "email": "test_user@mercadolibre.com",
  "password": "test"
}
```

_Response_:

Usted define la respuesta, hace parte del desafío
<br/><br/>

2. **Autentique un usuario**

_Request_:

```
POST v1/users/login
```

_Body_:

```json
{
  "email": "test_user@mercadolibre.com",
  "password": "test"
}
```

_Response_:

```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjY0MDA5YTAzNTc6MZy1NmQ3M2MwMDAyNiIsIkVtYWlsIjoib3NjYXIuc2Fsb21vbjg5QGdtYWlsLmNvbSIsIktleVR5cGUiOiJhY2Nlc3MiLCJleHAiOjE2Nzg4OTE0NTMsImlzcyI6ImF1dGguc2VydmljZSJ9.lu-svnQJklWOS37pq-eXDVhzLgshpIEnuL3QROBuR7o"
}
```

<br/>

3. **Retorna info del usuario autenticado**

Para este ejercicio solo retorne los los datos del usuario.

_Request_:

Adapte su _request_ para enviar los datos de autenticación del usuario

```
GET v1/users/me
```

_Response_:

```json
{
  "data": {
    "id": 1,
    "email": "test_user@mercadolibre.com",
    "created_at": "2020-05-10T04:20:33Z",
    "updated_at": "2020-05-10T05:30:00Z"
  }
}
```

## Criterios de calificación

Esperamos que el código que usted va a crear sea considerado por usted como _"Production Ready"_; por favor use las buenas prácticas a las cuáles usted está acostumbrado en su rutina de desarrollo de código.

Para la evaluación de su código, esperamos que su código sea portable. Esperamos que usted nos provea un comando para correr fácilmente en el ambiente local, la solución del problema.

Para el desarrollo del desafío vamos a utilizar Golang como lenguaje.

Dentro de los criterios que vamos a tener en cuenta a la hora de revisar su código, revisaremos:

- Resuelve el problema propuesto
- Organización y estructura del proyecto
- Mantenibilidad
- Facilidad para hacer tests
- Valoraremos adicionalmente si usa alguna arquitectura limpia (ej. arquitectura hexagonal).
