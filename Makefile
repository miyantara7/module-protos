clean:
	@echo "--- Cleanup all build and generated files ---"
	@rm -vf app/interface/grpc/proto/user_management/*.pb.go
	@rm -vf ./main

protoc: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p app/interface/grpc/proto/user_management

	@echo "--- Compiling all proto files ---"
	@cd ./app/proto/user_management && protoc -I. \
	--go_out=../../interface/grpc/proto/user_management \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/user_management *.proto \

gateway: 
	@cd ./app/proto/user_management && protoc -I. \
	--go_out=../../interface/grpc/proto/user_management \
    --grpc-gateway_out=output_directory ../../interface/grpc/proto/user_management *.proto \