# Sample usage of gRPC client stream 

# Setup

install Go
install ProtoBuf

install Grpc
```
go get google.golang.org/grpc
```

# Generate service 
```
protoc -I protos protos/msg.proto --go_out=plugins=grpc:streams
```


# Run server
```
cd server && go run main.go
```


# Run client(you can run more than one) 
```
go run main.go localhost:50051 test
```



