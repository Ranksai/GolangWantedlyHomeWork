FROM golang:1.9

RUN go get github.com/gin-gonic/gin

EXPOSE 8080
