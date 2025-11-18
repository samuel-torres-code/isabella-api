# Go Isabella API

This is a lightweight API for my server.

# Functionality

## Description

The Go Isabella API provides endpoints for...
- Live docker container information

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

3. **Run the application:**
   ```bash
   go run main.go
   ```
   The API will be available at `http://localhost:8080`.

## API Endpoints

### Health Check

- **GET /**: Returns the health status of the API.
  - **Response:**
    ```json
    {
      "status": "healthy",
      "message": "Docker Container API"
    }
    ```

### Container Endpoints

- **GET /containers**: Returns a list of all Docker containers.
  - **Response:**
    ```json
    {
      "containers": [
        {
          "id": "container_id",
          "name": "container_name",
          "image": "image_name",
          "state": "running",
          "status": "Up 2 hours",
          "created": 1678886400
        }
      ]
    }
    ```

- **GET /containers/{container_id}**: Returns details for a specific container.
  - **Response:**
    ```json
    {
      "id": "container_id",
      "name": "container_name",
      "image": "image_name",
      "state": "running",
      "created": 1678886400
    }
    ```


# Future
- [ ] Dockerize
- [ ] OpenAPI spec
- [ ] Gin for better routing
- [ ] Security Audit
- [ ] More data on network traffic, hard drive array health/size, system internals.

