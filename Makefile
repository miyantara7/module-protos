clean:
	@echo "--- Cleanup all build and generated files ---"
	@rm -vf app/interface/grpc/proto/user_management/*.pb.go
	@rm -vf app/interface/grpc/bussiness/user_management/*.pb.go
	@rm -vf ./main

protoc: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p app/interface/grpc/proto/user_management
	@mkdir -p app/interface/grpc/proto/bussiness


	@echo "--- Compiling all proto files ---"
	@cd ./app/proto/user_management && protoc -I. \
	--go_out=../../interface/grpc/proto/user_management \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/user_management *.proto \

	@cd ./app/proto/bussiness && protoc -I. \
	--go_out=../../interface/grpc/proto/bussiness \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/bussiness *.proto \

