FROM golang:1.8.3-alpine
MAINTAINER matsu-chara <matsuy00@gmail.com>

WORKDIR /go/src/github.com/matsu-chara/gol

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git openssh

RUN go get github.com/golang/dep/cmd/dep

COPY . .
RUN dep ensure && go install

EXPOSE 5656
ENTRYPOINT ["gol"]
