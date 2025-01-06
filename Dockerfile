# 使用官方 Go 镜像作为基础镜像
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将本地代码复制到容器中的工作目录
COPY main.go .

# 使用 Go 编译器编译 Go 程序
RUN go mod tidy && \
    go build -o shark main.go

# 使用更小的镜像来运行应用
FROM alpine:latest

# 安装依赖 (如果需要，如 ca-certificates 用于 HTTPS 请求)
#RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译后的二进制文件
COPY --from=builder /app/shark .

# 暴露容器的端口
EXPOSE 8000

# 启动应用
CMD ["./shark"]
