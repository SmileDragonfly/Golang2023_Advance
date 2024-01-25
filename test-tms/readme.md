# Generate protobuf files
``` bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative license_service.proto
``` 

