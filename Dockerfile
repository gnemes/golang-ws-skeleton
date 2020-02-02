FROM golang:1.13-alpine

WORKDIR /go/src/github.com/gnemes/golang-ws-skeleton/

RUN apk upgrade && \
    apk add --no-cache bash git openssh && \
    apk add --update alpine-sdk && \
    apk add protobuf

RUN go get -d -v -u github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher
