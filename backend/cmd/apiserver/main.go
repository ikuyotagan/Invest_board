package main

import (
	"flag"
	"github.com/Artemchikus/api/internal/app/apiserver"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

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

	//tinkoff.NewClient()
	//defer tinkoff.CloseClient()
	//tinkoff.StreamingListener()

	if err := apiserver.Start(config); err != nil {
		log.Fatalln(err)
	}
}
