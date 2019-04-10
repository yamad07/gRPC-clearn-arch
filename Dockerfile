FROM golang as builder

ADD . /go/src/github.com/yamad07/gRPC-sample
WORKDIR /go/src/github.com/yamad07/gRPC-sample

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

RUN go build -o /todo cmd/server/main.go

FROM debian:7.11-slim

COPY --from=builder /todo /todo

ENV PORT 19003
CMD ["/todo"]
