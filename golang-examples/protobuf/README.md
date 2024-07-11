# Protobuf comparado con xml y json

## Obtener y configurar protobuf

```shell
foo@bar:~$ go get google.golang.org/protobuf
foo@bar:~$ protoc --go_out=. --go_opt=paths=source_relative *.proto
```

**NOTA:** Para poder hacerlo en el mismo directorio se debe escribir en el archivo .proto.

```proto
option go_package = ".;nombre_paquete";
```
