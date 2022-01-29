FROM golang:1.17.6-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build ./cmd/bot/

ENTRYPOINT [ "./bot" ]
