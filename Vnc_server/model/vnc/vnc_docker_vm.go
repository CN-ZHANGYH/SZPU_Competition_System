package model

type DockerVmInfo struct {
	ContainerName  string
	ContainerImage string
	VncPort        string
	WeBasePort     string
	FrontPort      string
	ChannelPort    string
	SshPort        string
}
