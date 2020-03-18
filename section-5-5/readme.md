Run a nginx container
```
$ docker run -v $HOME/Documents/hands-on-courses/microservices-with-go/section-5-5/conf:/etc/nginx:ro --net host -d nginx
```
```
$ go run main.go -port=8000
```
