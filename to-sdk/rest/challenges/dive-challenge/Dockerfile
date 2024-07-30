FROM golang:1.22.3-alpine3.20

# Install dependencies and Delve for debugging
RUN apk update && \
    apk add --no-cache bash git libc-dev g++ inotify-tools && \
    go install github.com/go-delve/delve/cmd/dlv@latest

# Download the clickhouse-jdbc file
RUN wget -O /opt/clickhouse-jdbc.jar https://repo1.maven.org/maven2/ru/yandex/clickhouse/clickhouse-jdbc/0.3.1/clickhouse-jdbc-0.3.1.jar

# Set the working directory inside the container
WORKDIR /app

# Copy the project files to the container
COPY . .

# Download and verify the project dependencies
RUN go mod download && go mod verify

# Create the bin directory if it doesn't exist
RUN mkdir -p /app/bin

# Copy the entrypoint script and make it executable
COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Expose the necessary ports
EXPOSE 8080 2345

# Define the entrypoint script for the container
CMD ["/entrypoint.sh"]