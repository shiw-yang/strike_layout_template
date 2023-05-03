package server

import (
	v1 "strike_layout_template/api/helloworld/v1"
	"strike_layout_template/internal/middleware"
	"strike_layout_template/internal/service"

	"github.com/gin-gonic/gin"

	"strike_layout_template/internal/conf"
)

// NewHTTPServer .
func NewHTTPServer(c *conf.Server, greeterService *service.GreeterService) *gin.Engine {
	httpServer := gin.New()
	if c.Http.Timeout != nil {
		httpServer.Use(middleware.HTTPResponseTimeoutMiddleware(c.Http.Timeout.AsDuration()))
	}
	httpServer.Use(gin.Recovery())
	registerHTTPService(httpServer)
	return httpServer
}

func registerHTTPService(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	v1.RegisterGreeterHTTPServer(engine, &service.GreeterService{})
}
