#!/bin/sh

# shellcheck disable=SC2154  # Desactivar aviso de shellcheck para APP_NAME
# shellcheck source=./config/.env disable=SC1091 # Desactivar aviso de shellcheck de archivo no especificado

# Load environment variables from the ./config/.env file without overwriting existing ones
loadEnv() {
  if [ -f ./config/.env ]; then
    log "Loading environment variables from ./config/.env"
    # Read each line in .env file
    while IFS= read -r line || [ -n "$line" ]; do
      # Ignore empty lines and comments
      if [ -n "$line" ] && [ "${line#\#}" = "$line" ]; then
        VAR_NAME=$(echo "$line" | cut -d '=' -f1)
        VAR_VALUE=$(echo "$line" | cut -d '=' -f2-)
        # Check if variable is already set
        if [ -z "$(printenv "$VAR_NAME")" ]; then
          export "$VAR_NAME=$VAR_VALUE"
        else
          log "Variable $VAR_NAME is already set to $(printenv "$VAR_NAME"), not overwriting"
        fi
      fi
    done < ./config/.env
  else
    echo "ERROR: ./config/.env file not found in the parent directory. Please create ./config/.env with the necessary environment variables"
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
    log "ERROR: APP_NAME is not set. Please check ./config/.env file"
    exit 1
  fi

  if [ -z "${DEBUG}" ]; then
    log "ERROR: DEBUG is not set. Please check ./config/.env file"
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

runServer() {
  log "Running service"

  # Kill any existing server processes
  log "Killing old processes"
  pkill -f dlv || true
  pkill -f "/app/tmp/${APP_NAME}" || true

  if [ "${DEBUG}" = "true" ]; then
    log "Running in debug mode with Air and Delve"
    
    # Log APP_ROLE to verify correct loading
    log "App Role is set to: ${APP_ROLE}"

    # Check if the role is client or server
    if [ "$APP_ROLE" = "client" ]; then
      log "Starting client"
      air -c "$AIR_CONFIG"
      # air -c "$AIR_CONFIG_CLIENT"
    elif [ "$APP_ROLE" = "server" ]; then
      log "Starting server"
      air -c "$AIR_CONFIG"
      # air -c "$AIR_CONFIG_SERVER"
    else
      log "ERROR: APP_ROLE is not set correctly. Must be 'client' or 'server'"
      exit 1
    fi
  else
    log "Running in production mode"
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
