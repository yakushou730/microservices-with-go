```
// Compile App
$ env GOOS=linux go build -o main main.go

// Build Docker image
$ docker build -t example:1 .

// Starting Container in Foreground Mode
$ docker run --name-example-container-1 -p 8080:8080 example:1

// Starting Container in Background Mode
$ docker run -d --name-example-container-1 -p 8080:8080 example:1

// Entering a Container Command Line (Bash)
$ docker exec -it example-container-1 /bin/bash
```
