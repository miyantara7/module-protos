clean:
	@echo "--- Cleanup all build and generated files ---"
	@rm -vf app/interface/grpc/proto/user_management/*.pb.go
	@rm -vf app/interface/grpc/proto/bussiness/*.pb.go
	@rm -vf app/interface/grpc/proto/e_money_service/*.pb.go
	@rm -vf app/interface/grpc/proto/top_up_service/*.pb.go
	@rm -vf ./main

protoc: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p app/interface/grpc/proto/user_management
	@mkdir -p app/interface/grpc/proto/bussiness
	@mkdir -p app/interface/grpc/proto/e_money_service
	@mkdir -p app/interface/grpc/proto/top_up_service


	@echo "--- Compiling all proto files ---"
	@cd ./app/proto/user_management && protoc -I. \
	--go_out=../../interface/grpc/proto/user_management \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/user_management *.proto \

	@cd ./app/proto/bussiness && protoc -I. \
	--go_out=../../interface/grpc/proto/bussiness \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/bussiness *.proto \

	@cd ./app/proto/e_money_service && protoc -I. \
	--go_out=../../interface/grpc/proto/e_money_service \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/e_money_service *.proto \

	@cd ./app/proto/top_up_service && protoc \
	--go_opt=paths=import \
	--go_out=../../interface/grpc/proto/top_up_service \
	--go-grpc_opt=paths=import \
	--go-grpc_out=require_unimplemented_servers=false:../../interface/grpc/proto/top_up_service *.proto \


