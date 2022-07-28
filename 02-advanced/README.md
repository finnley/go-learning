# advanced

## grpc

```shell
go get github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
```

```shell
protoc -I . helloworld.proto --go_out=plugins=grpc:.
```

## redis

If you are using Redis 6, install go-redis/v8:
```shell
go get github.com/go-redis/redis/v8
```

```shell
go get github.com/gomodule/redigo/redis
```