package main

import (
	"flag"
	"fmt"
	"strike_layout_template/internal/conf"
	"strike_layout_template/internal/server"

	"github.com/spf13/viper"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	// init
	flag.Parse()
	c := readConfig()
	// start Server
	err := server.NewServer(c.Server)
	if err != nil {
		panic(err)
	}
	//	stop channel
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
	fmt.Println(&c)
	return &c
}
