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

	"go-isabella-api/pkg/types" // Import the new types package
)

const (
	dockerOutputFile = "docker_data.json" // Specific file for Docker data
	interval         = 5 * time.Second    // How often to collect data
)

func getDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithVersion("1.41"))
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}
	return cli, nil
}

func collectDockerInfo(cli *client.Client) ([]types.Container, error) { // Use types.Container
	containers, err := cli.ContainerList(context.Background(), client.ContainerListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	containerList := make([]types.Container, 0, len(containers.Items)) // Use types.Container
	for _, c := range containers.Items { // Use types.Container
		containerList = append(containerList, types.Container{ // Use types.Container
			ID:      c.ID,
			Name:    strings.Join(c.Names, ", "),
			Image:   c.Image,
			Status:  c.Status,
			Created: c.Created,
		})
	}
	return containerList, nil
}

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

		data := types.DockerInfo{ // Use types.DockerInfo
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
