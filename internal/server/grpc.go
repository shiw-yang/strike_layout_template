package server

import (
	"strike_layout_template/internal/conf"
	"strike_layout_template/internal/service"

	"google.golang.org/grpc"
)

// NewGRPCServer .
func NewGRPCServer(c *conf.Server, greeterService *service.GreeterService) *grpc.Server {
	// TODO: implement
	return &grpc.Server{}
}
