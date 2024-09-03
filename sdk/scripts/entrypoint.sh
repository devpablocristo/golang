#!/bin/sh

# shellcheck disable=SC2154  # Desactivar aviso de shellcheck para APP_NAME
# shellcheck source=./config/.env disable=SC1091 # Desactivar aviso de shellcheck de archivo no especificado

# Load environment variables from the ./config/.env file in the parent directory
loadEnv() {
  if [ -f ./config/.env ]; then
    log "Loading environment variables from ./config/.env"
    # Export all variables from .env file
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

# Function to initialize the file change logger
initializeFileChangeLogger() {
  echo "" > /tmp/filechanges.log
  tail -f /tmp/filechanges.log &
}


# Function to run the server with Air and Delve for debugging
runServer() {
  log "Run server"

  # Kill any existing server processes
  log "Killing old server processes"
  pkill -f dlv || true
  pkill -f "/app/tmp/${APP_NAME}" || true

  if [ "${DEBUG}" = "true" ]; then
    log "Run in debug mode with Air and Delve"

    # Start Air with the provided configuration for live reloading
    air -c "$AIR_CONFIG"
  else
    log "Run in production mode with Air"
    
    # Start Air normally for live reloading in production mode
    air -c "$AIR_CONFIG"
  fi
}

# Main function to orchestrate the process
main() {
  log "Starting script"
  log "Current directory: $(pwd)"
  loadEnv
  validateEnv
  initializeFileChangeLogger
  
  # Start server with Air and possibly Delve
  runServer
}

# Call the main function to start the process
main
