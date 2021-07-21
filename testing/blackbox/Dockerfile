FROM golang:1.16.3-alpine3.13

RUN mkdir /app

COPY . /app

WORKDIR /app

ENV CGO_ENABLED=0

RUN go test -c ./testing/blackbox -o newblackbox

CMD ["/app/newblackbox"]