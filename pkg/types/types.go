package types

// Container represents a Docker container.
type Container struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Status  string `json:"status"`
	Created int64  `json:"created"`
}

// DockerInfo represents the collected Docker information.
type DockerInfo struct {
	Containers []Container `json:"containers"`
	
}
