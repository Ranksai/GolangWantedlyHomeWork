FROM golang:1.9

EXPOSE 8080

RUN go get github.com/gin-gonic/gin

RUN go get github.com/go-xorm/xorm

RUN go get github.com/lib/pq

WORKDIR /go/src/GolangWantedlyHomeWork/src

COPY ./ /go/src/GolangWantedlyHomeWork/

RUN go build -o server .

CMD ["/go/src/GolangWantedlyHomeWork/src/server"]
