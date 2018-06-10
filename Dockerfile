FROM golang:alpine

ADD . /go/src/github.com/golang/ahab94/AddMyKeyPlease

WORKDIR /go/src/github.com/golang/ahab94/AddMyKeyPlease

RUN go install

ENV KEY="none"

ENV USERNAME="root"

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]

