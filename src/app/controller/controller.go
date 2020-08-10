package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bernnabe/mp/app/model"
	"github.com/bernnabe/mp/app/service"
)

func PostTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var error error
	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	x, y, error := service.GetPosition(request.Distance.Kenobi, request.Distance.Skywalker, request.Distance.Sato)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		return
	}

	message, error := service.GetMessage(request.Message.Kenobi, request.Message.Skywalker, request.Message.Sato)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		return
	}

	json.NewEncoder(w).Encode(model.TopSecretResponse{
		Message:  message,
		Position: model.Position{X: x, Y: y},
	})
}

func PostTopSecret(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var error error
	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	x, y, error := service.GetPosition(request.Distance.Kenobi, request.Distance.Skywalker, request.Distance.Sato)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		return
	}

	message, error := service.GetMessage(request.Message.Kenobi, request.Message.Skywalker, request.Message.Sato)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		return
	}

	json.NewEncoder(w).Encode(model.TopSecretResponse{
		Message:  message,
		Position: model.Position{X: x, Y: y},
	})

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Message Api")
}
