FROM golang:1.9.3-alpine as builder
MAINTAINER matsu-chara <matsuy00@gmail.com>

WORKDIR /go/src/github.com/matsu-chara/gol

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git

RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.toml .
COPY Gopkg.lock .
RUN dep ensure -vendor-only

COPY . .
RUN go build

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/matsu-chara/gol/gol /opt/bin/gol
EXPOSE 5656
ENTRYPOINT ["/opt/bin/gol"]
