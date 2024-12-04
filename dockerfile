FROM golang:1.23.3-alpine3.20
LABEL authors="zen"
#RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN apk update
RUN apk add build-base sqlite-dev sqlite-libs
RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
RUN go env -w CGO_ENABLED=1
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /usr/local/bin/gin main.go
EXPOSE 8192
ENTRYPOINT ["/usr/local/bin/gin"]
# docker run -dit --name sql -v C:\Users\zen\Github\DeepLXByLinuxdo:/data -p 8192:8192 golang:1.23.3-alpine3.20 ash
# docker exec -it sql ash