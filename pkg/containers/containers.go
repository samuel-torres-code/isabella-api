package containers

import (
	"net/http"
	"os"
)

const (
	dockerDataFile = "/app/cache/docker_data.json"
)

// GetContainers handles the request to fetch metrics from the Docker daemon.
func GetContainers(w http.ResponseWriter, r *http.Request) {
	// Read the content of the docker_data.json file
	data, err := os.ReadFile(dockerDataFile)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Docker data file not found. Please ensure docker_exporter is running.", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
