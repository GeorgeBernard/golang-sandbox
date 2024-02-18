FROM golang:bookworm

RUN apt update && apt upgrade -y

RUN apt install -y vim

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest