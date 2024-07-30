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
