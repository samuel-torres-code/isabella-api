.PHONY: all build docker-build-api docker-build-exporter docker-run-api docker-run-exporter docker-stop clean

APP_NAME := isabella
EXPORTER_NAME := docker-exporter

DOCKER_USERNAME := racecar246

# Image names
API_IMAGE_NAME := $(DOCKER_USERNAME)/$(APP_NAME)-api
EXPORTER_IMAGE_NAME := $(DOCKER_USERNAME)/$(APP_NAME)-exporter

# Container names (different from image names)
API_CONTAINER_NAME := isabella-api
EXPORTER_CONTAINER_NAME := isabella-exporter

# Network name
NETWORK_NAME := isabella-network

# Volume name for shared cache
CACHE_VOLUME := isabella-cache

# Cache directory path inside containers
CACHE_PATH := /app/cache

all: build docker-build-api docker-build-exporter

build:
	@echo "Building Go executables..."
	go build -o $(APP_NAME) main.go
	go build -o $(EXPORTER_NAME) cmd/docker_exporter/docker_exporter.go

docker-build-api:
	@echo "Building Docker image for API..."
	docker build -f Dockerfile.api -t $(API_IMAGE_NAME) .

docker-build-exporter:
	@echo "Building Docker image for Exporter..."
	docker build -f Dockerfile.exporter -t $(EXPORTER_IMAGE_NAME) .

docker-tag-api:
	@echo "Tagging API Docker image..."
	docker tag $(API_IMAGE_NAME) $(API_IMAGE_NAME)

docker-tag-exporter:
	@echo "Tagging Exporter Docker image..."
	docker tag $(EXPORTER_IMAGE_NAME) $(EXPORTER_IMAGE_NAME)

docker-push-api:
	@echo "Pushing API Docker image to Docker Hub..."
	docker push $(API_IMAGE_NAME)

docker-push-exporter:
	@echo "Pushing Exporter Docker image to Docker Hub..."
	docker push $(EXPORTER_IMAGE_NAME)

docker-push-all: docker-push-api docker-push-exporter
	@echo "All Docker images pushed to Docker Hub."

# Create shared network if it doesn't exist
docker-network:
	@echo "Creating Docker network..."
	docker network create $(NETWORK_NAME) 2>/dev/null || true

# Create shared volume if it doesn't exist
docker-volume:
	@echo "Creating shared cache volume..."
	docker volume create $(CACHE_VOLUME) 2>/dev/null || true

docker-run-api: docker-network docker-volume
	@echo "Running API Docker container..."
	docker run -d \
		-p 8080:8080 \
		-v $(CACHE_VOLUME):$(CACHE_PATH) \
		--network $(NETWORK_NAME) \
		--name $(API_CONTAINER_NAME) \
		$(API_IMAGE_NAME)

docker-run-exporter: docker-network docker-volume
	@echo "Running Exporter Docker container..."
	docker run -d \
		-v /var/run/docker.sock:/var/run/docker.sock:ro \
		-v $(CACHE_VOLUME):$(CACHE_PATH) \
		--network $(NETWORK_NAME) \
		--name $(EXPORTER_CONTAINER_NAME) \
		$(EXPORTER_IMAGE_NAME)

docker-run: docker-run-exporter docker-run-api
	@echo "All containers started successfully!"
	@echo "API: http://localhost:8080"

docker-stop:
	@echo "Stopping containers..."
	docker stop $(API_CONTAINER_NAME) $(EXPORTER_CONTAINER_NAME) 2>/dev/null || true

docker-logs-api:
	docker logs -f $(API_CONTAINER_NAME)

docker-logs-exporter:
	docker logs -f $(EXPORTER_CONTAINER_NAME)

clean: docker-stop
	@echo "Cleaning up..."
	rm -f $(APP_NAME) $(EXPORTER_NAME)
	docker rm -f $(API_CONTAINER_NAME) $(EXPORTER_CONTAINER_NAME) 2>/dev/null || true
	docker rmi $(API_IMAGE_NAME) $(EXPORTER_IMAGE_NAME) 2>/dev/null || true
	docker volume rm $(CACHE_VOLUME) 2>/dev/null || true
	docker network rm $(NETWORK_NAME) 2>/dev/null || true
	rm -f docker_data.json

# Useful for development
clean-cache:
	@echo "Cleaning cache volume..."
	docker volume rm $(CACHE_VOLUME) 2>/dev/null || true
	docker volume create $(CACHE_VOLUME)

restart: docker-stop docker-run

status:
	@echo "=== Container Status ==="
	@docker ps -a | grep isabella || echo "No containers running"
	@echo ""
	@echo "=== Network Status ==="
	@docker network inspect $(NETWORK_NAME) 2>/dev/null || echo "Network not created"
	@echo ""
	@echo "=== Volume Status ==="
	@docker volume inspect $(CACHE_VOLUME) 2>/dev/null || echo "Volume not created"