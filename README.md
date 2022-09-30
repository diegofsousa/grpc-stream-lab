# GRPC stream lab

> Repository for testing with GRPC and its streaming data functionality.

## Useful commands

Stubs from proto file:

`protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.`