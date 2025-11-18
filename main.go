package main

import (
	"fmt"
	"net/http"

	"go-isabella-api/pkg/containers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "healthy", "message": "Docker Container API"}`)
	})

	http.HandleFunc("/containers", containers.GetContainers)
	http.HandleFunc("/containers/", containers.GetContainer)

	http.ListenAndServe(":8080", nil)
}
