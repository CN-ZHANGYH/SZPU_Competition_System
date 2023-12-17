package docker

import (
	"Social_Gin/model/response"
	model "Social_Gin/model/vnc"
	"Social_Gin/model/vo"
	"encoding/json"
	"fmt"
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
	"strings"
)

var DOCKER_CLI *client.Client

func InitDockerClient() {

	remoteDockerURL := "tcp://106.52.203.212:2376"

	cli, err := client.NewClientWithOpts(
		client.WithHost(remoteDockerURL),
		// client.WithVersion("1.41"),
		client.WithAPIVersionNegotiation(),
		client.WithTLSClientConfig("cert/ca.pem", "cert/cert.pem", "cert/key.pem"))

	DOCKER_CLI = cli
	if err != nil {
		log.Fatal("连接Docker失败")
	}

	ctx := context.Background()
	dockerInfo, err := cli.Info(ctx)
	if err != nil {
		fmt.Println("Failed to get Docker info:", err)
		os.Exit(1)
	}

	// 将Info信息转换为JSON格式并打印到控制台
	infoJSON, err := json.MarshalIndent(dockerInfo, "", "    ")
	if err != nil {
		fmt.Println("Failed to convert to JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Docker Info:")
	fmt.Println(string(infoJSON))
}

// CreateVmContainer 创建Docker的虚拟机
func CreateVmContainer(dockerVmInfo *model.DockerVmInfo, ctx *gin.Context) error {
	fmt.Println(dockerVmInfo)
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
	mounts := []mount.Mount{}
	if !strings.EqualFold(dockerVmInfo.Host_Volume_List[0], "") {
		for index, source_path := range dockerVmInfo.Host_Volume_List {
			mounts = append(mounts, mount.Mount{
				Type:     mount.TypeBind,
				Source:   source_path,                               // 宿主机上要挂载的文件夹路径
				Target:   dockerVmInfo.Container_Volume_List[index], // 容器内挂载的路径
				ReadOnly: false,                                     // 是否只读
			})
		}
	}

	portBindings := nat.PortMap{}

	for index, container_port := range dockerVmInfo.Container_Port_List {
		new_container_port := nat.Port(fmt.Sprintf("%s/tcp", container_port))
		portBindings[new_container_port] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: dockerVmInfo.Host_Port_List[index],
			},
		}
	}

	// 配置容器自动重启
	hostConfig := &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: dockerVmInfo.RestartPolicy, // 设置重启策略为"always"，容器将总是自动重启
			// 可选的重启策略：
			// - "no"：无重启策略
			// - "always"：容器总是自动重启
			// - "on-failure"：容器在非零退出状态时重启（默认最多重启3次）
			// - "unless-stopped"：除非手动停止，否则容器总是自动重启
		},

		PortBindings: portBindings,

		Mounts: mounts,
	}

	// 设置容器名称
	name := "SZPT_VM_" + dockerVmInfo.ContainerName
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
	containerList, err := DOCKER_CLI.ContainerList(context.Background(), types.ContainerListOptions{
		All:    true,
		Size:   true,
		Latest: true,
	})
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
	err := DOCKER_CLI.ContainerRemove(context.Background(), containerId, types.ContainerRemoveOptions{
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

func StartDockerContainer(containerId string, ctx *gin.Context) {
	err := DOCKER_CLI.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err == nil {
		response.OkWithMessage("开启容器成功", ctx)
		return
	}
	response.FailWithMessage("当前容器不存在", ctx)
	return
}

func StopDockerContainer(containerId string, ctx *gin.Context) {
	timeout := -1
	err := DOCKER_CLI.ContainerStop(context.Background(), containerId, container.StopOptions{
		Timeout: &timeout,
	})
	if err == nil {
		response.OkWithMessage("停止容器成功", ctx)
		return
	}
	response.FailWithMessage("当前容器不存在", ctx)
	return
}

func SelectDockerContainer(containerId string, ctx *gin.Context) {
	containerList, _ := DOCKER_CLI.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	for _, containerInfo := range containerList {
		if strings.EqualFold(containerInfo.ID[:12], containerId) {
			var containerInfoList []vo.ContainerInfo
			var container_info vo.ContainerInfo
			container_info.Id = containerInfo.ID[:12]
			container_info.Names = containerInfo.Names[0][1:]
			container_info.Image = containerInfo.Image
			container_info.Status = containerInfo.State
			container_info.Ports = containerInfo.Ports

			containerInfoList = append(containerInfoList, container_info)
			response.OkWithDetailed(containerInfoList, "查询成功", ctx)
			return
		}
	}
	response.FailWithMessage("查询失败", ctx)
	return
}

func SelectDockerImagesLabel(ctx *gin.Context) {
	imagesList, _ := DOCKER_CLI.ImageList(context.Background(), types.ImageListOptions{
		All: true,
	})
	if len(imagesList) < 0 {
		response.FailWithMessage("当前没有镜像", ctx)
		return
	}
	var resultMapList []map[string]interface{}
	for _, image := range imagesList {

		var resultMap = make(map[string]interface{})
		resultMap["label"] = image.RepoTags[0]
		resultMap["value"] = image.RepoTags[0]
		resultMapList = append(resultMapList, resultMap)
	}
	response.OkWithDetailed(resultMapList, "查询成功", ctx)
	return
}

func SelectDockerNetworkLabel(ctx *gin.Context) {
	networkList, err := DOCKER_CLI.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		response.FailWithMessage("当前还未创建网络", ctx)
		return
	}
	var resultMapList []map[string]interface{}
	for _, network := range networkList {

		var resultMap = make(map[string]interface{})
		resultMap["label"] = network.Name
		resultMap["value"] = network.Name
		resultMapList = append(resultMapList, resultMap)
	}
	response.OkWithDetailed(resultMapList, "查询成功", ctx)
	return
}

func SearchDockerImage(imageName string, ctx *gin.Context) {
	searchResults, err := DOCKER_CLI.ImageSearch(context.Background(), imageName, types.ImageSearchOptions{})
	if err != nil {
		response.FailWithMessage("没有该镜像", ctx)
		return
	}
	fmt.Println(searchResults)
	var resultMapList []map[string]interface{}
	for _, image := range searchResults {
		var resultMap = make(map[string]interface{})
		resultMap["Name"] = image.Name
		resultMap["Desc"] = image.Description
		resultMapList = append(resultMapList, resultMap)
	}
	response.OkWithDetailed(resultMapList, "查询成功", ctx)
	return
}

func PullDockerImage(imageName string, ctx *gin.Context) {
	imagePull, err := DOCKER_CLI.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		response.FailWithMessage("镜像拉取失败", ctx)
		return
	}
	defer imagePull.Close()
	response.OkWithMessage("镜像拉取成功", ctx)
	return
}
