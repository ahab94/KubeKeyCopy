FROM golang:alpine

ADD . /go/src/github.com/golang/ahab94/AddMyKeyPlease

WORKDIR /go/src/github.com/golang/ahab94/AddMyKeyPlease

Run go install

ENV KEY="none"

ENV USERNAME="root"

ENTRYPOINT ["./entrypoint.sh"]

