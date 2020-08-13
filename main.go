package main

import (
	"flag"
	"os"

	"github.com/bernnabe/mp/app"
)

func main() {
	var configFilePath string
	var serverPort string = os.Getenv("PORT")

	if serverPort == "" {
		serverPort = "8080"
	}

	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.StringVar(&serverPort, "server_port", serverPort, "port on which server runs")
	flag.Parse()

	application := app.New(configFilePath)

	// start http server
	application.Start(serverPort)
}
