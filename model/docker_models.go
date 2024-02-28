package model

import (
	"github.com/docker/docker/api/types"
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

type ContainerInspect struct {
	ID              string                       `json:"id"`
	Image           string                       `json:"image"`
	Command         string                       `json:"command"`
	Created         string                       `json:"created"`
	Status          string                       `json:"status"`
	Ports           []types.Port                 `json:"ports"`
	Names           []string                     `json:"names"`
	NetworkSettings types.SummaryNetworkSettings `json:"network_settings"`
	RestartCount    int                          `json:"restart_count"`
	Driver          string                       `json:"driver"`
	Platform        string                       `json:"platform"`
	ExecIDs         []string                     `json:"exec_ids"`
	GraphDriver     types.GraphDriverData        `json:"graph_driver"`
	SizeRw          int64                        `json:"size_rw"`
	SizeRootFs      int64                        `json:"size_root_fs"`
}
