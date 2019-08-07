# 1.使用docker-compose进行容器编排
## 编写docker-compose.yml文件
```
指定服务名、及镜像的来源、暴露的端口号映射关系、启动container的先后关系
项目代码中多使用服务名进行访问
```

## 启动一个docker-compose
```
#进入到docker-compose文件的层级
docker-compose up --build
```

## 关闭服务并删除container
```
##进入到docker-compose文件的层级
docker-compose down
```

## 各服务介绍
```
consul-manager：配置的推送与管理
userwebPRJ：websocket服务实例
```