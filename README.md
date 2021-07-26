# grpc_app

Simple grpc app for practise

Codegen command:
```
protoc -I ./proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative ./proto/app.proto
```
