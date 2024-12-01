# 第一阶段：使用 Golang 1.23.3 编译文件
FROM golang:1.23.3-alpine3.20 AS builder
LABEL authors="zen"
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN apk update
RUN apk add build-base
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
RUN mkdir /gin
WORKDIR /gin
COPY . .
RUN go mod tidy
RUN go build -o /usr/local/bin/gin main.go

FROM alpine:3.20.3
# 从第一阶段复制编译好的二进制文件到最终镜像中
COPY --from=builder /usr/local/bin/gin /usr/local/bin/gin
ENTRYPOINT ["/usr/local/bin/gin"]
