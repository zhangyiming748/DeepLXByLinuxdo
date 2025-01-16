# DeepLXByLinuxdo
使用linuxdo官网提供的deeplx翻译服务
# token
https://connect.linux.do 
Ll3eVegZvHU4dB0JrEB8Ee6hGHgE5ZwUDdFeuzXrCM0
在[deeplx](https://github.com/OwO-Network/DeepLX/issues/149#issuecomment-2508875209)项目失效后，我最终选择了自己搭建依赖始皇的本地翻译服务
----
我自己的局域网网络环境如下
一台树莓派5作为长期运行的服务器，容器化运行`alist`、`qbittorrent`、`navidrome`、`mysql9`、`cpolar`、`gitea`
一台nVidia 4060的Windows系统笔记本
一台macbookpro(intel)
一台macbookair(Apple Silicon)
一台Intel nuc
主要运行的程序
https://github.com/zhangyiming748/MultimediaProcessingPipeline
其中不同的单元使用不同的硬件优势
比如ffmpeg转码libh265嵌入硬字幕使用Intel nuc(cpu硬刚 坏了不心疼)
----
由于最稳定的deeplx服务还得是这里
所以选择在长期运行的设备上搭建查询服务
```yml
name: deeplx
services:
  mysql:
    container_name: mysql9
    volumes:
      - ./mysql:/var/lib/mysql
    environment:
      - MYSQL_ROOT_PASSWORD=123456
    ports:
      - 3306:3306
    image: mysql:9.1.0
    restart: always
  golang:
    depends_on:
      - mysql
    container_name: program
    volumes:
      - ./:/data
    ports:
      - 3389:8192
    environment:
      - TOKEN=${你自己的token}
    #image: golang:alpine3.20
    build:
      context: .
      tags:
        - gin:latest
      dockerfile: dockerfile
    restart: always
```
其中dockerfile如下
```
FROM golang:alpine3.20
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
RUN mkdir /gin
WORKDIR /gin
COPY . .
RUN go mod tidy
RUN go build -o /usr/local/bin/gin main.go
ENTRYPOINT ["/usr/local/bin/gin"]
```
核心调用方法
```golang
func Deeplx(src string) (dst string, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	data := map[string]string{
		"src":         src,
		"source_lang": "auto",
		"target_lang": "zh",
	}

	uri := "http://192.168.1.9:3389/api/v1/translate"
	j, err := util.HttpPostJson(headers, data, uri)
	if err != nil {
		return "deeplx 请求发生错误", err
	}
	fmt.Println(string(j))
	var result RaspiRep
	json.Unmarshal(j, &result)
	return result.Dst, nil
}
```