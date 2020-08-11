package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bernnabe/mp/app/model"
	"github.com/bernnabe/mp/app/service"
)

// PostTopSecretSplit Recibe el mensaje en partes e intenta devolver el mensaje original y la posicion
func PostTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	distanceService := service.NewDistanceService()
	messageService := service.NewMessageService()

	x, y, getPositionError := distanceService.GetPosition(request.Distance.Kenobi, request.Distance.Skywalker, request.Distance.Sato)
	message, getMessagerror := messageService.GetMessage(request.Message.Kenobi, request.Message.Skywalker, request.Message.Sato)

	if getPositionError != nil || getMessagerror != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		return
	}

	json.NewEncoder(w).Encode(model.TopSecretResponse{
		Message:  message,
		Position: model.Position{X: x, Y: y},
	})
}

// PostTopSecret Recibe el mensaje desde los tres satelites y devuelve el mensaje tal cual se genero en el sender y su posición
func PostTopSecret(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	messageService := service.NewMessageService()
	distanceService := service.NewDistanceService()

	x, y, getPositionError := distanceService.GetPosition(request.Distance.Kenobi, request.Distance.Skywalker, request.Distance.Sato)
	message, getMessagerror := messageService.GetMessage(request.Message.Kenobi, request.Message.Skywalker, request.Message.Sato)

	if getPositionError != nil || getMessagerror != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		return
	}
	json.NewEncoder(w).Encode(model.TopSecretResponse{
		Message:  message,
		Position: model.Position{X: x, Y: y},
	})

}

// Home Muestra un mensaje
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Message Api")
}
