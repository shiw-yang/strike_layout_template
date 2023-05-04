package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strike_layout_template/internal/conf"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
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

	httpSrv    *http.Server
	ginServer  *gin.Engine
	grpcServer *grpc.Server
}

// Option .
type Option func(o *Server)

// NewServer 配置server
func NewServer(hs *gin.Engine, gs *grpc.Server) Option {
	return func(server *Server) {
		server.ginServer = hs
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

// Metadata .
func Metadata(metadata map[string]string) Option {
	return func(server *Server) {
		server.metadata = metadata
	}
}

// RunServer 启动server
func (server *Server) RunServer(c *conf.Server) error {
	//TODO: use errgroup to Run http and grpc server
	sctx := context.Background()
	eg, ctx := errgroup.WithContext(sctx)
	wg := sync.WaitGroup{}

	wg.Add(1)
	// http server
	eg.Go(func() error {
		wg.Done()
		server.httpSrv = &http.Server{
			Addr:    c.Http.Addr,
			Handler: server.ginServer,
		}
		return server.httpSrv.ListenAndServe()
	})
	wg.Add(1)
	// grpc server
	eg.Go(func() error {
		wg.Done()
		l, err := net.Listen("tcp", ":"+c.Grpc.Port)
		if err != nil {
			return err
		}
		if err := server.grpcServer.Serve(l); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	wg.Wait()
	// wait for exit signal
	ch := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	signal.Notify(ch, sigs...)
	eg.Go(func() error {
		select {
		case <-ch:
			return server.stop()
		case <-ctx.Done():
			return nil
		}
	})
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// New .
func New(opts ...Option) *Server {
	server := &Server{}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

func (server *Server) stop() error {
	server.grpcServer.GracefulStop()
	server.httpSrv.Shutdown(context.Background())
	return nil
}
