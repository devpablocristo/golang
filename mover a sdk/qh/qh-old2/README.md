Implementar:
SOLID
CORS
JWT

### POST JSON example

`
json
{
    "uuid": "1",
    "name": "Homero",
    "lastname":"Simpson"
    "age":39
}
`

`
json
{
    "uuid": "2",
    "name": "Bart",
    "lastname": "Simpson"
    "age":10
}
`ç




###################### model for readme file: ###########################


# Backend - Patients

This is a PoC for a patients API

## API endpoints


| Method | URL                             | Description                       |
|--------|---------------------------------|-----------------------------------|
| GET    | /api/v1/patients                | Get all patients                  |
| GET    | /api/v1/patients/:id            | Get one patient                   |
| POST   | /api/v1/patients                | Add one patient                   |

## Directories

cmd: entrypoint, lauch all microservices with goroutines.
apps: microservices backend and frontend.
src: microservices business concepts.



Post

[
    {
        "book": {
            "author": {
                "firstname": "Gabriel",
                "lastname": "Garcia Marquez"
            },
            "title": "100 años de soledad",
            "price": 97,
            "isbn": "0060929790"
        },
        "stock": 31
    },
    {
        "book": {
            "author": {
                "firstname": "Frank",
                "lastname": "Herbert"
            },
            "title": "Dune",
            "price": 53.79,
            "isbn": "0340960191"
        },
        "stock": 12
    },
    {
        "book": {
            "author": {
                "firstname": "Isaac",
                "lastname": "Asimov"
            },
            "title": "Fundation",
            "price": 28.5,
            "isbn": "0-553-29335-4"
        },
        "stock": 41
    }
]

### Ports

In one hand, we have the ports which are interfaces that define how the communication between an actor and the core has to be done. Depending on the actor, the ports has different nature:

- Ports for driver actors, define the set of actions that the core provides and expose to the outside. Each action generally correspond with a specific case of use.

- Ports for driven actors, define the set of actions that the actor has to implement.


### Adapters

In the other hand, we have the adapters that are responsible of the transformation between a request from the actor to the core, and vice versa. This is necessary, because as we said earlier the actors and the core “speaks” different languages.

- An adapter for a driver port, transforms a specific technology request into a call on a core service.

- An adapter for a driven port, transforms a technology agnostic request from the core into an a specific technology request on the actor.


Dependency Injection
After the implementation is done, then it is necessary to connect, somehow, the adapters to the corresponding ports. This could be done when the application starts and it allow us to decide which adapter has to be connected in each port, this is what we call “Dependency injection”. For example, if we want to save data into a mysql database, then we just have to plug an adapter for a mysql database into the corresponding port or if we want to save data in memory (for testing) we need to plug an “in memory database” adapter into that port.








/////////////////////////////////


# Hexagonal Architecture in Golang

This application implements a Golang backend using Hexagonal Architecture which aims to decouple the components of the domain/application from the frameworks like http server, storage and others.

To demonstrate this, this application implements a simple URL Shortener which makes use of different adapters. The application can be executed from command-line, or as a service (HTTP and gRPC).

## Goals
- Implement some service using the principles of the Hexagonal Architecture.
- Use more than one framework for the same function, to demonstrante how the ports and adapters works in practice.
- Create a project template.

#### Storage adapters:
- MongoDB
- Redis

#### Serializer adapters:
- JSON
- MessagePack

#### Communication adapters:
- Gin http server
- Chi http server
- gRPC

## Steps to run the cli application
The application needs an instance of Redis or MongoDB running. Just check the [cmd/cli/main.go](cmd/cli/main.go) or [cmd/http/main.go](cmd/http/main.go) and adjust the connection string for your servers.

```bash
$ go run cmd/cli/main.go
```
You would get an output like the following:

```bash
Generating Code from URL: http://www.google.com
{"code":"LWK9v1Qng","url":"http://www.google.com","created_at":1650829428}

Retrieve the URL from code: LWK9v1Qng
{"code":"LWK9v1Qng","url":"http://www.google.com","created_at":1650829428}
```


## Steps to run the service application
```bash
$ go run cmd/http/main.go
```
You would get an output like the following:

```bash
CHI  listening on :8000
GIN  listening on :9000
GRPC listening on :7000
```

From now, you can perform HTTP requests on http://localhost:8000 and http://localhost:9000

### HTTP: Get Url (or redirect if call from browser)
```
Method GET
Endpoint: <hostname:port>/:code
```
Code is the short code generated from a given URL.

### HTTP: Generate Shortcode
```
Method POST
Endpoint: <hostname:port>/
```
Payload: 
```json
{
    "url": "http://www.google.com"
}
```
![](_assets/post.jpg)

Example using curl:
```bash
curl --request POST 'localhost:9000/' \
--header 'Content-Type: application/json' \
--data '{"url": "http://www.google.com.br"}'
```

Expected output
```bash
{"code":"AS4YOJQ7R","url":"http://www.google.com.br","created_at":1650830243}
```

### gRPC
The code show how to use [evans](https://github.com/ktr0731/evans) in interactive mode to perform gRPC call.

```bash
cd adapter/grpc/proto 
evans --host localhost --port 7000 --proto shortener_msg.proto,shortener_service.proto
```
![](_assets/evans.jpg)


## TO-DO
- Add adapter for Jwt and Paseto tokens



directorio internal
El compilador no permite que nadie importe los paquetes definidos en internal.

Es una buena manera para reguardar el codigo que no quiero que otro importe, por lo tanto ni modifique ni nada.

Esta bueno poner en este paquete cosas que sean comunes (como commmons o utils que se usa en otros lenguajes). Lo que sea comun a todo el projecto esta bien ponerlo aqui.

O sea, no se puede ver el codigo, pero no puedo usarlo. 

Lo que esta internal es algo que son cosas que estan dentro de la aplicacion y no queres que otros las usen. 

Se puede ver, pero no usar.
















# Backend - Patients

This is a PoC for a patients API

## API endpoints


| Method | URL                             | Description                       |
|--------|---------------------------------|-----------------------------------|
| GET    | /api/v1/patients                | Get all patients                  |
| GET    | /api/v1/patients/:id            | Get one patient                   |
| POST   | /api/v1/patients                | Add one patient                   |

## Directories

cmd: entrypoint, lauch all microservices with goroutines.
apps: microservices backend and frontend.
src: microservices business concepts.



Post

[
    {
        "book": {
            "author": {
                "firstname": "Gabriel",
                "lastname": "Garcia Marquez"
            },
            "title": "100 años de soledad",
            "price": 97,
            "isbn": "0060929790"
        },
        "stock": 31
    },
    {
        "book": {
            "author": {
                "firstname": "Frank",
                "lastname": "Herbert"
            },
            "title": "Dune",
            "price": 53.79,
            "isbn": "0340960191"
        },
        "stock": 12
    },
    {
        "book": {
            "author": {
                "firstname": "Isaac",
                "lastname": "Asimov"
            },
            "title": "Fundation",
            "price": 28.5,
            "isbn": "0-553-29335-4"
        },
        "stock": 41
    }
]

### Ports

In one hand, we have the ports which are interfaces that define how the communication between an actor and the core has to be done. Depending on the actor, the ports has different nature:

- Ports for driver actors, define the set of actions that the core provides and expose to the outside. Each action generally correspond with a specific case of use.

- Ports for driven actors, define the set of actions that the actor has to implement.


### Adapters

In the other hand, we have the adapters that are responsible of the transformation between a request from the actor to the core, and vice versa. This is necessary, because as we said earlier the actors and the core “speaks” different languages.

- An adapter for a driver port, transforms a specific technology request into a call on a core service.

- An adapter for a driven port, transforms a technology agnostic request from the core into an a specific technology request on the actor.


Dependency Injection
After the implementation is done, then it is necessary to connect, somehow, the adapters to the corresponding ports. This could be done when the application starts and it allow us to decide which adapter has to be connected in each port, this is what we call “Dependency injection”. For example, if we want to save data into a mysql database, then we just have to plug an adapter for a mysql database into the corresponding port or if we want to save data in memory (for testing) we need to plug an “in memory database” adapter into that port.







Orders
{
    "client": {
        "firstname": "Homero",
        "lastname": "Simpson"
    },
    "date": "2021-11-29T09:55:34.854041217-03:00",
    "quantity": 51,
    "books": {
        book[]
        "author": {
            "firstname": "Isaac",
            "lastname": "Asimov"
        },
        "title": "Fundation",
        "price": 28.5,
        "isbn": "0-553-29335-4"
    }
}

{
    "client": {
        "firstname": "Juan",
        "lastname": "Perez"
    },
    "date": "2021-11-30T10:40:54.412167134-03:00",
    "details": [
        {
            "books": {
                "author": {
                    "firstname": "Isaac",
                    "lastname": "Asimov"
                },
                "title": "Fundation",
                "price": 28.5,
                "isbn": "0-553-29335-4"
            },
            "quantity": 1
        },
        {
            "books": {
                "author": {
                    "firstname": "Stanislaw",
                    "lastname": "Lem"
                },
                "title": "Solaris",
                "price": 65.2,
                "isbn": "0156027607"
            },
            "quantity": 42
        }
    ]
}


Inventory

{
    "book": {
        "author": {
            "firstname": "Gabriel",
            "lastname": "Garcia Marquez"
        },
        "title": "100 años de soledad",
        "isbn": "0060929790,
        "price": 97.00,

    },
    "stock": 31
}

{
    "book": {
        "author": {
            "firstname": "Frank",
            "lastname": "Herbert"
        },
        "title": "Dune",
        "price": 53.79,
        "isbn": "0340960191"
    },
    "stock": 12
}

{
    "book": {
        "author": {
            "firstname": "Isaac",
            "lastname": "Asimov"
        },
        "title": "Fundation",
        "price": 28.5,
        "isbn": "0-553-29335-4"
    },
    "stock": 41
}



{
    "username":"toribio.gato",
    "password":"12345",
    "email":"tori@gmail.com"
}

En service se hace todo cq cuestion logica calculos o lo sea

repository interactua con la base de datos, o con los elementos de consulta, recibe todo procesado desde service


main->app->