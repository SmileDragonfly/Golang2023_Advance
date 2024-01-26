# Generate protobuf files
``` bash
protoc --go_out=.\tms.proto.licensing\ --go_opt=paths=source_relative --go-grpc_out=.\tms.proto.licensing\ --go-grpc_opt=paths=source_relative license_service.proto
``` 

