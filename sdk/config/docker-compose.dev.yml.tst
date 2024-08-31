version: "3.8"

services:
  greeter-client-api:
    container_name: "greeter-client"
    build:
      context: ..
      dockerfile: config/Dockerfile.dev # Ruta al Dockerfile del cliente
    image: "greeter-client:${APP_VERSION:-1.0}"
    ports:
    - "${DELVE_PORT_SERVER:-2345}:${DELVE_PORT_SERVER:-2345}"
    environment:
      # - GRPC_SERVER_HOST=greeter-server # Nombre del servicio del servidor gRPC
      # - GRPC_SERVER_HOST=${GRPC_SERVER_HOST:-localhost} # Usa localhost cuando estás en modo host
      # - GRPC_SERVER_HOST=172.17.0.1
      - GRPC_SERVER_PORT=${GRPC_SERVER_PORT:-50051}
      - MAIN_DIR=/app/cmd/examples/greeter-client/main.go
      - APP_NAME=greeter-client-api
      - APP_VERSION=${APP_VERSION:-1.0}
      - DEBUG=${DEBUG:-true}
    # networks:
    #   - app-network
    network_mode: "host"  # Usa la red del host
    volumes:
      - type: bind
        source: ..
        target: /app
    restart: on-failure
    profiles:
      - greeter-client
      - greeter-service

networks:
  app-network:
    driver: bridge
