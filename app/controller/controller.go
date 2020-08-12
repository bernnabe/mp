package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bernnabe/mp/app/model"
	"github.com/bernnabe/mp/app/service"
)

type Controller interface {
	GetTopSecretSplit(w http.ResponseWriter, r *http.Request)
	PostTopSecretSplit(w http.ResponseWriter, r *http.Request)
	PostTopSecret(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
}

type GenericController struct {
	MessageService  service.MessageServiceInterface
	DistanceService service.DistanceServiceInterface
}

// New : build new Controller
func NewController(messageService service.MessageServiceInterface, distanceService service.DistanceServiceInterface) Controller {
	return &GenericController{
		MessageService:  messageService,
		DistanceService: distanceService,
	}
}

// GetTopSecretSplit Intena devolver el mensaje que ha sido enviado en partes
func (controller *GenericController) GetTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	message, getMessageError := controller.MessageService.TryGetSplitedMessage()
	x, y, getDistanceError := controller.DistanceService.TryGetSplitedDistance()

	if getMessageError != nil || getDistanceError != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("not enough information")
		return
	}

	json.NewEncoder(w).Encode(model.TopSecretResponse{
		Message:  message,
		Position: model.Position{X: x, Y: y},
	})

	//Esto es una especie de akc del mensaje
	controller.MessageService.ClearParts()
	controller.DistanceService.ClearParts()
}

// PostTopSecretSplit Recibe el mensaje en partes e intenta devolver el mensaje original y la posicion
func (controller *GenericController) PostTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	controller.MessageService.AddMessagePart(request.Message.Kenobi, request.Message.Skywalker, request.Message.Sato)
	controller.DistanceService.AddDistancePart(request.Distance.Kenobi, request.Distance.Skywalker, request.Distance.Sato)

	w.WriteHeader(http.StatusOK)
}

// PostTopSecret Recibe el mensaje desde los tres satelites y devuelve el mensaje tal cual se genero en el sender y su posición
func (controller *GenericController) PostTopSecret(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	x, y, getPositionError := controller.DistanceService.GetPosition(request.Distance.Kenobi, request.Distance.Skywalker, request.Distance.Sato)
	message, getMessagerror := controller.MessageService.GetMessage(request.Message.Kenobi, request.Message.Skywalker, request.Message.Sato)

	if getPositionError != nil || getMessagerror != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(model.TopSecretResponse{
		Message:  message,
		Position: model.Position{X: x, Y: y},
	})
}

// Home Muestra un mensaje
func (controller *GenericController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Message Api")
}
