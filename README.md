# gRPC_example
This is a simple example of gRPC. Modified from https://github.com/fafalina/https_service.
1. Use golang to create gRPC service. (server)
2. The https service obtains data from the firebase real-time database. (SQL)
3. Use golang to create gRPC service. (client)

# Docker
```bash
docker compose up
```

# Generate grpc files command
## Golang
Golang client and server work normally.
```shell
protoc --go_out=. --go-grpc_out=. .\proto\timestamp.proto
```
## Angular
I haven't found a way to use gRPC normally for this part. So instead, I'll use the Golang client.
```shell
ng new client --no-standalone --routing --ssr=false
cd client
protoc --grpc-web_out=import_style=typescript,mode=grpcweb:".\client\src\app" --js_out="import_style=commonjs,binary:.\client\src\app" .\proto\timestamp.proto
```
