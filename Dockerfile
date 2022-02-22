FROM golang:1.17-alpine

RUN apk add openssl

WORKDIR /app

COPY . .

WORKDIR /app/scripts

RUN ./init.sh

WORKDIR /app/src

RUN go build -o /app/proxy main.go

WORKDIR /app

CMD ["/app/proxy"]

EXPOSE 8080
