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

### Installation and Running

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/go-isabella-api.git
   cd go-isabella-api
   ```

2. **Tidy the dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the API application:**
   ```bash
   go run main.go
   ```
   The API will be available at `http://localhost:8080`.

## Docker Exporter

This project uses a separate background process, `docker_exporter`, to collect Docker container information and save it to a local JSON file (`docker_data.json`). The main API then reads this file, completely decoupling the API from direct Docker daemon interaction and enhancing security.

### Running the Docker Exporter

To run the `docker_exporter` in the background, execute the following command in a separate terminal:

```bash
go run cmd/docker_exporter/docker_exporter.go &
```

This exporter needs access to the Docker daemon. Ensure that the user running this command has permissions to access `/var/run/docker.sock`.

## API Endpoints

### Health Check

- **GET /**: Returns the health status of the API.
  - **Response:**
    ```json
    {
      "status": "healthy",
      "message": "Docker Exporter API"
    }
    ```

### Metrics Endpoint

- **GET /metrics**: Returns a list of all Docker container information from `docker_data.json`.

# Future
- [ ] Dockerize
- [ ] OpenAPI spec
- [ ] Gin for better routing
- [ ] Security Audit
- [ ] More data on network traffic, hard drive array health/size, system internals.

