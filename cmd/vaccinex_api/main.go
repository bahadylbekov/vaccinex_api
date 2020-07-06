package main

import (
	"flag"
	"log"

	"github.com/bahadylbekov/vaccinex_api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/vaccinex.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
