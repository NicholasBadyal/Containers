# Docker

### Commands
```shell script
/* 
 * Creates a new container
 * docker create <flags> <image> <alt command> <command args...>
 */ 

docker create busybox ping google.com
```

```shell script
/*
 * Start one or more stopped containers
 * docker start <flags> <containers...>
 */

docker start -a 9dq9qdc98cd8gb9wvg98hbe98egb98egb

// NOTE: 
// -a, --attach          Attach STDOUT/STDERR and forward signals
// -i, --interactive     Attach container's STDIN
```

```shell script
/*
 * Run a command in a new container
 * docker run <flads> <image> <alt command> <command args...>
 */

docker run busybox
```

```shell script
/*
 * List containers
 * docker ps <flags>
 */

docker ps

// NOTE:
// -a, --all             Show all containers (default shows just running)
```

```shell script
/*
 * Stop one or more running containers
 * docker stop <flags> <containers...>
 */

docker stop 3jdir345vn

// NOTE:
// -t, --time int   Seconds to wait for stop before killing it (default 10)
```

```shell script
/*
 * Kill one or more running containers
 * docker kill <flags> <containers...>
 */

docker kill 654jf4983n

// NOTE:
// -s, --signal string   Signal to send to the container (default "KILL")
```

```shell script
/*
 * Run a command in a running container
 * docker exec <flags> <container> <alt command> <command args>
 */

docker exec -it 3hndfu3ehn redis-cli

// NOTE:
// -i, --interactive          Keep STDIN open even if not attached
// -t, --tty                  Allocate a pseudo-TTY
```

```shell script
/*
 * Open a shell in a running container
 * docker exec -it <container> <shell>
 */

docker exec -it 34jn35jn54 sh

// NOTE:
// -i, --interactive          Keep STDIN open even if not attached
// -t, --tty                  Allocate a pseudo-TTY
```

```shell script
/*
 * Run a new container with an open shell
 * docker run -it <image> <shell>
 */

docker run -it busybox sh

// NOTE:
// -i, --interactive          Keep STDIN open even if not attached
// -t, --tty                  Allocate a pseudo-TTY
```

