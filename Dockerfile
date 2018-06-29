FROM alpine:3.7
MAINTAINER xujintao <xujintao@126.com>

# 打包应用程序
ADD ./testgin /go/

# 打包静态文件
ADD ./static /go/

ENTRYPOINT ["/go/testgin"]