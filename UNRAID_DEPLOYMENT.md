# Deploying to Unraid Server

This guide explains how to deploy the Go Isabella API to your Unraid server.

## Prerequisites

1. Unraid server with Docker enabled
2. Community Applications plugin installed
3. Docker Compose Manager plugin installed (see instructions below)

## Option 1: Using Pre-built Images (Recommended)

This is the easiest method - uses images from Docker Hub.

### Step 1: Build and Push Images to Docker Hub

First, build and push your images from your development machine:

```bash
# Build the images
make docker-build-api
make docker-build-exporter

# Login to Docker Hub
docker login

# Push the images
make docker-push-all
```

### Step 2: Install Docker Compose Manager on Unraid

1. Open your Unraid web interface
2. Go to the **Apps** tab
3. Search for "Docker Compose Manager"
4. Click **Install**

### Step 3: Create the Stack in Unraid

1. Go to the **Docker** tab in Unraid
2. Scroll down to the **Compose** section
3. Click **Add New Stack**
4. Name it `isabella-api` (or any name you prefer)
5. Click the gear icon next to your stack
6. Select **Edit Stack**
7. In the **Compose File** section, paste the contents of `docker-compose.unraid.yml`
8. Click **Save Changes**

### Step 4: Deploy the Stack

1. Back in the **Docker** tab, under **Compose** section
2. Find your `isabella-api` stack
3. Click **Compose Up** to start the services

The API will be available at `http://your-unraid-ip:8191`

## Option 2: Building on Unraid (Advanced)

If you want to build the images directly on your Unraid server:

### Step 1: Transfer Files to Unraid

You have a few options:

**Option A: Using SMB/NFS Share**
1. Copy your project files to an Unraid share (e.g., `/mnt/user/appdata/isabella-api/`)
2. Access the files from Unraid

**Option B: Using Git**
1. SSH into your Unraid server
2. Navigate to your appdata directory: `cd /mnt/user/appdata/`
3. Clone your repository: `git clone <your-repo-url> isabella-api`
4. Navigate into the directory: `cd isabella-api`

### Step 2: Install Docker Compose Manager

Same as Option 1, Step 2.

### Step 3: Create Stack with Build Context

1. Go to the **Docker** tab → **Compose** section
2. Click **Add New Stack**
3. Name it `isabella-api`
4. Click the gear icon → **Edit Stack**
5. In the **Compose File** section, paste the contents of `docker-compose.yml`
6. **Important**: Update the build context paths if your files are in a different location:
   ```yaml
   build:
     context: /mnt/user/appdata/isabella-api
     dockerfile: Dockerfile.api
   ```
7. Click **Save Changes**

### Step 4: Deploy

1. Click **Compose Up** to build and start the services

## Updating the Stack

When you need to update:

**For Pre-built Images:**
1. Build and push new images from your dev machine: `make docker-push-all`
2. In Unraid, go to your stack → **Compose Down**
3. Click **Compose Up** to pull and start the new images

**For Building on Unraid:**
1. Update your source files on Unraid
2. In Unraid, go to your stack → **Compose Down**
3. Click **Compose Up** to rebuild and start

## Managing the Stack

- **View Logs**: Click the stack name → **View Logs**
- **Stop Services**: Click **Compose Down**
- **Restart Services**: Click **Compose Down**, then **Compose Up**
- **Edit Configuration**: Click gear icon → **Edit Stack**

## Troubleshooting

### Port Already in Use
The default port is set to 8191 to avoid conflicts. If you need to change it, edit the stack and modify:
```yaml
ports:
  - "8191:8080"  # Change first number (8191) to an available port
```
Note: The second number (8080) is the container's internal port and should remain 8080.

### Docker Socket Permission Issues
The exporter needs access to Docker socket. This should work by default, but if you have issues:
- Ensure the exporter container has access to `/var/run/docker.sock`
- Check Unraid's Docker settings

### Images Not Found
If using pre-built images and getting "image not found":
- Verify images are pushed to Docker Hub: `docker pull racecar246/isabella-api:latest`
- Check you're logged into Docker Hub if using private images
- Verify the image names match in `docker-compose.unraid.yml`

## Alternative: Manual Container Setup

If you prefer not to use Docker Compose Manager, you can add containers manually:

1. Go to **Docker** tab → **Add Container**
2. Add the API container:
   - Repository: `racecar246/isabella-api:latest`
   - Name: `isabella-api`
   - Port: `8191:8080`
   - Add volume: `isabella-cache:/app/cache`
3. Add the Exporter container:
   - Repository: `racecar246/isabella-exporter:latest`
   - Name: `isabella-exporter`
   - Add volume: `/var/run/docker.sock:/var/run/docker.sock:ro`
   - Add volume: `isabella-cache:/app/cache`
4. Create a custom network `isabella-network` and add both containers to it

However, Docker Compose is recommended as it's easier to manage.

