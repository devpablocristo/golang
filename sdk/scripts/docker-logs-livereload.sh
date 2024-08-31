#!/bin/bash

# Mata cualquier proceso existente que coincida con el comando de logs de Docker Compose
pkill -f "docker compose -f config/docker-compose.dev.yml logs -f greeter-client-api"

# Espera un momento para asegurarse de que el proceso ha terminado
sleep 1

# Inicia el comando de logs de Docker Compose
docker compose -f config/docker-compose.dev.yml logs -f greeter-client-api
