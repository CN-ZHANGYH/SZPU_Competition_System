import request from "@/utils/request";


export function getVmList() {
    return request({
        url: "/vm/getList",
        method: "get"
    })
}

export function getVm(query){
    return request({
        url: "/vm/get",
        method: "get",
        params: query
    })
}    


export function removeVm(query) {
    return request({
        url: "/vm/delete",
        method: "delete",
        params: query
    })
}

export function addVm(data) {
    return request({
        url: "/vm/add",
        method: "post",
        data: data
    })
}

export function getVmsLabel() {
    return request({
        url: "/vm/getHostAndUrl",
        method: "get",
    })
}