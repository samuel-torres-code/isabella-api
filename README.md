# Go Isabella API

This is a lightweight API for my server.

# Functionality

## Description

The Go Isabella API provides endpoints for...
- Live docker container information.

## Getting Started

### Prerequisites

- Go 1.18 or later
- Docker installed and running

### Building and Running

This project uses a `Makefile` to streamline building the Go executables and Docker images.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/go-isabella-api.git
    cd go-isabella-api
    ```

2.  **Tidy the dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Build Go executables and Docker images:**
    ```bash
    make all
    ```
    This command will:
    *   Build the `go-isabella-api` executable.
    *   Build the `docker-exporter` executable.
    *   Build the `go-isabella-api-api` Docker image.
    *   Build the `go-isabella-api-exporter` Docker image.

4.  **Run the Docker Exporter container:**
    ```bash
    make docker-run-exporter
    ```
    This will start the `docker-exporter` in a container. This exporter needs access to the Docker daemon, so it mounts `/var/run/docker.sock`. Ensure that the user running this command has permissions to access `/var/run/docker.sock`.

5.  **Run the API application container:**
    ```bash
    make docker-run-api
    ```
    This will start the main API application in a container, exposing it on port `8080`. The API will be available at `http://localhost:8080`.

## API Endpoints

### Health Check

- **GET /**: Returns the health status of the API.
  - **Response:**
    ```json
    {
      "status": "healthy",
      "message": "Isabella API"
    }
    ```

### Containers

- **GET /containers**: Returns a list of all Docker container information from `docker_data.json`.

# Future
- [x] Dockerize
- [ ] OpenAPI spec
- [ ] Gin for better routing
- [ ] Security Audit
- [ ] More data on network traffic, hard drive array health/size, system internals.

