# base image for build environment
FROM golang:alpine AS build-env

# set workdir
WORKDIR /usr/app

# handle dependencies, minimize cache busting and rebuilds
COPY go.mod .
RUN go mod download

# COPY and build executable
COPY . .
RUN go build -o rediboard

# reduce image size
FROM alpine

# set region
RUN apk update --no-cache
RUN apk add --no-cache tzdata
ENV TZ America/Los_Angeles

# set workdir
WORKDIR /usr/rediboard

# COPY templates
COPY --from=build-env /usr/app/templates templates
# COPY executable
COPY --from=build-env /usr/app/rediboard server

# default command
CMD ["./server"]
