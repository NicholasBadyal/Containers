# Use an existing docker image as a base
FROM alpine

# Download and install dependencies
# Order of operations matters in build
RUN apk add --update gcc
RUN apk add --update redis

# Tell the image what to do when it starts as a container
CMD ["redis-server"]