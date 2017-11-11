FROM golang:1.9

EXPOSE 8080

RUN go get github.com/gin-gonic/gin

RUN go get github.com/go-xorm/xorm

RUN go get github.com/lib/pq

ADD ../GolangWantedlyHomeWork /go/src

WORKDIR /go/src/github.com/Ranksai/GolangWantedlyHomeWork/src

COPY ./ /go/src/github.com/Ranksai/GolangWantedlyHomeWork/src

RUN go build -o server .

CMD ["/go/src/github.com/Ranksai/GolangWantedlyHomeWork/src/server"]