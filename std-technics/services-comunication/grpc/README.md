Sacar  de a poco el codigo comprimido y actualizar las dependencias.

Es un buen path para enteder esta parte mas compleja de go.


Context
Protobuf
gRPC


para los modulos:

$ go mod init github.com/devpablocristo/golang-examples/grpc   
$ go mod tidy


protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    todo/task/task.proto




REST, gRPC y Graph son paradigmas 

Protobuf (gRPC): Mas rapido y mas eficiente. 




Apache thrift es los mismo que protobuf, pero mas incomodo.

O sea:

Thrift
Json
XML
Protobuf

Son todos simplemente formatos de intercambio de datos!



# Comando

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    book/book.proto
```





Now it is your turn to write code!

In this exercise, your goal is to implement a Sum RPC Unary API in a CalculatorService:

The function takes a Request message that has two integers, and returns a Response that represents the sum of them.
Remember to first implement the service definition in a .proto file, alongside the RPC messages
Implement the Server code first
Test the server code by implementing the Client
Example:

The client will send two numbers (3 and 10) and the server will respond with (13)

Good luck!



Project Setup

1) Protoc Setup
In order to perform code generation, you will need to install protoc  on your computer.

============ MacOSX =============

It is actually very easy, open a command line interface and type brew install protobuf 

============ Ubuntu (Linux) ============

Find the correct protocol buffers version based on your Linux Distro: https://github.com/google/protobuf/releases

Example with x64:

# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
============ Windows ============

Download the windows archive: https://github.com/google/protobuf/releases

Example: https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-win32.zip

Extract all to C:\proto3  

Your directory structure should now be

C:\proto3\bin 

C:\proto3\include 

Finally, add C:\proto3\bin to your PATH:

From the desktop, right-click My Computer and click Properties.
In the System Properties window, click on the Advanced tab
In the Advanced section, click the Environment Variables button.
Finally, in the Environment Variables window (as shown below), highlight the Path variable in the Systems Variable section and click the Edit button. Add or modify the path lines with the paths you wish the computer to access. Each different directory is separated with a semicolon as shown below.

C:\Program Files; C:\Winnt; ...... ; C:\proto3\bin
(you need to add ; C:\proto3\bin  at the end)


2) Setup gRPC for Golang

https://github.com/grpc/grpc-go

Module: "google.golang.org/grpc"

3) Setup Protbuf for Golang

https://github.com/golang/protobuf

Module: "google.golang.org/protobuf"

4) Create folders and files
    Example: 
        bookstorepb: bookstore.proto
        bookstore_server: server.go
        bookstores_client: cliente. go

5) Generate de protobuf files

protoc bookstorepb/bookstore.proto --go_out=plugins=grpc:.

It usefull to create a generate_proto.sh file to reuse this command.













# Guia Basica gRPC

1. En el archivo .proto define the gRPC service and the method request and response types using protocol buffers:
    - To define a service, you specify a named service in your .proto file.
    - Es muy importante.
    - Es lo que se recibe y lo que se devuelve con API.

        ```proto
        // lo que hará el servicio
        service BookstoreInventory {
        ...
        }
        ```

2. NewBook es el input de CreateNewBook y Book es el output.

    ```proto
    message NewBook {
        string title = 1;
        string author = 2;
    }

    message Book {
        string title = 1;
        string author = 2;
        int32 id = 3;
    }

    service BookstoreInventory {
        rpc CreateNewBook(NewBook) returns (Book) {}
    }
    ```

## Comando

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    book/book.proto
```
# gRPC

## 1. Instalar protobuf

```shell
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
foo@bar:~$ go get -u google.golang.org/grpc
```

## 2. Agregar el import en el codigo

```go
import "google.golang.org/grpc"
```

Se define el archivo protobuf, para este ejemplo chat.proto

### *NOTA: debe estar instalado protobuf en el sistema*

```shell
foo@bar:~$ sudo apt install protobuf-compiler
foo@bar:~$ sudo apt install golang-goprotobuf-dev
```

## 3. Crear dir para el paquete chat

```shell
foo@bar:~$ mkdir chat
foo@bar:~$ protoc --go_opt=paths=source_relative --go_out=plugins=grpc:chat chat.proto
```

Genera codigo para gRPC
