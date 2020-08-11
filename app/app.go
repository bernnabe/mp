package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bernnabe/mp/app/controller"
	"github.com/bernnabe/mp/config"
	"github.com/gorilla/mux"
)

type App interface {
	// Start http server
	Start(serverPort string)
}

type ApiApplication struct {
	config *config.GeneralConfig
}

// New : build new ApiApplication
func New(configFilePaths ...string) App {
	return &ApiApplication{
		config: nil,
	}
}

// Start serve http server
func (app *ApiApplication) Start(serverPort string) {

	// init new handler
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", controller.Home)

	// myRouter.HandleFunc("/topsecret_split", postTopSecretSplit).Methods("GET")
	// myRouter.HandleFunc("/topsecret_split", getTopSecretSplit).Methods("POST")

	myRouter.HandleFunc("/topsecret", controller.PostTopSecret).Methods("POST")

	fmt.Println("Server Started at http://localhost:" + serverPort)

	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}
