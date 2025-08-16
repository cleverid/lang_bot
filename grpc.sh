# user
mkdir -p ./services/user/clients/user
protoc -I=./services/user/contracts/ \
   --go_out=./services/user/clients/user --go_opt=paths=source_relative \
   --go-grpc_out=./services/user/clients/user --go-grpc_opt=paths=source_relative \
   ./services/user/contracts//*.proto

# gate
mkdir -p ./services/gate/clients/user
protoc -I=./services/user/contracts/ \
   --go_out=./services/gate/clients/user --go_opt=paths=source_relative \
   --go-grpc_out=./services/gate/clients/user --go-grpc_opt=paths=source_relative \
   ./services/user/contracts//*.proto

