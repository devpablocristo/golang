#!/bin/sh

# shellcheck disable=SC2154  # Desactivar aviso de shellcheck para APP_NAME
# shellcheck source=./config/.env disable=SC1091 # Desactivar aviso de shellcheck de archivo no especificado

MAIN_DIR="./cmd/examples/greeter-client/"  # Ajusta esta ruta al directorio correcto de tu proyecto


# Load environment variables from the ./config/.env file in the parent directory
loadEnv() {
  if [ -f ./config/.env ]; then
    log "Loading environment variables from ./config/.env"
    # Use `set -a` to export all variables
    set -a
    # shellcheck source=./config/.env
    . ./config/.env
    set +a
  else
    echo "ERROR: ./config/.env file not found in the parent directory. Please create ./config/.env with the necessary environment variables."
    exit 1
  fi
}

# Function to log messages
log() {
  echo "ENTRYPOINT: $1"
}

# Validate essential environment variables
validateEnv() {
  if [ -z "${APP_NAME}" ]; then
    log "ERROR: APP_NAME is not set. Please check ./config/.env file."
    exit 1
  fi

  if [ -z "${DEBUG}" ]; then
    log "ERROR: DEBUG is not set. Please check ./config/.env file."
    exit 1
  fi

  log "Environment variables loaded successfully"
  log "App Name: ${APP_NAME}"
  log "Debug: ${DEBUG}"
}

# Function to build the server binary
buildServer() {
  log "Building server binary"
  mkdir -p ./live
  go build -gcflags "all=-N -l" -buildvcs=false -o "./live/${APP_NAME}" "${MAIN_DIR}"
  # Verify if the binary file has been created and is executable
  if [ -f "/app/bin/${APP_NAME}" ]; then
    log "Binary file created successfully"
    chmod +x "/app/bin/${APP_NAME}"
  else
    log "Failed to create binary file"
    exit 1
  fi
}

# Function to run the server
runServer() {
  log "Run server"

  log "Killing old server"
  pkill -f dlv || true
  pkill -f "/app/bin/${APP_NAME}" || true

  if [ "${DEBUG}" = "true" ]; then
    log "Run in debug mode"
    dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec "/app/bin/${APP_NAME}" &
  else
    log "Run in production mode"
    "/app/bin/${APP_NAME}"
  fi
}

# Function to rebuild and rerun the server
rerunServer() {
  log "Rerun server"
  buildServer
  runServer
}

# Main function to orchestrate the process
main() {
  log "Starting script"
  log "Current directory: $(pwd)"
  loadEnv
  validateEnv
  buildServer
  runServer
}

# If the script is called with the argument "reload", only run the rerunServer function
if [ "$1" = "reload" ]; then
  loadEnv
  validateEnv
  rerunServer
else
  # Call the main function to start the process
  main
fi

# FIXME: no funca! este es lanzado, o deberia ser lanzando por uuna task. pero no funcina, es para  que cuando se vaya al debugger solo con apretar f5 se guarde y se lance el debugger