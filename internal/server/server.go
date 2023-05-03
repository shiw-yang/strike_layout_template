package server

import (
	"log"
	"strike_layout_template/internal/conf"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// Server .
type Server struct {
	id       string
	name     string
	version  string
	metadata map[string]string

	logger *log.Logger

	httpServer *gin.Engine
	grpcServer *grpc.Server
}

// Option .
type Option func(o *Server)

// NewServer 配置server
func NewServer(hs *gin.Engine, gs *grpc.Server) Option {
	return func(server *Server) {
		server.httpServer = hs
		server.grpcServer = gs
	}

}

// ID .
func ID(id string) Option {
	return func(server *Server) {
		server.id = id
	}
}

// Name .
func Name(name string) Option {
	return func(server *Server) {
		server.name = name
	}
}

// Version .
func Version(version string) Option {
	return func(server *Server) {
		server.version = version
	}
}

// Logger .
func Logger(logger *log.Logger) Option {
	return func(server *Server) {
		server.logger = logger
	}
}

// Metadata .
func Metadata(metadata map[string]string) Option {
	return func(server *Server) {
		server.metadata = metadata
	}
}

// RunServer 启动server
func (server *Server) RunServer(c *conf.Server) error {
	//TODO: get url
	err := server.httpServer.Run(c.Http.Addr)
	return err
}

// New .
func New(opts ...Option) *Server {
	server := &Server{}
	for _, opt := range opts {
		opt(server)
	}
	return server
}
