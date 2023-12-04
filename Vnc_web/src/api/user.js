import request from "@/utils/request";


export function register(data){
    return request({
        url: "/user/register",
        method: "post",
        data: data
    })
}


export function login(data){
    return request({
        url: "/user/login",
        method: "post",
        data: data
    })
}

export function getUserInfo(query){
    return request({
        url: "/user/userInfo",
        method: "get",
        params: query
    })
}


export function UpdateUserInfo(data){
    return request({
        url: "/user/update",
        method: "put",
        data: data
    })
}
