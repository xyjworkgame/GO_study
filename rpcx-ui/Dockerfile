FROM docker.io/golang:1.14 as builder
ENV GOSUMDB off
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

#2.指明该镜像的作者和其电子邮件
MAINTAINER Firewine "1451661318@qq.com"


COPY . /go/src/project/rpcx-ui
WORKDIR /go/src/project/rpcx-ui

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/project/rpcx-ui/rpcx-ui *.go

FROM alpine:latest

RUN echo http://mirrors.aliyun.com/alpine/v3.10/main/ > /etc/apk/repositories && \
    echo http://mirrors.aliyun.com/alpine/v3.10/community/ >> /etc/apk/repositories

RUN apk update && apk add --no-cache ca-certificates


COPY --from=builder /go/src/project/rpcx-ui/rpcx-ui /rpcx/
COPY conf /rpcx/conf
COPY templates /rpcx/templates
COPY web /rpcx/web

WORKDIR /rpcx

EXPOSE 8972

CMD ./rpcx-ui



