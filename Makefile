export PROTOGOPPATH := app/interface/grpc/proto
export PROTOPATH := app/proto

generate-proto: 
	@echo "Generating proto file for golang and javascript ...."
	@echo "Using protoc version" $$(protoc --version), "protoc-gen-go version", $$(protoc-gen-go --version)
	@echo "Generate go struct to" $$(pwd)/$$PROTOGOPPATH
	@rm -rf $$(pwd)/$$PROTOGOPPATH && mkdir -p $$(pwd)/$$PROTOGOPPATH
	@for entry in $$(find $$PROTOPATH -iname "*.proto"); do\
		protoc --go_out=$$(pwd)/$$PROTOGOPPATH \
		--go_opt=paths=import \
		--go_opt=module=github.com/vins7/module-protos/$$PROTOGOPPATH \
		--go-grpc_out=$$(pwd)/$$PROTOGOPPATH \
		--go-grpc_opt=paths=import \
		--go-grpc_opt=module=github.com/vins7/module-protos/$$PROTOGOPPATH \
		$$entry;\
    done


