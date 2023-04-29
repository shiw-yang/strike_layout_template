GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api -name *.proto)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/shiw-yang/strike/cmd/protoc-gen-gin@latest
	go install github.com/shiw-yang/strike/cmd/protoc-gen-go-errors@latest
	go install github.com/shiw-yang/strike/cmd/strike@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
  	       --go_out=paths=source_relative:./api \
  	       --gin_out=paths=source_relative:./api \
  	       --go-grpc_out=paths=source_relative:./api \
 	       --openapi_out=fq_schema_naming=true,default_response=false:. \
 	       $(API_PROTO_FILES)

.PHONY: build
# build todo
build:
	echo "todo"

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...
