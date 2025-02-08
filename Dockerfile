# 使用官方的 Go 镜像作为构建阶段的基础镜像
FROM golang:1.23.3-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到容器的 /app 目录
COPY . .

# 下载依赖包
RUN go mod tidy

# 构建可执行文件
RUN go build -o main .

# 使用更小的基础镜像作为最终镜像
FROM alpine:latest

RUN apk add --no-cache tzdata

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件到最终镜像
COPY --from=builder /app/main .

# 定义容器启动时运行的命令
CMD ["./main"]