## Dockerfile
### Content
1.  [ADD](#ADD)
2.  [ARG](#ARG)
3.  [CMD](#CMD)
4.  [COPY](#COPY)
5.  [ENTRYPOINT](#ENTRYPOINT)
6.  [ENV](#ENV)
7.  [EXPOSE](#EXPOSE)
8.  [FROM](#FROM)
9.  [HEALTHCHECK](#HEALTHCHECK)
10. [LABEL](#LABEL)
11. [ONBUILD](#ONBUILD)
12. [RUN](#RUN)
13. [SHELL](#SHELL)
14. [STOPSIGNAL](#STOPSIGNAL)
15. [USER](#USER)
16. [VOLUME](#VOLUME)
17. [WORKDIR](#WORKDIR)

### Instructions
```shell script
# Key:
#     <x> required variation
#     [x] optional variation
#     | OR
```

###### ADD
```dockerfile
# Usage:
#       ADD [--chown=<user>:<group>] <src>... <dest>
#       ADD [--chown=<user>:<group>] ["<src>", ..., "<dest>"] (required for paths with whitespace)
#
# ADD copies new files, directories, or remote file URLS from <src> and adds them to the filesystem of the image at path <dest>
#
# Rules:
#       Multiple <src> resources may be specified but if they are files or directories, their paths are interpreted as relative to the source of the context of the build
#       Each <src> may contain wildcards and matching will be done using Go’s filepath.Match rules
#       The <dest> is an absolute path, or a path relative to WORKDIR, into which the source will be copied inside the destination container
#
# Notes:
#       OPT FOR COPY INSTEAD OF ADD
#       ADD can get remote files and uncompress files so it can have unpredictable behavior

ADD test.txt /test/
```

###### ARG
```dockerfile
# Usage:
#       ARG <name>[=<default value>]
#
# ARG defines a variable that users can pass at build-time to the builder using `docker build --build-arg <varname>=<value>`
#
# Rules:
#       There can be multiple ARGs in a dockerfile
#       Scope of an ARG is dictated by dockerfile, not CLI, and is only valid in the same build stage it was defined
#       ARG can appear before the first FROM line
#       ENV always overrides ARG
#
# Note:
#       DO NOT USE BUILD-TIME VARIABLES FOR PASSING SECRETS LIKE github keys, user credentials etc.

# No default value
ARG user1

# With default value
ARG user1=someuser

# Split across build stages
FROM busybox
ARG SETTINGS
RUN ./run/setup $SETTINGS

FROM busybox
ARG SETTINGS
RUN ./run/other $SETTINGS

# ENV overriding
# docker build --build-arg CONT_IMG_VER=v2.0.1 .
FROM ubuntu
ARG CONT_IMG_VER
ENV CONT_IMG_VER=v1.0.0
RUN echo $CONT_IMG_VER
# prints v1.0.0
```

###### CMD
```dockerfile
# Usage:
#       CMD ["executable", "param1", "param2"] (exec form, this is the preferred form)
#       CMD ["param1", "param2"] (as default parameters to ENTRYPOINT)
#       CMD command param1 param2 (shell form)
#
# CMD provides defaults for an executing container
#
# Rules:
#       Only the last CMD will take effect
#       If executable is omitted, ENTRYPOINT must be specified
#       In exec form, a command shell is not invoked    
#       In exec form, CMD must be expressed as a JSON array

CMD ["./server"]
```

###### COPY
```dockerfile
# Usage:
#       COPY [--chown=<user>:<group>] <src>... <dest>
#       COPY [--chown=<user>:<group>] ["<src>", ..., "<dest>"]
#       COPY --from=<name> <src>... <dest> (used to set the src location to a previous build stage that will be used instead of context sent by user)
#
# COPY copies new files or directories from <src> and adds them to the filesystem of the container at the path <dest>
#
# Rules:
#       Multiple allowed in a Dockerfile
#       Relative paths interpreted relative to the source of the context of the the build
#
# Notes:
#       Matching done by Go's filepath.Match rules

COPY test.txt relativeDir/
COPY test.txt /absoluteDir/
```

###### ENTRYPOINT
```dockerfile
# Usage:
#       ENTRYPOINT ["executable", "param1", "param2"]
#       ENTRYPOINT command param1 param2
#
# ENTRYPOINT allows you to configure a container that will run as an executable. Command line arguments to docker run <image> will be appended after all elements in an exec form ENTRYPOINT, and will override all elements specified using CMD. This allows arguments to be passed to the entry point, i.e., docker run <image> -d will pass the -d argument to the entry point. You can override the ENTRYPOINT instruction using the docker run --entrypoint flag.
#   
# Notes:
#       SEE CMD

# Exec form
FROM ubuntu
ENTRYPOINT ["top", "-b"]
CMD ["-c"]

# Shell form
FROM ubuntu
ENTRYPOINT exec top -b
```

###### ENV
```dockerfile
# Usage:
#       ENV <key>=<value> ...
#
# ENV sets the environment variable <key> to the value <value>
#
# Rules:
#       The value will be in the environment of all subsequent instructions in the build stage
#       The environment variables set using ENV will persist when a container is run from the resulting image
#
# Notes:
#       View the values using `docker inspect`
#       Change the values using `docker run --env <key>=<value>

# Multiple
ENV MY_NAME="John Doe"
ENV MY_DOG=Rex\ The\ Dog
ENV MY_CAT=fluffy

# Split
ENV MY_NAME="John Doe" MY_DOG=Rex\ The\ Dog \
    MY_CAT=fluffy
```

###### EXPOSE
```dockerfile
# Usage:
#       EXPOSE <port>[/<protocol>] ...
#
# EXPOSE informs Docker that the container listens on the specified network ports at runtime
#
# Rules:
#       
# Notes:
#       By default, EXPOSE assumes TCP but can also be specified for UDP
#       EXPOSE does not publish the port, it's as documentation between builder and runner.
#       To publish one or more ports, use `docker run -p ...`
#       To publish all ports and map them to high-order ports, use `docker run -P ...`

# docker run -p 80:80/tcp -p 80:80/udp ...
EXPOSE 80/udp
EXPOSE 80/tcp
```

###### FROM
```dockerfile
# Usage:
#       FROM [--platform=<platform>] <image>[:tag | @<digest>] [AS <name>]
#
# FROM initializes a new build stage and sets the base image for subsequent instructions
#
# Rules:
#       FROM must be first non-comment, non-ARG instruction in the dockerfile
#       FROM can appear multiple times within a single dockerfile
#       <tag> & <digest> values are optional. If ommitted, the builder assumes 'latest'

FROM alpine
```

###### HEALTHCHECK
```dockerfile
# Usage:
#       HEALTHCHECK [OPTIONS] CMD <command>
#       HEALTHCHECK NONE
#
# HEALTHCHECK tells Docker how to test a container to check that it is still working
#
# OPTIONS:   
#       --interval=DURATION (default: 30s) (The health check will first run interval seconds after the container is started, and then again interval seconds after each previous check completes)
#       --timeout=DURATION (default: 30s) (If a single run of the check takes longer than timeout seconds then the check is considered to have failed)
#       --start-period=DURATION (default: 0s) (provides initialization time for containers that need time to bootstrap)
#       --retries=N (default: 3) (It takes retries consecutive failures of the health check for the container to be considered `unhealthy`)
#
# Rules:
#       Only one HEALTHCHECK instruction in a Dockerfile, the last one will take effect
#       
# Notes:
#       HEALTHCHECK can detect cases such as a web server that is stuck in an infinite loop and unable to handle new connections, even though the server process is still running
#       Statuses -- 
#                0: success - the container is healthy and ready for use
#                1: unhealthy - the container is not working correctly
#                2: reserved - do not use this exit code
#       Use `docker inspect` to review health status

# check every five minutes or so that a web-server is able to serve the site’s main page within three seconds
HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost/ || exit 1
```

###### LABEL
```dockerfile
# Usage:
#       LABEL [<key>=<value>...]
#
# LABEL add metadata to an image
#
# Rules:
#       Labels included in the base or parent images (images in the FROM line) are inherited by your image
#       Use quotes and backslashes to include spaces within a value
#       Images can have multiple labels and they can be specified on the same line
#       If a label already exists, it can be overridden
#
# Note:
#       To view an image's labels use `docker image inspect`
#       Use `--format` option to show just the labels

# Quoted spaces
LABEL "com.example.vendor"="ACME Incorporated"

# New line
LABEL description="This text illustrates \
that label-values can span multiple lines."

# Multiple on single line
LABEL multi.label1="value1" multi.label2="value2" other="value3"

# Multiple on a split line
LABEL multi.label1="value1" \
      multi.label2="value2" \
      other="value3"
```

###### ONBUILD
```dockerfile
# Usage:
#       ONBUILD <INSTRUCTION>
#
# ONBUILD adds to the image a trigger instruction to be executed at a later time, when the image is used as the base for another build 

ONBUILD ADD . /app/src
ONBUILD RUN /usr/local/bin/python-build --dir /app/src
```

###### RUN
```dockerfile
# Usage:
#       RUN <command> (shell form, the command is run in a shell)
#       RUN ["<executable>", "<param1>", "<param2>"] (exec form)
#
# RUN executes any commands in a new layer on top of the current image and commits the results
#
# Rules:
#       In shell form, use `\` to continue a single RUN instruction onto the next line
#       In exec form, a command shell is not invoked

RUN go mod download
```

###### SHELL
```dockerfile
# Usage:
#       SHELL ["executable", "parameters"]
#
# SHELL allows the default shell used for the shell form of commands to be overridden
#
# Rules:
#       SHELL can appear multiple times, each overriding the previous and affects all subsequent instructions
#
# Notes:
#       SHELL must be written in JSON form in a Dockerfile
#       RUN, CMD, ENTRYPOINT are all affected by SHELL

# example
FROM microsoft/windowsservercore

# Executed as cmd /S /C echo default
RUN echo default

# Executed as cmd /S /C powershell -command Write-Host default
RUN powershell -command Write-Host default

# Executed as powershell -command Write-Host hello
SHELL ["powershell", "-command"]
RUN Write-Host hello

# Executed as cmd /S /C echo hello
SHELL ["cmd", "/S", "/C"]
RUN echo hello
```

###### STOPSIGNAL
```dockerfile
# Usage:
#       STOPSIGNAL <signal>
#
# STOPSIGNAL sets the system call signal that will be sent to the container to exit
    
STOPSIGNAL SIGKILL
```

###### USER
```dockerfile
# Usage:
#       USER <user>[:<group>]
#       USER <UID>[:<GID>]
# 
# USER sets the user name or UID and optionally the user group or GID to use when running the image and for any RUN, CMD, and ENTRYPOINT instructions that follow in the Dockerfile
    

FROM microsoft/windowsservercore
# Create Windows user in the container
RUN net user /add patrick
# Set it for subsequent commands
USER patrick
```

###### VOLUME
```dockerfile
# Usage:
#       VOLUME ["<mountable>"]
#
# VOLUME creates a mount point with the specified name and marks it as holding externally mounted volumes from native host or other containers
#
# Rules:
#       
# Notes:
#       The value can be a JSON array, VOLUME ["/var/log/"], or a plain string with multiple arguments, such as VOLUME /var/log or VOLUME /var/log /var/db
# The docker run command initializes the newly created volume with any data that exists at the specified location within the base image

FROM ubuntu
RUN mkdir /myvol
RUN echo "hello world" > /myvol/greeting
VOLUME /myvol
```

###### WORKDIR
```dockerfile
# Usage:
#       WORKDIR /path/to/workdir
#
# WORKDIR sets the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions that follow it.
#
# Rules:
#       Can be used multiple times in a Dockerfile, relative paths are relative to previous working directory
#               WORKDIR /a
#               WORKDIR b
#               WORKDIR c
#               RUN pwd
#               result: /a/b/c
#       You can only use environmental variables explicitly set in Dockerfile      

ENV DIRPATH=/path
WORKDIR $DIRPATH/$DIRNAME
RUN pwd
```
