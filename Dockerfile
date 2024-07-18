# 使用 Ubuntu 22.04 LTS 基础镜像
FROM ubuntu:22.04

# 设置工作目录为 /app
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates 
# 拷贝 server 可执行文件以及 config.yaml 到容器中的 /app 目录
COPY server /app/
COPY cfg_prod.yaml /app/cfg.yaml

# 容器启动执行命令
CMD ["./server", "cron"]
