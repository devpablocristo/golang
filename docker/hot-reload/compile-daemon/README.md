# Documentación: Configuración de NoSQL Workbench for Amazon DynamoDB y AWS CLI

## Configuración de NoSQL Workbench for Amazon DynamoDB

Para ejecutar NoSQL Workbench for Amazon DynamoDB, sigue los siguientes pasos:

1. Abre una terminal y navega al directorio que contiene el ejecutable de NoSQL Workbench. Puedes utilizar el siguiente comando:

    ```bash
    cd /home/pablo/DynamoDBWorkbench
    ```

2. Ejecuta NoSQL Workbench utilizando el siguiente comando:

    ```bash
    ./NoSQL\ Workbench-linux-3.10.0.AppImage
    ```

Este paso te permitirá iniciar NoSQL Workbench y comenzar a trabajar con tus tablas de DynamoDB.

## Instalación de la AWS CLI

Para resolver el problema de "Unable to locate credentials" al usar la AWS CLI con DynamoDB Local, realiza los siguientes pasos:

1. Descarga la AWS CLI ejecutando el siguiente comando:

    ```bash
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    ```

2. Descomprime el archivo descargado:

    ```bash
    unzip awscliv2.zip
    ```

3. Instala la AWS CLI ejecutando el siguiente comando con privilegios de superusuario:

    ```bash
    sudo ./aws/install
    ```

Este proceso instalará la AWS CLI en tu sistema.

## Configuración de Credenciales Ficticias

Para interactuar con DynamoDB Local usando la AWS CLI, configura credenciales ficticias mediante el comando `aws configure`. Ejecuta el siguiente comando y proporciona los valores ficticios cuando se te solicite:

```bash
aws configure
```

Proporciona los siguientes valores ficticios:

- `AWS Access Key ID`: Cualquier valor, por ejemplo, "fake-access-key".
- `AWS Secret Access Key`: Cualquier valor, por ejemplo, "fake-secret-key".
- `Default region name`: Cualquier valor, por ejemplo, "local".
- `Default output format`: Puedes dejar esto en blanco o seleccionar un formato de salida, como "json".

Estas credenciales ficticias permitirán que la AWS CLI funcione con DynamoDB Local.

## Ejecución de Comandos de AWS CLI para DynamoDB Local

Ahora puedes ejecutar comandos de la AWS CLI para DynamoDB Local sin recibir el mensaje "Unable to locate credentials". Por ejemplo:

### Listar Tablas en DynamoDB Local:

```bash
aws dynamodb list-tables --endpoint-url http://localhost:9876 --region local
```

Este comando listará las tablas en DynamoDB Local en el puerto 9876, utilizando las credenciales ficticias y la región "local".

Recuerda que este enfoque de credenciales ficticias es adecuado para entornos locales de desarrollo o pruebas con DynamoDB Local. En un entorno de producción o en la nube, deberías proporcionar credenciales válidas y seguras.



Por un problema particular de permisos, creo que porque tengo que correr docker como root, tengo que correr estos comandos:
$ sudo chown $USER docker-volumes -R
$ chmod 775 -R docker-volumes

Y recien luego puedo crear las tablas:
$ aws dynamodb create-table \
--table-name Employee \
--attribute-definitions \
    AttributeName=EmployeeID,AttributeType=N \
--key-schema \
    AttributeName=EmployeeID,KeyType=HASH \
--provisioned-throughput \
    ReadCapacityUnits=5,WriteCapacityUnits=5 \
--endpoint-url http://localhost:9876

    
nota: los volumes son etiquetas para espacios q define docker, donde guardara la persistencia.