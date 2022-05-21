package main 

import (
	"github.com/BurntSushi/toml"
	"flag"
	"log"
	"github.com/savaukr/restApiGolang/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.SringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New()
	if err :=s.Start(); err != nil {
		log.Fatal(err)
	}
}