clean:
	@echo "--- Cleanup all build and generated files ---"
	@rm -vf app/interface/grpc/proto/user_management/*.pb.go
	@rm -vf app/interface/grpc/gateway/user_management/*.pb.gw.go
	@rm -vf ./main

protoc: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p app/interface/grpc/proto/user_management
	@mkdir -p app/interface/grpc/gateway/user_management


	@echo "--- Compiling all proto files ---"
	@cd ./app/proto/user_management && protoc -I. \
	--go_out=../../interface/grpc/proto/user_management \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/user_management *.proto \


gateway: 
	@echo "--- Preparing proto output directories ---"
	@mkdir -p app/interface/grpc/gateway/user_management

	@cd ./app/proto/user_management && protoc -I. \
	--go_out=../../interface/grpc/gateway/user_management  \
  	--grpc-gateway_out=../../interface/grpc/gateway/user_management --grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt generate_unbound_methods=true *.proto \