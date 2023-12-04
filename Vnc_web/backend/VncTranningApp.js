const express = require("express");
const fs = require("fs");
const {resolve} = require("path")
const app = express();

const bodyParser = require('body-parser')
const multer = require('multer')
const upload = multer()
app.use(bodyParser.json()) // for parsing application/json
app.use(bodyParser.urlencoded({ extended: true })) // for parsing application/x-www-form-urlencoded


var vmList = []
var contentList = []

const FILE_PATH = resolve('./vmList.txt')
const FILE_PATH_CONTENT = resolve("./markdown.txt")

load()

// 定义变量
const STATUS_OK = 200;
const STATUS_ERROR = 400;
const STATUS_SYSTEM_ERROR = 500;

// 定义返回结果
class Result {
    constructor(code,data,msg) {
        this.code = code;
        this.data = data;
        this.msg = msg;
    }
}

// 定义VNC的实例信息
class VncInstance {
    constructor(name,host,password,port,url) {
        this.id = null;
        this.name = name;
        this.host = host;
        this.password = password;
        this.port = port;
        this.url = url;
    }
}

class ContentInfo {
    constructor(id,name,type,vmList,content) {
        this.id = id;
        this.name = name;
        this.type = type;
        this.vmList = vmList;
        this.content = content;
    }
}

// 添加VNC虚拟机实例信息
app.post('/addVm', async (req, res) => {
    // 接收来自客户端的书籍信息，并存储到数据库中
    const {name, host, password, port, url} = req.body;
    const vm = new VncInstance(name,host,password,port,url);
    let vmUrl = "http://" + host + ":" + port + "/vnc.html"
    vm.id = vmList.length
    vm.url = vmUrl
    vmList.push(vm);
    append()
    res.json(new Result(STATUS_OK,vm,"添加成功"))
});


// 查询所有的VNC实例信息
app.get("/getVms",async (req,res) => {
    res.json(vmList)
})
app.get("/getVmsLabel",async (req,res) => {
    let data = []
    for (let i = 0; i < vmList.length; i++) {
        let info = {};
        info.value = vmList[i].host
        info.label = vmList[i].host
        data.push(info)
    }
    res.json(data)
})

// 查询VNC实例的详细信息
app.get("/getVmInfo",async (req,res) => {
    const {id} = req.query
    let instance = new VncInstance();
    for (let i = 0; i < vmList.length; i++) {
        if (vmList[i].id == id){
            instance = vmList[i]
            break
        }
    }
    res.json(new Result(STATUS_OK,instance,"查询成功"))
})

// 删除VNC的实例信息
app.delete("/removeVm",async (req,res) => {
    const {id} = req.query
    save(id)
    res.json(new Result(STATUS_OK,id,"删除成功"))
})



app.post("/addExam",async (req,res) => {
    const {name,type,vmList,content} = req.body
    var newContent = JSON.parse(content)._rawValue;
    const contentId = contentList.length;
    var contentInfo = new ContentInfo(contentId,name,type,vmList,newContent);
    contentList.push(contentInfo);
    appendContentInfo()
    res.json(new Result(STATUS_OK,null,"添加成功"));
})

app.post("/getExamInfo",async (req,res) => {
    const {id} = req.query
    let contentInfo = new ContentInfo();
    for (let i = 0; i < contentList.length; i++) {
        if (contentList[i].id == id){
            contentInfo = contentList[i]
            break
        }
    }
    console.log(contentInfo)
    res.json(new Result(STATUS_OK,contentInfo,"查询成功"))
})

app.post("/getExamList",async (req,res) => {
    res.json(contentList)
})


// 获取所有的试题的选项
app.get("/getCountExam",async (req,res) => {
    let data = []
    for (let i = 0; i < contentList.length; i++) {
        let info = {};
        info.value = contentList[i].id
        info.label = contentList[i].name
        data.push(info)
    }
    res.json(data)
})

// 初始化加载数据
function load(){
    let data = fs.readFileSync(FILE_PATH,"utf-8");
    let oldContentData = fs.readFileSync(FILE_PATH_CONTENT,"utf-8");
    let oldVmList = data.split("\n")
    let oldContentList = oldContentData.split("\n");

    oldVmList.pop()
    oldContentList.pop()

    for (let i = 0; i < oldVmList.length; i++) {
        let vmInfo = JSON.parse(oldVmList[i])
        vmList.push(vmInfo)
    }

    for (let i = 0; i < oldContentList.length; i++) {
        let contentInfo = JSON.parse(oldContentList[i])
        contentList.push(contentInfo)
    }

}


// 文件内容追加数据
function append() {
    fs.appendFileSync(FILE_PATH,JSON.stringify(vmList[vmList.length-1]) + "\n")
}

function appendContentInfo() {
    const filePath = resolve("./markdown.txt")
    fs.appendFileSync(filePath,JSON.stringify(contentList[contentList.length -1]) + "\n")
}

// 删除覆盖文件内容
function  save(id) {
    for (let i = 0; i < vmList.length; i++) {
        if (vmList[i].id == id){
            vmList[i] = vmList[vmList.length -1]
            vmList.pop()
        }
    }
    let data = ""
    for (let i = 0; i < vmList.length; i++) {
        data += JSON.stringify(vmList[i]) + "\n"
    }
    fs.writeFileSync(FILE_PATH,data);

}

// 服务端设置响应头
app.use(function(req,res,next){
    res.header("Access-Control-Allow-Origin","*")
    res.header("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
    next()
})


// 启动项目
app.listen(3000, () => console.log('Server started at http://localhost:3000'));




