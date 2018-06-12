FROM golang:alpine as builder

COPY . /go/src/AddMyKeyPlease/

WORKDIR /go/src/AddMyKeyPlease/

RUN go get -d -v

RUN go build -o /go/bin/AddMyKeyPlease

FROM alpine

COPY --from=builder /go/src/AddMyKeyPlease/entrypoint.sh .

COPY --from=builder /go/bin/AddMyKeyPlease /go/bin/AddMyKeyPlease

ENV KEY="none"

ENV USERNAME="root"

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]

