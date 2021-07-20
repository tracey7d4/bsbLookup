FROM golang:1.16.3-alpine3.13

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build cmd/main.go

EXPOSE 8080

CMD ["/app/main"]