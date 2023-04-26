# Docker

## Requerimientos

Debe tener instalado localmente:

- docker
- docker-compose

## Ejecución

Ejecutar el proyecto: $ sudo make up

Detener el proyecto: $ sudo make down

## Configuración

Para configurar MySQL, usar 'mysql-repo' como servidor en el código de golang.

## PhpMyAdmin

PhpMyAdmin es un administrador de bases de datos para MySQL.
Para ejecutarlo:

<http://localhost:9090>
usuario: root
contraseña: secret

sudo make down; docker rm -f $(docker ps -a -q); docker volume rm $(docker volume ls -q)
