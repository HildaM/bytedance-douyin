FROM golang:latest

# 为我们的镜像设置必要的环境变量
ENV	GOPROXY="https://goproxy.cn,direct"

# 工作目录
WORKDIR /home/sota/bytedance

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o server .

# Run the server executable
CMD [ "/app/server" ]


