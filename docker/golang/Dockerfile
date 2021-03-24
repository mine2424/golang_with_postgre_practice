FROM golang:1.16.2

RUN go get github.com/cespare/reflex
ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY go.mod go.sum main.go ./

RUN go mod download
