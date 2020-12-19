# Dockerfile

### Instructions
```dockerfile
# FROM
# Usage:
#       FROM <image>
#       FROM <image>:<tag>
#       FROM <image>:<digest>
#
# FROM initializes a new build stage and sets the base image for subsequent instructions
#
# Rules:
#       FROM must be first non-comment instruction in the Dockerfile
#       FROM can appear multiple times within a single Dockerfile
#       <tag> & <digest> values are optional. If ommitted, the builder assumes 'latest'

FROM alpine AS builder
```

```dockerfile
# RUN
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

```dockerfile
# CMD
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