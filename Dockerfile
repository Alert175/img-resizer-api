FROM golang:1.19.0-alpine3.16 as builder

# RUN cat /etc/os-release
RUN apk update
RUN apk add build-base
RUN apk add --update --no-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing --repository http://dl-3.alpinelinux.org/alpine/edge/main vips-dev

RUN mkdir /app
RUN mkdir /app/logs

ADD . /app

WORKDIR /app

RUN go build -o . main.go

CMD ["/app/main"]