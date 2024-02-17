package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Container struct {
	ID      string       `json:"id"`
	Image   string       `json:"image"`
	Command string       `json:"command"`
	Created int64        `json:"created"`
	Status  string       `json:"status"`
	Ports   []types.Port `json:"ports"`
	Names   []string     `json:"names"`
}

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer cli.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /containers/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		containers, err := cli.ContainerList(ctx, container.ListOptions{
			All: true,
		})
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		var containerList []Container = make([]Container, len(containers))
		for i, container := range containers {
			containerList[i] = Container{
				ID:      container.ID,
				Image:   container.Image,
				Command: container.Command,
				Created: container.Created,
				Status:  container.Status,
				Ports:   container.Ports,
				Names:   container.Names,
			}
		}

		json.NewEncoder(w).Encode(containerList)
	})

	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Error: %s", err)
	}
}
