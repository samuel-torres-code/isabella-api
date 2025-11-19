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

This project uses a `Makefile` to streamline building the Go executables and Docker images, and managing the Docker containers.

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

4.  **Run with Docker Compose (Recommended):**
    The simplest way to run the application is using Docker Compose:
    ```bash
    make compose-up-build
    ```
    This will build the images and start both services. The API will be available at `http://localhost:8080`.
    
    Other useful Docker Compose commands:
    - `make compose-up` - Start services (assumes images are already built)
    - `make compose-down` - Stop and remove services
    - `make compose-logs` - View logs from all services
    - `make compose-logs-api` - View logs from API service only
    - `make compose-logs-exporter` - View logs from exporter service only
    - `make compose-restart` - Restart all services
    - `make compose-ps` - Show status of services
    - `make compose-clean` - Stop services and remove volumes
    
    Or use `docker-compose` directly:
    ```bash
    docker-compose up -d          # Start services
    docker-compose down           # Stop services
    docker-compose logs -f        # View logs
    docker-compose ps             # Show status
    ```

5.  **Alternative: Manual Docker Setup (Legacy):**
    If you prefer the manual Docker approach:
    
    **Docker Setup (Network and Volume):**
    Before running the Docker containers, ensure the shared network and volume are created:
    ```bash
    make docker-network
    make docker-volume
    ```
    These commands create the `isabella-network` and `isabella-cache` volume, respectively.

    **Run the Docker Exporter container:**
    ```bash
    make docker-run-exporter
    ```
    This will start the `isabella-exporter` container. This exporter needs access to the Docker daemon, so it mounts `/var/run/docker.sock`. Ensure that the user running this command has permissions to access `/var/run/docker.sock`. It also uses the `isabella-network` and `isabella-cache` volume.

    **Run the API application container:**
    ```bash
    make docker-run-api
    ```
    This will start the `isabella-api` container, exposing it on port `8080`. The API will be available at `http://localhost:8080`. It also uses the `isabella-network` and `isabella-cache` volume.

    **Run both Exporter and API containers:**
    ```bash
    make docker-run
    ```
    This command will start both the `isabella-exporter` and `isabella-api` containers.

    **Stop all running containers:**
    ```bash
    make docker-stop
    ```

    **View container logs:**
    ```bash
    make docker-logs-api
    make docker-logs-exporter
    ```

    **Clean up:**
    ```bash
    make clean
    ```
    This command stops and removes containers, images, executables, network, and volume.

    **Clean cache volume only:**
    ```bash
    make clean-cache
    ```

    **Restart all containers:**
    ```bash
    make restart
    ```

    **Check status of Docker components:**
    ```bash
    make status
    ```

### Publishing to Docker Hub

To publish your Docker images to Docker Hub, follow these steps:

1.  **Log in to Docker Hub:**
    ```bash
    docker login
    ```
    Enter your Docker Hub username and password when prompted.

2.  **Tag your Docker images:**
    ```bash
    make docker-tag-api
    make docker-tag-exporter
    ```
    These commands will tag your locally built images with your Docker Hub username (e.g., `racecar246/isabella-api`).

3.  **Push your Docker images to Docker Hub:**
    ```bash
    make docker-push-api
    make docker-push-exporter
    ```
    Alternatively, push both images with:
    ```bash
    make docker-push-all
    ```

### Deploying to Unraid

For detailed instructions on deploying this application to your Unraid server, see [UNRAID_DEPLOYMENT.md](UNRAID_DEPLOYMENT.md).

Quick summary:
1. Build and push images: `make docker-push-all`
2. Install Docker Compose Manager plugin on Unraid
3. Create a new stack using `docker-compose.unraid.yml`
4. Deploy with **Compose Up**

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

