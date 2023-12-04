FROM nginx:latest

# 复制配置文件
COPY nginx.conf /etc/nginx/nginx.conf
COPY vnc-vue.conf /etc/nginx/conf.d/vnc-vue.conf
ADD dist/ /usr/share/nginx/dist/
# 暴露当前的端口
EXPOSE 9999

# 运行Nginx服务
CMD ["nginx", "-g", "daemon off;"]
