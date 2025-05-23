.PHONY: requirements common gw geo nc ns als ad ps as as-integration as-external-api fuota

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/../third_party/googleapis"

all: requirements common gw geo als ad ps nc ns fuota as as-integration as-external-api

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/golang/protobuf/protoc-gen-go

common:
	protoc -I=../protobuf -I=../protobuf/common --go_out=plugins=grpc,paths=source_relative:. common/common.proto

gw:
	protoc -I=../protobuf -I=../protobuf/gw --go_out=plugins=grpc,paths=source_relative:. gw/gw.proto

geo:
	protoc -I=../protobuf -I=../protobuf/geo --go_out=plugins=grpc,paths=source_relative:. geo/geo.proto

als:
	protoc -I=../protobuf -I=../protobuf/als --go_out=plugins=grpc,paths=source_relative:. als/als.proto

ad:
	protoc -I=../protobuf -I=../protobuf/ad --go_out=plugins=grpc,paths=source_relative:. ad/ad.proto
	
# ps:
# 	protoc -I=../protobuf -I=../protobuf/ps --go_out=plugins=grpc,paths=source_relative:. ps/ps.proto

nc:
	protoc -I=../protobuf -I=../protobuf/nc --go_out=plugins=grpc,paths=source_relative:. nc/nc.proto

ns:
	protoc -I=../protobuf -I=../protobuf/ns --go_out=plugins=grpc,paths=source_relative:. ns/profiles.proto
	protoc -I=../protobuf -I=../protobuf/ns --go_out=plugins=grpc,paths=source_relative:. ns/ns.proto

fuota:
	protoc -I=../protobuf -I=../protobuf/fuota --go_out=plugins=grpc,paths=source_relative:. fuota/fuota.proto

as:
	protoc -I=../protobuf -I=../protobuf/as --go_out=plugins=grpc,paths=source_relative:. as/as.proto

as-integration:
	protoc -I=../protobuf -I=../protobuf/as/integration --go_out=plugins=grpc,paths=source_relative:. as/integration/integration.proto

as-external-api:
	# gRPC
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/application.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/frameLog.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/device.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/deviceProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/deviceQueue.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/gateway.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/gatewayProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/internal.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/multicastGroup.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/networkServer.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/organization.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/profiles.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/serviceProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/user.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/data.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/alarm.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/payment.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/zone.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/notification.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/automation.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/kitchen.proto

	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/application.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/frameLog.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/device.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/deviceProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/deviceQueue.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/gateway.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/gatewayProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/internal.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/multicastGroup.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/networkServer.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/organization.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/profiles.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/serviceProfile.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/user.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/data.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/alarm.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/zone.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/notification.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/payment.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/automation.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/kitchen.proto

