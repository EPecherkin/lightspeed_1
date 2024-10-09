FROM golang:1.23.2-alpine

WORKDIR /app

COPY ./perfmon.go ./go.mod ./7_table_v2/. /app
RUN go build -o app .

CMD /app/app
