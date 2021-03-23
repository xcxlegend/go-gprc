
protoc -I ./api --go_out=plugins=grpc:./pkg/pb ./api/api.proto