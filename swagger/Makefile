.PHONY: requirements kitchen-api

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/../third_party/googleapis"

all: requirements kitchen-api

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

kitchen-api:
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --swagger_out=json_names_for_fields=true:. as/external/api/kitchen.proto
