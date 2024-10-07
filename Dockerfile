FROM golang:1.23.2-alpine

WORKDIR /app

COPY . /app
RUN go build .
