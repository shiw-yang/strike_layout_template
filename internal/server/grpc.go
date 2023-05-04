package server

import (
	v1 "strike_layout_template/api/helloworld/v1"
	"strike_layout_template/internal/conf"
	"strike_layout_template/internal/service"

	"google.golang.org/grpc"
)

// NewGRPCServer .
func NewGRPCServer(c *conf.Server, greeterService *service.GreeterService) *grpc.Server {
	grpcServer := grpc.NewServer()
	v1.RegisterGreeterServer(grpcServer, greeterService)

	return grpcServer
}
