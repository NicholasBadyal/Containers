# Custom Images

### Build from Dockerfile via Docker Server

```shell script
# 
# Build from Dockerfile
#  

docker build .
```

```shell script
# 
# Tag image
# docker build -t <docker id>/<project name>:<version> .
#  

docker build -t captkirk/redis-server:latest .
```

### Build image manually with Docker Commit
** NOT PREFERRED
```shell script
# from initial command-line

# start base container
docker run -it alpine sh

#  manually install redis
apk add --update redis
```

```shell script
# from second command-line

# get container id
docker ps

# commit changes to image
docker commit -c 'CMD ["redis-server"]' <container id>
```