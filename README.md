# strike_layout_template
strike project template

# TODO
- 写一个最小化部署的模板
- 加入日志配置
- 加入grpc服务
- 优雅监听服务退出
- 写一个快速部署的dockerfile
- 连接数据库
- dockerfile快速部署数据库

## Docker
```bash
# build
docker build -t <your-docker-image-name> ./docker

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```