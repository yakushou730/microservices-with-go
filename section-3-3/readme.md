What you need is create `*.proto` file.
Then compile it by proto command.
```
$ protoc section-3-3/proto/fibonacci.proto --go_out=plugins=grpc:.
```
Then the related go file would be generated for you.

The gRPC course as reference.
```
https://github.com/yakushou730/grpc-go-course/blob/master/generate.sh
```
