package service

import (
	"context"
	"fmt"
	"log"

	"dogo/model"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerService struct {
	cli *client.Client
}

func NewDockerService() *DockerService {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	return &DockerService{cli: cli}
}

func (d *DockerService) ListContainers() ([]model.Container, error) {
	ctx := context.Background()
	containers, err := d.cli.ContainerList(ctx, container.ListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}

	var result []model.Container
	for _, container := range containers {
		result = append(result, model.Container{
			ID:      container.ID[:10],
			Image:   container.Image,
			Command: container.Command,
			Created: container.Created,
			Status:  container.Status,
			Ports:   container.Ports,
			Names:   container.Names,
		})
	}
	return result, nil
}

func (d *DockerService) StartContainer(id string) error {
	ctx := context.Background()
	return d.cli.ContainerStart(ctx, id, container.StartOptions{})
}

func (d *DockerService) StopContainer(id string) error {
	ctx := context.Background()
	return d.cli.ContainerStop(ctx, id, container.StopOptions{})
}

func (d *DockerService) RemoveContainer(id string) error {
	ctx := context.Background()
	return d.cli.ContainerRemove(ctx, id, container.RemoveOptions{})
}

func (d *DockerService) InspectContainer(id string) (model.ContainerInspect, error) {
	ctx := context.Background()
	container, err := d.cli.ContainerInspect(ctx, id)
	if err != nil {
		return model.ContainerInspect{}, err
	}
	return model.ContainerInspect{
		ID:      container.ID[:10],
		Image:   container.Config.Image,
		Command: fmt.Sprintf("%s", container.Config.Cmd),
		Created: container.Created,
		Status:  container.State.Status,
	}, nil
}

func (d *DockerService) CreateContainer(image string, name string) (string, error) {
	ctx := context.Background()
	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
		Image: image,
	}, nil, nil, nil, name)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}
