# Backend - Patients

This is a PoC for a patients API

### API endpoints

| Method | URL                             | Description                       |
|--------|---------------------------------|-----------------------------------|
| GET    | /api/v1/patients                | Get all patients                  |
| GET    | /api/v1/patients/:id            | Get one patient                   |
| POST   | /api/v1/patients                | Add one patient                   |

- package cmd: entrypoint - package main and function main
- package internal: cannot import an internal package

## Directories

cmd:
- entrypoint - package main and function main.
api:
- web o http tambien puede llamarse, depende de como sea la infrastructura.
internal:
- Cannot import an internal package
- Buena forma de resguardar el codigo.
- Se puede usar para proteger el codigo.
- Se puede para las cosas comunes.
bd:
- Los modelos de las bases de datos en sql.
migrations:
- Para cambios en la base de datos.
src:
- Donde van los boundaries del negocio.     


Port:

A port is a special type of interface between a machine and the outside world that was created for a specific purpose or protocol. As a result, we can define the port as a technology-independent application programming interface (API) that was created for a specific form of interaction with the application (hence the word “protocol”). It’s entirely up to you how you define this protocol, which is part of what makes this method so appealing. Here are some examples of the various ports you might have:

A port used by your application to access a database
A port used by your application to send out messages such as e-mails or text messages
A port used by human users to access your application
A port used by other systems to access your application
A port used by a particular user group to access your application
A port exposing a particular use case
A port designed for polling clients
A port designed for subscribing clients
A port designed for synchronous communication
A port designed for asynchronous communication
A port designed for a particular type of device

como quiero que mi capa (domain por ejemplo) se comunique con el mundo.

Adapter:
As I already stated, ports are technology agnostic. Even so, you use technology to communicate with the system: a web browser, a mobile device, a dedicated hardware device, a desktop client, and so on. Adapters are very in this situation. An adapter enables interaction through a specific port and with a certain technology. Consider the following scenario:

A REST adapter allows REST clients to interact with the system through some port
A RabbitMQ adapter allows RabbitMQ clients to interact with the system through some port
An SQL adapter allows the system to interact with a database through some port
An adapter allows human users to interact with the system through some port
Numerous adapters for a single port are possible, as well as a single adapter for multiple ports. You can add as many adapters as you want or need to the system without affecting the other adapters, ports, or domain model.