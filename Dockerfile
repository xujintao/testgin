FROM alpine:3.7
MAINTAINER xujintao <xujintao@126.com>

# 打包应用程序
ADD ./testgin /go/

ENTRYPOINT ["/go/testgin"]