FROM golang:alpine

ENV GOPATH /go
ENV GOBIN /go/bin

RUN apk add --no-cache libc-dev git zeromq-dev libzmq libsodium-dev czmq czmq-dev gcc pkgconf

COPY . /app

WORKDIR /app

RUN go get github.com/zeromq/goczmq

RUN go install

RUN go build main.go

CMD ["./main"]
