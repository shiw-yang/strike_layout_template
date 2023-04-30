package server

import "strike_layout_template/internal/conf"

// NewServer 优雅启动server
func NewServer(c *conf.Server) error {
	server, err := NewHTTPServer(c)
	if err != nil {
		return err
	}
	err2 := server.Run(c.Http.Addr)
	if err2 != nil {
		return err2
	}
	return nil
}
