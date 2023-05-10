FROM golang:1.16-alpine

WORKDIR /app
# 设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
# 此处的 . 即为 /app 目录
COPY main/lesson2.go ./

RUN go build -o /go-http lesson2.go

EXPOSE 80

CMD [ "/go-http" ]