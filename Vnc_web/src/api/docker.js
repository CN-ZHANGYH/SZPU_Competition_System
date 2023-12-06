import request from "@/utils/request";


export function getContainerList(){
    return request({
        url: "/docker/getContainerList",
        method: "get",
    })
}


export function createDockerContainer(data){
    return request({
        url: "/docker/create",
        method: "post",
        data: data
    })
}


export function getDockerInfo(){
    return request({
        url: "/docker/info",
        method: "get",
    })
}

export function startDockerContainer(query){
    return request({
        url: "/docker/start",
        method: "post",
        params: query
    })
}

export function stopDockerContainer(query){
    return request({
        url: "/docker/stop",
        method: "post",
        params: query
    })
}

export function removeDockerContainer(query){
    return request({
        url: "/docker/delete",
        method: "delete",
        params: query
    })
}


export function searchContainer(query){
    return request({
        url: "/docker/searchContainer",
        method: "get",
        params: query
    })
}

export function selectDockerImagesLabel(){
    return request({
        url: "/docker/selectDockerImagesLabel",
        method: "get",
    })
}
