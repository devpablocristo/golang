# Challenge

## Descripción del problema

En **Mercado Libre** trabajamos con `articulos` de sellers que los venden a traves de nuestro marketplace. El objetivo de este desafío es realizar una aplicación la cual exponga un _API_ que permita realizar algunas operaciones de _CRUD_ para cada una de esas dos entidades con algunas reglas de negocio sobre ellas.

### Articulos

Un item, tiene la información básica sobre el artículo que será anunciado por nosotros.

#### Representación de un item

A continuación un ejemplo de la representación _JSON_ de un item:

```json
{
  "id": 10,
  "code": "SAM27324354",
  "title": "Tablet Samsung Galaxy Tab S7",
  "description": "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
  "price": 150000.35,
  "stock": 3,
  "status": "ACTIVE",
  "created_at": "2020-05-10T04:20:33Z",
  "updated_at": "2020-05-10T05:30:00Z"
}
```

#### Reglas sobre artículos

1. Los _ids_ deben ser generados automáticamente.
2. Los campos `code`, `title`, `description`, `price`, `stock`, `photos` son obligatórios.
3. El campo `code` debe ser único.
4. Los campos `status`, `created_at`, `updated_at` son automáticamente generados por el sistema. La API no debería permitir modificaciones sobre ellos.
5. El campo `status` puede tener los siguientes valores:

- `ACTIVE`: Un item que tiene stock disponible.
- `INACTIVE`: Un item que el valor de stock es cero (0).

#### Desafío

Usando la siguientes estructura de `Item` permita las siguientes funcionalidades. Para crear o editar items, tenga en cuenta las reglas descritas anteriormente.

1. **Cree nuevos items.**

_Request_:

```
POST v1/items
```

_Body_:

```json
{
  "code": "SAM27324354",
  "title": "Tablet Samsung Galaxy Tab S7",
  "description": "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
  "price": 150000.31,
  "stock": 15
}
```

_Response_:

Usted define la respuesta, hace parte del desafío

2. **Actualice un item**

_Request_:

```
PUT v1/items/{id}
```

_Body_:

```json
{
  "code": "SAM27324354",
  "title": "Tablet Samsung Galaxy Tab S7",
  "description": "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
  "price": 158000.65,
  "stock": 25
}
```

_Response_:

Usted define la respuesta, hace parte del desafío

3. **Obtener un item por ID**

Retorna el item correspondiente al _ID_.

_Request_:

```
GET v1/items/{id}
```

_Response_:

Usted define la respuesta, hace parte del desafío

4. **Eliminar Item**

Elimina el item correspondiente al _ID_.

_Request_:

```
DELETE v1/items/{id}
```

_Response_:

Usted define la respuesta, hace parte del desafío

5. **Obtener todos los items (opcional: permitir filtrado por _Status_)**

Retorna los items filtrados por _Status_ (opcional). Los resultados vienen organizados por fecha de actualización del más reciente al más antiguo. Debe dar la opción al usuario de poner un límite de resultados a la busqueda.

_Request_:

```
GET v1/items?status={status}&limit={limit}
```

Donde:

- `status`: Es el filtro por el estado del item; es un parámetro opcional:

  - `No esta especificado`: Retorna todos los items sin importar su valor en el campo el estado
  - `ACTIVE`: Retorne los items activos
  - `INACTIVE`: Retorne los items inactivos

- `limit`: Es el tamaño solicitado de resultados en la página. Es un parámetro opcional, su valor default es 10, y su valor máximo es 20.

_Response_:

La respuesta debe seguir la siguiente estructura de campos:

- `totalPages`: El número total de items que contienen resultados para la búsqueda hecha.
- `data`: Un array con los objetos conteniendo los items solicitados en el request.

```json
{
  "totalPages": 1,
  "data": [
    {
    "id": 10,
    "code": "SAM27324354",
    "title": "Tablet Samsung Galaxy Tab S7",
    "description": "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
    "price": 150000,
    "stock": 3,
    "status": "ACTIVE"
    "created_at": "2020-05-10T04:20:33Z",
    "updated_at": "2020-05-10T05:30:00Z"
    }
  ]
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


{
  "code": "apple-iphone13",
  "title": "Apple iPhone 13",
  "description": "The latest iPhone 13 with A15 Bionic chip, 5G capable, and 128GB of storage",
  "price": 999.99,
  "stock": 100
}

{
  "code": "samgts7",
  "title": "Tablet Samsung Galaxy Tab S7",
  "description": "Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 4GB de memoria RAM",
  "price": 158000.65,
  "stock": 25
}