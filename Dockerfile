# 使用官方 Golang 镜像作为构建环境
FROM golang:1.21.3 as builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# 使用 scratch 作为最终镜像
FROM scratch

# 将工作目录设置为 /
WORKDIR /

# 从构建器镜像中复制编译好的应用程序
COPY --from=builder /app/myapp .

# 暴露端口 8080
EXPOSE 8080

# 运行应用程序
CMD ["./myapp"]
