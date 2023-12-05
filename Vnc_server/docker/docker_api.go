package docker

import (
	"Social_Gin/model/response"
	model "Social_Gin/model/vnc"
	"Social_Gin/model/vo"
	"fmt"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"log"
	"strings"
)

var DOCKER_CLI *client.Client

func InitDockerClient() {

	remoteDockerURL := "tcp://fallingcreams.top:2375"

	cli, err := client.NewClientWithOpts(
		client.WithHost(remoteDockerURL),
		// client.WithVersion("1.41"),
		client.WithAPIVersionNegotiation())
	DOCKER_CLI = cli
	if err != nil {
		log.Fatal("连接Docker失败")
	}
}

// CreateVmContainer 创建Docker的虚拟机
func CreateVmContainer(dockerVmInfo *model.DockerVmInfo, ctx *gin.Context) error {
	// 获取本地已启动的Docker容器，如果要查看全部容器，可以配置types.ContainerListOptions{}参数
	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//for _, container := range containers {
	//	fmt.Printf("容器ID：%s，容器名称：%s,容器状态：%s\n", container.ID[:3], container.Image, container.State)
	//}

	// 配置要启动的容器
	containerConfig := &container.Config{
		Image: dockerVmInfo.ContainerImage, // 指定要使用的镜像
	}

	// 配置端口暴露
	portBindings := nat.PortMap{}
	portBindings["6080/tcp"] = []nat.PortBinding{
		{
			HostIP:   "0.0.0.0",
			HostPort: dockerVmInfo.VncPort,
		},
	}
	portBindings["5000/tcp"] = []nat.PortBinding{
		{
			HostIP:   "0.0.0.0",
			HostPort: dockerVmInfo.WeBasePort,
		},
	}
	portBindings["5002/tcp"] = []nat.PortBinding{
		{
			HostIP:   "0.0.0.0",
			HostPort: dockerVmInfo.FrontPort,
		},
	}
	portBindings["20200/tcp"] = []nat.PortBinding{
		{
			HostIP:   "0.0.0.0",
			HostPort: dockerVmInfo.ChannelPort,
		},
	}
	portBindings["22/tcp"] = []nat.PortBinding{
		{
			HostIP:   "0.0.0.0",
			HostPort: dockerVmInfo.SshPort,
		},
	}

	// 配置容器自动重启
	hostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: "always", // 设置重启策略为"always"，容器将总是自动重启
			// 可选的重启策略：
			// - "no"：无重启策略
			// - "always"：容器总是自动重启
			// - "on-failure"：容器在非零退出状态时重启（默认最多重启3次）
			// - "unless-stopped"：除非手动停止，否则容器总是自动重启
		},

		PortBindings: portBindings,

		//Mounts: []mount.Mount{
		//	{
		//		Type:     mount.TypeBind,
		//		Source:   "./",                    // 宿主机上要挂载的文件夹路径
		//		Target:   "/usr/share/nginx/html", // 容器内挂载的路径
		//		ReadOnly: false,                   // 是否只读
		//	},
		//},
	}
	// 设置容器名称
	name := "szpt_vm-" + dockerVmInfo.ContainerName
	// 创建并启动容器
	resp, err := DOCKER_CLI.ContainerCreate(context.Background(), containerConfig, hostConfig, nil, nil, name)
	if err != nil {
		fmt.Println("容器创建失败:", err)
		return err
	}

	// 启动容器
	DOCKER_CLI.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	//log.Fatal("容器ID:", resp.ID)
	// 关闭Docker客户端连接
	//defer DOCKER_CLI.Close()
	return nil
}

// GetDockerContainerList  查看所有的Docker容器
func GetDockerContainerList(ctx *gin.Context) {
	containerList, err := DOCKER_CLI.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		response.FailWithMessage("查询失败", ctx)
		return
	}

	var containerInfoList []vo.ContainerInfo
	for _, containerInfo := range containerList {
		var container_info vo.ContainerInfo
		container_info.Id = containerInfo.ID[:12]
		container_info.Names = containerInfo.Names[0][1:]
		container_info.Image = containerInfo.Image
		container_info.Status = containerInfo.State
		container_info.Ports = containerInfo.Ports

		containerInfoList = append(containerInfoList, container_info)
	}

	response.OkWithDetailed(containerInfoList, "查询成功", ctx)
	return

}

// GetDockerInfo 获取Docker的信息
func GetDockerInfo(ctx *gin.Context) {
	info, err := DOCKER_CLI.Info(context.Background())
	if err != nil {
		response.FailWithMessage("获取状态失败", ctx)
		return
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("get memory info fail. err： ", err)
	}
	// 获取总内存大小，单位GB
	memTotal := memInfo.Total / 1024 / 1024 / 1024
	// 获取已用内存大小，单位MB
	memUsed := memInfo.Used / 1024 / 1024 / 1024
	// 可用内存大小
	memAva := memInfo.Available / 1024 / 1024 / 1024
	// 内存可用率
	memUsedPercent := memInfo.UsedPercent

	var resultMap = make(map[string]interface{})
	resultMap["id"] = info.ID
	resultMap["system"] = info.OperatingSystem
	resultMap["version"] = info.ServerVersion
	resultMap["status"] = "Running"
	resultMap["memory_total"] = memTotal
	resultMap["memory_ava"] = memAva
	resultMap["memory_used"] = 100 - memUsed
	resultMap["memory_percent"] = memUsedPercent
	response.OkWithDetailed(resultMap, "查询成功", ctx)
	return
}

// RemoveDockerContainer 删除容器
func RemoveDockerContainer(containerId string, ctx *gin.Context) {
	containerList, _ := DOCKER_CLI.ContainerList(context.Background(), types.ContainerListOptions{})
	for _, containerInfo := range containerList {
		if strings.EqualFold(containerInfo.ID[:12], containerId) {
			err := DOCKER_CLI.ContainerRemove(context.Background(), containerInfo.ID, types.ContainerRemoveOptions{
				RemoveVolumes: true,
				Force:         true,
			})
			if err != nil {
				response.FailWithMessage("删除失败", ctx)
				return
			}
			response.OkWithMessage("删除成功", ctx)
			return
		}
		response.FailWithMessage("没有该容器", ctx)
		return
	}
}

func StartDockerContainer(containerId string, ctx *gin.Context) {
	dockerContainer, success := SelectDockerContainer(containerId)
	if success {
		err := DOCKER_CLI.ContainerStart(context.Background(), dockerContainer, types.ContainerStartOptions{})
		if err == nil {
			response.FailWithMessage("当前容器不存在", ctx)
			return
		}
	}
	response.FailWithMessage("当前容器不存在", ctx)
	return
}

func StopDockerContainer(containerId string, ctx *gin.Context) {
	dockerContainer, success := SelectDockerContainer(containerId)
	if success {
		err := DOCKER_CLI.ContainerStop(context.Background(), dockerContainer, container.StopOptions{})
		if err == nil {
			response.FailWithMessage("当前容器不存在", ctx)
			return
		}
	}
	response.FailWithMessage("当前容器不存在", ctx)
	return
}

func SelectDockerContainer(containerId string) (string, bool) {
	containerList, _ := DOCKER_CLI.ContainerList(context.Background(), types.ContainerListOptions{})
	for _, containerInfo := range containerList {
		if strings.EqualFold(containerInfo.ID[:12], containerId) {
			err := DOCKER_CLI.ContainerRemove(context.Background(), containerInfo.ID, types.ContainerRemoveOptions{
				RemoveVolumes: true,
				Force:         true,
			})
			if err == nil {
				return containerInfo.ID, true
			}
		}
	}
	return "", false
}
