FROM alpine:3.8
MAINTAINER xujintao <xujintao@126.com>

# 指定镜像工作目录
WORKDIR /go

# 打包应用程序
ADD ./testgin ./

# 打包静态文件
ADD ./static ./static/

RUN echo "====>apk add tzdata" && apk add -U tzdata && \
    echo "====>date" && date && \
    echo "====>set timezone" && cp /usr/share/zoneinfo/Asia/Shanghai \
                                /etc/localtime &&\
    echo "====>validate date" && date &&\
    echo "====>apk del tzdata" && apk del tzdata &&\
    echo "====>alpine automatic clean repository indexes on reboot"

ENTRYPOINT ["./testgin"]