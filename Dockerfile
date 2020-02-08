FROM golang:1.12-alpine 

ENV GO111MODULE=on

RUN apk add --no-cache git libgit2-dev alpine-sdk

WORKDIR /go/src/gitlab.com/miapago/open-banking/obd-server

# https://divan.github.io/posts/go_get_private/
COPY .gitconfig /root/.gitconfig
COPY go.mod .
COPY go.sum .
# install dependencies
RUN go mod download

COPY ./cmd/ ./cmd/
COPY ./obd/ ./obd/
COPY ./routers/ ./routers/
COPY ./certs/ ./certs/

RUN go build -o ./bin/obd-server ./cmd/

FROM alpine:latest

WORKDIR /go/src/gitlab.com/miapago/open-banking/obd-server
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# RUN apk add --no-cache git libgit2-dev alpine-sdk
RUN apk --no-cache add curl

COPY --from=0 /go/src/gitlab.com/miapago/open-banking/obd-server .

EXPOSE 8080

CMD ["/go/src/gitlab.com/miapago/open-banking/obd-server/bin/obd-server"]
