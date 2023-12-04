export function GetUser(){
    return JSON.parse(localStorage.getItem("user"))
}