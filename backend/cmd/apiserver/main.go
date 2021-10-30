package main

import (
	"flag"
	"log"

	"github.com/Artemchikus/api/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

// @title Invest board
// @version 0.0.1
// @description This is a sample server celler server.
// @contact.email support@swagger.io
//
// @host localhost:8080
// @BasePath /Invest_board

func init() {
	flag.StringVar(&configPath, "config_path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
