# specify a base image
FROM golang:alpine AS builder

# select working dir
WORKDIR /build/simple-web

# COPY all files in context to working dir
COPY . .

# build server
RUN go mod download
RUN go build -o /app/server server.go

# compress image
FROM alpine

# configure time
RUN apk update --no-cache
RUN apk add --no-cache tzdata
ENV TZ America/Los_Angeles

# select working dir
WORKDIR /app

# COPY directory from builder
COPY --from=builder /app/server /app/server

# default command
CMD ["./server"]