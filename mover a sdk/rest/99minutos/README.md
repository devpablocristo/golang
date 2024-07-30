# Order Manager API

Solucion para el challenge propuesto.

La solucion fue escrita complemante en Golang y utiliza dos in-memory databases para persistencia.

Por lo charlado con 99minutos, solo el desarrollo de un endpoint es realmente necesario para entregar una solución aceptable.

En este caso desarrolle el endpoint para crear ordenes.

El desarrollo abarca 2 metodos:

- handler.CreateOrder.
- application.CreateOrder.

La solucion implementa un orquestador concurrente muy simple para una posible expacion del proyecto.

---

## handler.CreateOrder

El metodo recibe una solicitud HTTP POST que contiene información del envío en formato JSON y luego crea una orden de envío en la base de datos del sistema.

El código comienza configurando la respuesta HTTP con el tipo de contenido "application/json". Luego, se decodifica el cuerpo de la solicitud JSON en una estructura de datos de envío de tipo ShippingRequest.

La función valida las credenciales de inicio de sesión del usuario y verifica si el usuario tiene los permisos necesarios para crear una orden. Si las credenciales no son válidas o el usuario no tiene los permisos necesarios, se devuelve un código de estado HTTP 401.

Si las credenciales son válidas y el usuario tiene los permisos necesarios, la función llama a la función CreateOrder del objeto orderManager para crear la orden en la base de datos del sistema. Si se produce algún error durante el proceso de creación de la orden, se devuelve un código de estado HTTP 403.

Si la orden se crea con éxito, se devuelve un código de estado HTTP 201 junto con la información de la orden en formato JSON.

---

## application.CreateOrder

El código está estructurado en torno a dos paquetes: domain y port.

El paquete domain contiene las estructuras y tipos de datos relacionados con el dominio del problema, como Order, Product, Coords y constantes como SM, MD y LG que representan los diferentes tipos de paquetes.

El paquete port contiene las interfaces que definen los repositorios de órdenes y los managers de órdenes. Además, en este paquete se implementa el gestor de órdenes OrderManager que se encarga de validar y crear una orden de envío.

### Validaciones de la orden

El gestor de órdenes realiza las siguientes validaciones antes de crear una orden:

- Validación de las coordenadas de origen y destino: se verifica que las coordenadas de latitud y longitud estén dentro de los límites válidos (-90 a 90 para latitud y -180 a 180 para longitud).
- Validación de los productos: se verifica que la cantidad, peso y tamaño de cada producto sean mayores que cero.
- Validación del peso total: si el peso total de los productos es mayor que 25 kg, se devuelve un error y se informa que la orden no está disponible con el servicio estándar.

### Creación de la orden

Si la orden pasa todas las validaciones, se crea con la siguiente información:

- UUID: un identificador único generado por la librería github.com/google/uuid.
- Tipo de paquete: se determina según el peso total de los productos.
- Peso total: la suma de los pesos de los productos multiplicados por su cantidad.
- Estado: inicialmente se establece como CREATED.
- Fecha de creación: se establece como la fecha y hora actuales.
- Finalmente, la orden creada se almacena en el repositorio de órdenes proporcionado en el constructor del gestor de órdenes.

---

## Build the API

```shell script

go build  -o  ord-man  cmd/main.go

```

---

## Run the API

```shell script

./order-manager

```

---

## With Docker compose

```shell script

make up

```

or

```shell script

docker-compose up

```

### Down all the containers

```shell script

make down

```

or

```shell script

docker-compose down  --remove-orphans

```

---

## Check status

```shell script

docker-compose ps

```

---

## Create order

Para crear ordenes es necesario utilizar usuarios validos, que existan en la DDBB.
Se crean 2 usuarios, con diferentes roles para realizar pruebas en la API:

```go
internal := domain.User{
  UUID:      "1",
  Username:  "internal10",
  Email:     "internal10@99minutos.com",
  Password:  "superPass",
  Role:      domain.INTERNAL,
  CreatedAt: time.Now().Unix(),
 }

customer := domain.User{
  UUID:      "2",
  Username:  "customer",
  Email:     "customer@mail.com",
  Password:  "12345",
  Role:      domain.CUSTOMER,
  CreatedAt: time.Now().Unix(),
 }
```

```shell script

curl -X  POST  \

http://localhost:8080/api/v1/orders/create  \

-d  '{
        "email":"customer@mail.com",
        "password":"12345",
        "orig_address":{
            "coords":{
                "lat":40.712776,
                "lng":-74.005974
            },
            "zipcode":"10007",
            "address":"Broadway",
            "city":"New York",
            "province":"NY",
            "country":"USA",
            "ext_num":"1",
            "int_num":"11"
        },
        "dest_address":{
            "coords":{
                "lat":-34.583655,
                "lng":-58.4305518
            },
            "zipcode":"C1425FOD",
            "address":"Fray Justo Sta. María de Oro",
            "city":"CABA",
            "province":"Buenos Aires",
            "country":"AR",
            "ext_num":"2148",
            "int_num":"5"
        },
        "products":[
            {
                "quantity":5,
                "weight":1.2,
                "size":10
            },
            {
                "quantity":1,
                "weight":2.5,
                "size":2
            },
            {
                "quantity":2,
                "weight":0.3,
                "size":5
            },
            {
                "quantity":8,
                "weight":1.8,
                "size":15
            }
        ]
    }'

```

---

## Postman

Para utilizar con Postman se puede importar el archivo: 99minutos.postman_collection.json
