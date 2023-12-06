package model

type DockerVmInfo struct {
	ContainerName         string
	ContainerImage        string
	Host_Port_List        []string
	Container_Port_List   []string
	Host_Volume_List      []string
	Container_Volume_List []string
	RestartPolicy         string
}
