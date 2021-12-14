FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app/
COPY . .
EXPOSE 8080

RUN apk add --no-cache gcc g++

RUN go get -d -v

RUN go build -o /app/bin

EXPOSE "8080"

ENTRYPOINT ["/app/bin"]
