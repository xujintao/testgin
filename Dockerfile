FROM alpine:3.7
MAINTAINER xujintao <xujintao@126.com>

# 指定镜像工作目录
WORKDIR /go

# 打包应用程序
ADD ./testgin ./

# 打包静态文件
ADD ./static ./static/

ENTRYPOINT ["./testgin"]