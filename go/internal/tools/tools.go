//go:build tools
// +build tools

package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "google.golang.org/protobuf/protoc-gen-go"
)
