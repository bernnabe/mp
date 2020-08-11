package main

import (
	"flag"

	"github.com/bernnabe/mp/app"
)

func main() {
	var configFilePath string
	var serverPort string

	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.StringVar(&serverPort, "server_port", "8080", "port on which server runs")
	flag.Parse()

	application := app.New(configFilePath)

	// start http server
	application.Start(serverPort)
}
