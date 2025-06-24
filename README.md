# BeTrueLangBot

# TODO

- Copy logger from nm
- Run GRPC server
- Call request from grpc client
- Work with GRPC with Postman

### GRPC

```bash
protoc --proto_path=./contracts --go_out=./contracts --go_opt=paths=source_relative --go-grpc_out=./contracts --go-grpc_opt=paths=source_relative ./contracts/user.proto
```
