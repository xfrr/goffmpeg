FROM golang:1.14-alpine as mod

WORKDIR /usr/share/app

RUN apk add --update \
  --no-cache \
  build-base \
  git \
  ca-certificates \
  pkgconfig \
  openssl \
  ffmpeg \
  ffmpegthumbnailer

COPY go.mod .
COPY go.sum .

ENV CWD=/usr/share/app
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on

# Because of how the layer caching system works in Docker, the go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download

# Fetch compile daemon
RUN ["go", "get", "-u", "github.com/githubnemo/CompileDaemon"]
