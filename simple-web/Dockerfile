# specify a base image
FROM golang:alpine AS builder

WORKDIR /build/simple-web

COPY . .
RUN go mod download
RUN go build -o /app/server server.go

FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache tzdata
ENV TZ America/Los_Angeles

WORKDIR /app
COPY --from=builder /app/server /app/server

CMD ["./server"]