import request from "@/utils/request";

export function addExam(data) {
    return request({
        url: "/content/add",
        method: "post",
        data: data
    })
}

export function getExamList() {
    return request({
        url: "/content/getList",
        method: "post"
    })
}

export function getCountExam() {
    return request({
        url: "/content/getLabel",
        method: "get"
    })
}


export function getExamInfo(query) {
    return request({
        url: "/content/get",
        method: "post",
        params: query
    })
}
export function removeContentInfo(query) {
    return request({
        url: "/content/delete",
        method: "delete",
        params: query
    })
}



