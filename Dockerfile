FROM golang:1.8.3-alpine
MAINTAINER matsu-chara <matsuy00@gmail.com>

WORKDIR /go/src/app
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache git openssh

RUN go-wrapper download && go-wrapper install

EXPOSE 5656
ENTRYPOINT ["app"]
