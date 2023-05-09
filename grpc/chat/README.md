# API gRPC bidireccional para un chat

## Definir los servicios y mensajes en un archivo .proto

Correr: $ ./generate

# API gRPC

## Estructura de directorios y arquitectura hexagonal

chat/
├── cmd/
│   └── grpc/
│       └── main.go
├── internal/
│   ├── application/
│   │   └── chat/
│   │       ├── chat.go
│   │       └── chat_test.go
│   ├── domain/
│   │   ├── chat.go
│   │   └── chat_test.go
│   └── infrastructure/
│       ├── grpc/
│       │   ├── chat.go
│       │   └── chat_test.go
│       └── persistence/
│           ├── mongo/
│           │   ├── chat.go
│           │   └── chat_test.go
│           └── persistence.go
└── pkg/
    ├── chat/
    │   ├── chat.pb.go
    │   └── chat_grpc.pb.go
    └── config/
        ├── config.go
        └── config_test.go

## Arquitectura Hexagonal

En esta estructura, domain contiene las entidades y reglas del negocio, application implementa casos de uso, e infrastructure se encarga de detalles técnicos como la comunicación gRPC y la persistencia en MongoDB.

## Uso de MongoDB para persistencia

Para usar MongoDB como base de datos en tu aplicación, primero necesitarás instalar el paquete oficial de MongoDB para Golang:

$ go get go.mongodb.org/mongo-driver/mongo

Luego, en infrastructure/persistence/mongo/chat.go, implementa las funciones de persistencia del chat utilizando el paquete de MongoDB. Por ejemplo:



El directorio pkg (abreviatura de "package") es una convención en proyectos de Golang utilizada para almacenar paquetes de código reutilizable que no son específicos del dominio de la aplicación y pueden ser compartidos entre múltiples proyectos o aplicaciones. En general, estos paquetes son independientes de la lógica de negocio y proporcionan funcionalidades que pueden ser útiles en una amplia variedad de contextos.

En el caso de tu proyecto de chat, los archivos chat.pb.go y chat_grpc.pb.go se colocan en el directorio pkg porque se generan automáticamente a partir de la definición del archivo .proto y no están vinculados a la lógica de negocio de tu aplicación. Estos archivos contienen el código necesario para la comunicación gRPC y pueden ser importados por otros proyectos que necesiten interactuar con tu servicio de chat a través de gRPC.

Además, en tu proyecto, también se incluye el paquete config en el directorio pkg. Este paquete es responsable de cargar la configuración de la aplicación desde las variables de entorno o un archivo de configuración y podría ser utilizado en otros proyectos que requieran un enfoque similar para manejar la configuración.

En resumen, el directorio pkg debe contener paquetes de código reutilizable y genérico que no estén atados a la lógica de negocio específica de tu aplicación. Estos paquetes pueden ser compartidos e importados por otros proyectos o aplicaciones, y suelen ser independientes de la arquitectura de la aplicación.



En el contexto de la arquitectura hexagonal, el cliente gRPC no pertenece directamente a ninguna de las capas principales (dominio, aplicación e infraestructura). Esto se debe a que el cliente gRPC no es una parte esencial de la lógica de negocio de la aplicación de chat. En cambio, el cliente gRPC es un consumidor externo de la API que se comunica con el servidor gRPC.

Por lo general, se considera una buena práctica mantener el código del cliente separado del proyecto principal, especialmente si el cliente es un ejemplo o una herramienta de prueba. En el caso de tu proyecto, puedes crear un directorio separado en la raíz del proyecto para alojar el código del cliente gRPC.

Una posible estructura de directorios para tu proyecto podría ser:

css
Copy code
.
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── internal
│   ├── application
│   ├── domain
│   └── infrastructure
└── pkg


En esta estructura, el cliente gRPC se encuentra en el directorio cmd/grpc-client, mientras que el servidor gRPC y la lógica de la aplicación se encuentran en el directorio cmd/server y en las carpetas de la arquitectura hexagonal (internal/application, internal/domain, e internal/infrastructure). Esta estructura mantiene el cliente gRPC separado de la lógica principal de la aplicación, facilitando su uso como herramienta de prueba o ejemplo de uso de la API de chat.


En gRPC, tanto los clientes como los servidores pueden tener comunicación bidireccional. La diferencia entre un cliente y un servidor en gRPC radica en cómo se inicia y se maneja la comunicación:

Servidor gRPC: El servidor gRPC se encarga de escuchar y aceptar conexiones entrantes de los clientes gRPC. El servidor también implementa los servicios y métodos definidos en el archivo .proto. Cuando se define un servicio bidireccional en el archivo .proto, el servidor es responsable de manejar los streams de entrada y salida para recibir y enviar mensajes a los clientes. En el caso de tu aplicación de chat, el servidor gRPC se encuentra en internal/infrastructure/grpc/chat.go y maneja el método StreamChat para la comunicación bidireccional.

Cliente gRPC: Un cliente gRPC es responsable de establecer una conexión con un servidor gRPC y realizar solicitudes a los métodos definidos en el archivo .proto. Cuando se utiliza un servicio bidireccional, el cliente también debe manejar los streams de entrada y salida para enviar y recibir mensajes del servidor. En general, el cliente gRPC se utiliza para interactuar con el servidor y consumir la API que expone.

Aunque ambos, cliente y servidor, pueden manejar la comunicación bidireccional, sus roles son distintos. El servidor gRPC implementa y expone los servicios y métodos, mientras que el cliente gRPC los consume y se comunica con el servidor para realizar acciones.

En resumen, en el contexto de gRPC, la diferencia entre un cliente y un servidor radica en sus responsabilidades y cómo manejan la comunicación. El servidor gRPC escucha y acepta conexiones, implementa los servicios y métodos, y maneja los streams de entrada y salida. Por otro lado, el cliente gRPC establece conexiones con el servidor, realiza solicitudes a los servicios y métodos expuestos, y también maneja los streams de entrada y salida para enviar y recibir mensajes del servidor.