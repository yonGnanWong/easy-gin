FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /gin

WORKDIR /gin

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application

FROM alpine:3.7

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /gin/application /www/gin/application

ADD ./Config /www/gin/Config

RUN chmod +x /www/gin/application

CMD ["/www/gin/application"]
