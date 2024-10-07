FROM golang:1.23.2-alpine

RUN apk update && apk upgrade && apk add procps

WORKDIR /app

COPY . /app
RUN go build main.go get_ip.go perfmon.go

CMD /app/main
