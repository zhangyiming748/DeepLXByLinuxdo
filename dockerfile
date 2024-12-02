FROM golang:1.23.3-alpine3.20
LABEL authors="zen"
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN apk update
RUN apk add build-base
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /usr/local/bin/gin main.go
ENTRYPOINT ["/usr/local/bin/gin"]