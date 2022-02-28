FROM golang:1.17-alpine

RUN apk add openssl build-base

WORKDIR /app

COPY src/go.mod src/go.sum src/

WORKDIR /app/src

RUN go mod download

WORKDIR /app

COPY . .

WORKDIR /app/scripts

RUN ./init.sh

WORKDIR /app/src

RUN go build -o /app/proxy main.go

WORKDIR /app

CMD ["/app/proxy"]

EXPOSE 8080
EXPOSE 80