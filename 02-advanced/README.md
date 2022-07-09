# advanced

## grpc

```shell
go get github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
```

```shell
protoc -I . helloworld.proto --go_out=plugins=grpc:.
```