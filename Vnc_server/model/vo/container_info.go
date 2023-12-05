package vo

import "github.com/docker/docker/api/types"

type ContainerInfo struct {
	Id     string
	Host   string
	Image  string
	Status string
	Names  string
	Ports  []types.Port
}
