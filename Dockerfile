FROM golang:1.17-stretch AS builder
RUN mkdir /api
ADD . /api
WORKDIR /api

RUN go env -w GOPROXY=https://proxy.golang.org
RUN go env -w CGO_ENABLED="0"
RUN go env -w GO111MODULE='on'

RUN go mod download
RUN go build -o bin/api -ldflags="-s -w" ./cmd/filesd

FROM alpine:latest
RUN mkdir -p /go/bin
WORKDIR /go/bin
RUN mkdir /go/bin/files

COPY --from=builder /api/bin/api .
COPY --from=builder /api/resources/application.yml resources/
COPY --from=builder /api/internal internal/

EXPOSE 8090
ENTRYPOINT /go/bin/api
