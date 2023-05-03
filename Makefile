GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api -name *.proto)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)

.PHONY: init
# init envdocke
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/shiw-yang/strike/cmd/protoc-gen-gin@latest
	go install github.com/shiw-yang/strike/cmd/protoc-gen-go-errors@latest
	go install github.com/shiw-yang/strike/cmd/strike@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/favadi/protoc-go-inject-tag@latest
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
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       ./internal/conf/conf.proto
	protoc-go-inject-tag -input=./internal/conf/conf.pb.go


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
