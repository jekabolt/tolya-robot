FROM golang:1.12-alpine 

ENV GO111MODULE=on

RUN apk add --no-cache git libgit2-dev alpine-sdk

WORKDIR /go/src/github.com/jekabolt/tolya-robot

COPY go.mod .
COPY go.sum .
# install dependencies
RUN go mod download

COPY ./ ./

RUN go build -o ./bin/tolya-robot ./cmd/

FROM alpine:latest

WORKDIR /go/src/github.com/jekabolt/tolya-robot

COPY --from=0 /go/src/github.com/jekabolt/tolya-robot .

EXPOSE 80

CMD ["/go/src/github.com/jekabolt/tolya-robot/bin/tolya-robot"]
