# gRPC_example
This is a simple example of gRPC. Modified from https://github.com/fafalina/https_service.

# Generate grpc files command
## Golang
Golang client and server work normally.
```shell
protoc --go_out=. --go-grpc_out=. .\proto\timestamp.proto
```
## Angular
I haven't found a way to use gRPC normally for this part. So TBD.
```shell
ng new client --no-standalone --routing --ssr=false
cd client
protoc --grpc-web_out=import_style=typescript,mode=grpcweb:".\client\src\app" --js_out="import_style=commonjs,binary:.\client\src\app" .\proto\timestamp.proto
```