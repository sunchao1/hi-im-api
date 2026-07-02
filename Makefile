.PHONY: test proto proto-grpc tidy lint

PROTOC ?= protoc
PROTO_DIR := proto
GEN_DIR := gen/go

test:
	go test ./...

tidy:
	go mod tidy

proto:
	@mkdir -p $(GEN_DIR)/im/v1
	cd $(PROTO_DIR) && $(PROTOC) --go_out=../$(GEN_DIR) --go_opt=paths=source_relative \
		im/v1/mesg.proto

proto-grpc: proto
	@mkdir -p $(GEN_DIR)/seq/v1
	cd $(PROTO_DIR) && $(PROTOC) --go_out=../$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=../$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		seq/v1/seq.proto

gen: proto-grpc

lint:
	go vet ./...
