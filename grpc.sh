# user
protoc --proto_path=./services/user/contracts/ \
   --go_out=./services/user/contracts/ --go_opt=paths=source_relative \
   --go-grpc_out=./services/user/contracts --go-grpc_opt=paths=source_relative \
   ./services/user/contracts/user.proto

# gate clients
mkdir -p ./services/gate/clients/user
protoc --proto_path=./services/user/contracts/ \
    --go_out=./services/gate/clients/user --go_opt=paths=source_relative \
    --go-grpc_out=./services/gate/clients/user --go-grpc_opt=paths=source_relative \
    ./services/user/contracts/user.proto

