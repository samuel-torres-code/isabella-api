package containers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/moby/moby/client"
)

// Container represents a Docker container.
type Container struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Status  string `json:"status"`
	Created int64  `json:"created"`
}

// ContainerList represents a list of Docker containers.
type ContainerList struct {
	Containers []Container `json:"containers"`
}

func getDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithVersion("1.41"))
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// GetContainers handles the request to list all Docker containers.
func GetContainers(w http.ResponseWriter, r *http.Request) {
	cli, err := getDockerClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	containers, err := cli.ContainerList(context.Background(), client.ContainerListOptions{All: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	containerList := make([]Container, 0, len(containers.Items))
	for _, c := range containers.Items {
		containerList = append(containerList, Container{
			ID:      c.ID,
			Name:    strings.Join(c.Names, ", "),
			Image:   c.Image,
			Status:  c.Status,
			Created: c.Created,
		})
	}

	response := ContainerList{Containers: containerList}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetContainer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetContainer"))
}
