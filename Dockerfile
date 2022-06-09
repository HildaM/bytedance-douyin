FROM jrottenberg/ffmpeg:4-alpine

# 安装Golang
RUN mkdir /go

RUN cd /go \
  && wget https://golang.google.cn/dl/go1.17.11.linux-amd64.tar.gz \
  && tar -C /usr/local -zxf go1.17.11.linux-amd64.tar.gz \
  && rm -rf /go/go1.17.11.linux-amd64.tar.gz \
  && mkdir /lib64 \
  && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH

# docker中的工作目录
WORKDIR $GOPATH/src/douyin
# 将当前目录同步到docker工作目录下，也可以只配置需要的目录和文件（配置目录、编译后的程序等）
ADD . ./

RUN mkdir -p /home/files/videos
RUN mkdir -p /home/files/images

# 由于所周知的原因，某些包会出现下载超时。这里在docker里也使用go module的代理服务
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"
# 指定编译完成后的文件名，可以不设置使用默认的，最后一步要执行该文件名
RUN go build -o douyin .
EXPOSE 8080
# 这里跟编译完的文件名一致
ENTRYPOINT  ["./douyin"]
