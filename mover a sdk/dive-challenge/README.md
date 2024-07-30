# Golang Developer Assignment

This project provides a service that retrieves the Last Traded Price (LTP) of Bitcoin for the following currency pairs:

- BTC/USD
- BTC/CHF
- BTC/EUR

The service is built using Go and provides an API endpoint at `/api/v1/ltp`. The response includes the LTP for the specified currency pairs with time accuracy up to the last minute.

## API Endpoint

### GET /api/v1/ltp

This endpoint retrieves the Last Traded Price of Bitcoin for the specified currency pairs. The request can be made for a single pair or a list of pairs.

#### Request

- **Path:** `/api/v1/ltp`
- **Query Parameters:**
  - `pair`: The currency pair(s) to retrieve the LTP for. Multiple pairs can be specified.

#### Response

The response will be in JSON format with the following structure:

```json
{
  "ltp": [
    {
      "pair": "BTC/CHF",
      "amount": 49000.12
    },
    {
      "pair": "BTC/EUR",
      "amount": 50000.12
    },
    {
      "pair": "BTC/USD",
      "amount": 52000.12
    }
  ]
}
```

## Requirements

- Time accuracy of the data up to the last minute.
- Code hosted in a remote public repository.
- `README.md` includes clear steps to build and run the app.
- Integration tests.
- Dockerized application.
- Documentation.

## Getting Started

### Prerequisites

- Go 1.16+
- Docker
- Docker Compose

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/devpablocristo/dive-challenge.git
   cd dive-challenge
   ```

### Building and Running the Application

#### Using Docker Compose (Recommended)

1. Build and start the Docker containers:
   ```sh
   make docker-up
   ```

2. The application will be available at `http://localhost:8080`.

#### Debugging Inside Docker

To enable debugging inside Docker, ensure that your `Dockerfile` and `docker-compose.debug.yml` are set up to expose the debug port (e.g., 2345) and that Delve is installed in the Docker container.

1. Build and start the Docker containers in debug mode:
   ```sh
   make docker-dup
   ```

2. The application will be available at `http://localhost:8080` and will be waiting for a debugger to attach on port 2345.

### Running Tests

To run the tests, use the following command:
```sh
make test
```

### Linting the Code

To lint the code, use the following command:
```sh
make lint
```

### Makefile Commands

- `make all`: Build and run the project.
- `make build`: Build the project.
- `make run`: Run the project.
- `make test`: Run tests.

- `make docker-build`: Build Docker images for production mode.
- `make docker-up`: Start Docker Compose services in production mode.
- `make docker-down`: Stop and remove Docker Compose services in production mode.

- `make docker-dbuild`: Build Docker images for debug mode.
- `make docker-dup`: Start Docker Compose services in debug mode.
- `make docker-ddown`: Stop and remove Docker Compose services in debug mode.

- `make clean`: Clean binary files.
- `make lint`: Lint the code.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.