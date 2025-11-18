// Package main provides a Docker exporter that collects information about running Docker containers
// and writes it to a JSON file at regular intervals.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/moby/moby/client"

	"go-isabella-api/pkg/types"
)

const (
	// dockerOutputFile is the specific file path for Docker data in the shared volume.
	dockerOutputFile = "/app/cache/docker_data.json"
	// interval defines how often to collect Docker data.
	interval = 5 * time.Second
)

// getDockerClient initializes and returns a new Docker client.
func getDockerClient() (*client.Client, error) {
	cli, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithVersion("1.41"))
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}
	return cli, nil
}

// collectDockerInfo collects information about all Docker containers
// and returns them as a slice of types.Container.
func collectDockerInfo(cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), client.ContainerListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	containerList := make([]types.Container, 0, len(containers.Items))
	for _, c := range containers.Items {
		containerList = append(containerList, types.Container{
			ID:      c.ID,
			Name:    strings.Join(c.Names, ", "),
			Image:   c.Image,
			Status:  c.Status,
			Created: c.Created,
		})
	}
	return containerList, nil
}

// writeDockerDataToFile marshals the DockerInfo data to JSON and writes it to the dockerOutputFile.
func writeDockerDataToFile(data types.DockerInfo) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(dockerOutputFile, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write Docker data to file: %w", err)
	}
	return nil
}

// main is the entry point of the Docker exporter.
// It initializes the Docker client, then periodically collects Docker container
// information and writes it to a JSON file.
func main() {
	cli, err := getDockerClient()
	if err != nil {
		log.Fatalf("Error initializing Docker client: %v", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Collecting Docker information...")
		containers, err := collectDockerInfo(cli)
		if err != nil {
			log.Printf("Error collecting Docker info: %v", err)
			continue
		}

		data := types.DockerInfo{
			Containers: containers,
		}

		err = writeDockerDataToFile(data)
		if err != nil {
			log.Printf("Error writing Docker data to file: %v", err)
		} else {
			log.Printf("Successfully wrote Docker info to %s", dockerOutputFile)
		}
	}
}
