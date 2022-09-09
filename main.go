package main

import (
	"flag"
	"log"

	rest "github.com/ericbutera/go-api/api"
	"github.com/ericbutera/go-api/config"
)

// TODO make cmd

var (
	path = flag.String("config_path", ".", "path to config")
	file = flag.String("config_file", "config.yaml", "config file name")
)

func main() {
	config, err := config.NewAppConfig(path, file)
	if err != nil {
		log.Fatal(err)
	}

	rest.Serve(config)
	// grpc.Serve() // TODO: move to different service
}
