#!/bin/bash

# Detener y eliminar todos los contenedores
docker stop "$(docker ps -aq)"
docker rm "$(docker ps -aq)"

# Eliminar todas las imágenes
docker rmi "$(docker images -q)"

# Eliminar todos los volúmenes
docker volume rm "$(docker volume ls -q)"

# Eliminar todas las redes
docker network rm "$(docker network ls -q)"

# Limpiar el sistema de Docker
docker system prune -a --volumes -f

echo "Docker cleanup completed."
