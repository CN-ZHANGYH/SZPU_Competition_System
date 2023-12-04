# 基于Vue3 + Naive + Go + Gin的竞赛系统


## 项目介绍
使用Vue3 + Naive UI实现的前端、使用Go + Gin作为Vnc_Server的后端服务的前后端分离的Web竞赛系统项目，并使用Docker容器化部署提高用户的体验效率以及生产环境的完美隔离。

## 项目结构
- Vnc_web  
- Vnc_server

## 技术栈
- Vue3
- Naive UI
- Go
- Gin
- Docker

## 安装步骤
1. 安装前端依赖：
   ```bash
   cd Vnc_web/
   npm install

1. 安装后端依赖：

   ```
   cd Vnc_server/
   go mod tidy
   ```

## 使用说明

1. 启动前端服务：

   ```
   cd Vnc_web/
   npm run serve
   ```

2. 启动后端服务：

   ```
   cd Vnc_server/
   go run main.go
   ```

3. 访问Web竞赛系统： 打开浏览器并访问 http://localhost:9999

## Docker容器化部署

1. 构建前端镜像：

   ```
   cd Vnc_web/
   docker build -t frontend-image .
   ```

2. 构建后端镜像：

   ```
   cd Vnc_server/
   docker build -t backend-image .
   ```

3. 运行Docker容器：

   ```
   docker run -d -p 9999:9999 szpt-vnc-web:latest
   docker run -d -p 8888:8888 szpt-vnc-server:latest
   ```

## 贡献者

- 张宇豪（邮箱：2756864799@qq.com）