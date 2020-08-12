package app

import (
	"net/http"

	"github.com/bernnabe/mp/app/controller"
	"github.com/bernnabe/mp/app/repository"
	"github.com/bernnabe/mp/app/service"
	"github.com/bernnabe/mp/config"
	"github.com/gorilla/mux"
	muxlogrus "github.com/pytimer/mux-logrus"
	log "github.com/sirupsen/logrus"
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

	// add logger middleware
	myRouter.Use(muxlogrus.NewLogger().Middleware)

	//Get Services
	messageService, distanceService := getServices()

	controller := controller.NewController(messageService, distanceService)

	//Api Routing map
	myRouter.HandleFunc("/", controller.Home)
	myRouter.HandleFunc("/topsecret", controller.PostTopSecret).Methods("POST")

	log.Info("Server Started at http://localhost:" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}

func getServices() (m service.Message, d service.Distance) {
	repository := repository.NewRepository()
	messageService := service.NewMessageService(repository)
	distanceService := service.NewDistanceService()

	return messageService, distanceService
}
