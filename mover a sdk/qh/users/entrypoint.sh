#!/bin/sh

# Load environment variables from the .env file
if [ -f .env ]; then
  export $(cat .env | grep -v ^# | xargs)
fi

# Function to log messages
log() {
  echo "ENTRYPOINT: $1"
}

# Function to build the server binary
buildServer() {
  log "Building server binary"
  go build -gcflags "all=-N -l" -o "/app/bin/$BINARY_NAME" "/app/cmd/api"
}

# Function to run the server
runServer() {
  log "Run server"

  log "Killing old server"
  pkill -f dlv || true
  pkill -f "/app/bin/$BINARY_NAME" || true

  log "Run in debug mode"
  dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec "/app/bin/$BINARY_NAME" &
}

# Function to rebuild and rerun the server
rerunServer() {
  log "Rerun server"
  buildServer
  runServer
}

# Function to monitor file changes and trigger server restart
liveReloading() {
  log "Run liveReloading"
  inotifywait -e modify,delete,move -m -r --format '%w%f' --exclude '.*(\.tmp|\.swp)$' /app | (
    while read file; do
      if [[ "$file" == *.go ]]; then
        log "File $file changed. Reloading..."
        rerunServer
      fi
    done
  )
}

# Function to initialize the file change logger
initializeFileChangeLogger() {
  echo "" > /tmp/filechanges.log
  tail -f /tmp/filechanges.log &
}

# Main function to orchestrate the process
main() {
  initializeFileChangeLogger
  buildServer
  runServer
  liveReloading
}

# Call the main function to start the process
main
