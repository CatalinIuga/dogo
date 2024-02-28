package main

import (
	"context"
	"log"
	"os/exec"

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

func openBrowser(url string) {
	var err error = exec.Command("cmd.exe", "/c", "start", url).Start()

	// switch runtime.GOOS {
	// case "linux":
	// 	err = exec.Command("xdg-open", url).Start()
	// case "windows":
	// 	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	// case "darwin":
	// 	err = exec.Command("open", url).Start()
	// default:
	// 	err = fmt.Errorf("unsupported platform")
	// }
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	dockerService, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer dockerService.Close()

	containers, err := dockerService.ContainerList(ctx, container.ListOptions{
		All: true,
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
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

	for _, container := range containerList {
		log.Println(container)
	}

}
