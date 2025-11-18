package main

import (
	"fmt"
	"go-isabella-api/pkg/containers"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "healthy", "message": "Docker Metrics API"}`)
	})

	http.HandleFunc("/containers", containers.GetContainers)

	http.ListenAndServe(":8080", nil)
}
