# 使用的基础镜像
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app
# 复制应用程序代码到容器中
COPY . .

# 编译应用程序
RUN go env -w GOPROXY=https://goproxy.cn,direct && go build -o stone .


# 使用轻量级的 alpine 镜像作为最终基础镜像
FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
# 设置工作目录
WORKDIR /app

# 从构建镜像中复制编译好的应用程序
COPY --from=builder /app/stone .
# 将配置文件复制到容器中
COPY conf /app/conf
# 暴露端口
EXPOSE 8880

# 启动应用程序
CMD ["./stone"]
