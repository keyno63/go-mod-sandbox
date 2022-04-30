FROM golang:1.18-stretch

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080 2345
