package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strike_layout_template/internal/conf"
	"strike_layout_template/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	_ "go.uber.org/automaxprocs"
)

var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version  string
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "./data/configs/conf.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	// init
	flag.Parse()
	loggerInit()
	// logger.Println("start")
	c := readConfig()
	server, cancel, err := wireApp(c.Server, c.Data)
	if err != nil {
		panic(err)
	}
	defer cancel()
	if err := server.RunServer(c.Server); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}

func readConfig() *conf.Bootstrap {
	var c conf.Bootstrap
	viper.SetConfigFile(flagconf)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	UnmarshalErr := viper.Unmarshal(&c)
	if UnmarshalErr != nil {
		panic(UnmarshalErr)
	}
	return &c
}

func newApp(gs *grpc.Server, hs *gin.Engine) *server.Server {
	return server.New(
		server.ID(id), server.Name(Name),
		server.Metadata(map[string]string{}),
		server.NewServer(hs, gs),
	)
}

func loggerInit() log.Logger {
	// TODO: init logger
	return log.Logger{}
}
