FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/land ./api

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl bash

WORKDIR /app

COPY --from=0 /app/land /usr/local/bin/land

COPY ./api/etc/land.yaml ./api/etc/land.yaml

COPY .aws /root/.aws

EXPOSE 8888

CMD ["/usr/local/bin/land", "-f", "./api/etc/land.yaml"]
