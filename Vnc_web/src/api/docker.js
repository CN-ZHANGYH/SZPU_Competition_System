import request from "@/utils/request";


export function getContainerList(){
    return request({
        url: "/docker/getContainerList",
        method: "get",
    })
}


export function createDockerVm(data){
    return request({
        url: "/docker/createVm",
        method: "post",
    })
}


export function getDockerInfo(){
    return request({
        url: "/docker/info",
        method: "get",
    })
}