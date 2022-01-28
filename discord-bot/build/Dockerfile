FROM golang:1.17.6-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/bot/" -command="./bot"