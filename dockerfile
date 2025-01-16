FROM golang:1.23.4-alpine3.21
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN go env -w GO111MODULE=on
WORKDIR /gin
COPY . .
RUN go mod tidy
RUN go build -o /usr/local/bin/gin main.go
ENTRYPOINT ["/usr/local/bin/gin"]