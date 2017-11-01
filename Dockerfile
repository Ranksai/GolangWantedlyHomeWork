FROM golang:1.9

EXPOSE 8080

RUN go get github.com/gin-gonic/gin

WORKDIR /go/src/GolangWantedlyHomeWork/src

COPY ./ /go/src/GolangWantedlyHomeWork/

RUN go build -o server .

CMD ["/go/src/github.com/Ranksai/GolangWantedlyHomeWork/src/server"]
