#!/bin/bash

# TODO: no funciona el script, sin embargo si puedo crear la base de datos si sigo los paso de la documentacion de pgAdmin.
# FIXME: ssadsadada
# BUG: asdasdsad
# ?: dasdadsadad
# Función para cargar las variables de entorno desde el archivo .env
load_env() {
  export $(grep -v '^#' ../.env | xargs)
}

# Llamar a la función para cargar las variables
load_env

# Variables de entorno cargadas desde el .env
DB_HOST=${DEV_DB_HOST:-postgres}
DB_PORT=${DEV_DB_HOST_PORT:-5432}
DB_NAME=${DEV_DB_DATABASE:-dev_events_db}
DB_USER=${DEV_DB_USERNAME:-postgres}
DB_PASSWORD=${DEV_DB_USER_PASSWORD:-root}
DB_ROOT_PASSWORD=${DEV_DB_ROOT_PASSWORD:-rootpassword}
DB_TABLE=${DEV_DB_TABLE:-events}
COMPOSE_PROJECT_NAME=${COMPOSE_PROJECT_NAME:-my_project}

# Obtener el nombre de la red de Docker Compose
NETWORK_NAME=$(docker network ls --filter name=${COMPOSE_PROJECT_NAME}_default --format "{{.Name}}")

if [ -z "$NETWORK_NAME" ]; then
  echo "Docker network not found. Please make sure Docker Compose is up and running."
  exit 1
fi

# Esperar a que el contenedor de PostgreSQL esté listo para aceptar conexiones
until docker run --rm --network=${NETWORK_NAM} -e PGPASSWORD=$DB_ROOT_PASSWORD postgres:16 psql -U postgres -h $DB_HOST -p $DB_PORT -c "\q"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

# Crear la base de datos
docker run --rm --network=$NETWORK_NAME -e PGPASSWORD=$DB_ROOT_PASSWORD postgres:16 psql -U postgres -h $DB_HOST -p $DB_PORT -c "CREATE DATABASE $DB_NAME;"

# Crear el usuario y otorgar privilegios
docker run --rm --network=$NETWORK_NAME -e PGPASSWORD=$DB_ROOT_PASSWORD postgres:16 psql -U postgres -h $DB_HOST -p $DB_PORT -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';"
docker run --rm --network=$NETWORK_NAME -e PGPASSWORD=$DB_ROOT_PASSWORD postgres:16 psql -U postgres -h $DB_HOST -p $DB_PORT -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"

# Crear la tabla
docker run --rm --network=$NETWORK_NAME -e PGPASSWORD=$DB_PASSWORD postgres:16 psql -U $DB_USER -h $DB_HOST -p $DB_PORT -d $DB_NAME -c "CREATE TABLE $DB_TABLE (
  id UUID PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  location VARCHAR(255),
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP,
  category VARCHAR(50),
  creator_id UUID NOT NULL,
  is_public BOOLEAN NOT NULL DEFAULT true,
  is_recurring BOOLEAN NOT NULL DEFAULT false,
  series_id UUID,
  status VARCHAR(50) NOT NULL
);"

echo "Database and table created successfully"
