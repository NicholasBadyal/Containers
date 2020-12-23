# specify a base image
FROM golang:alpine AS builder

# select working dir, set builder context
WORKDIR /app

# COPY dependency files, minimizes cache busting and rebuilds
COPY go.mod .
# install dependencies
RUN go mod download

# COPY all files
COPY . .
# build server
RUN go build -o /app/server server.go

# create new image
FROM alpine

# configure time
RUN apk update --no-cache
RUN apk add --no-cache tzdata
ENV TZ America/Los_Angeles

# COPY directory from builder context
COPY --from=builder app/server app/simple-server

# select working dir
WORKDIR /app

# default command
CMD ["./simple-server"]