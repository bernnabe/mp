package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

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
	PositionService service.PositionServiceInterface
}

// New : build new Controller
func NewController(messageService service.MessageServiceInterface, positionService service.PositionServiceInterface) Controller {
	return &GenericController{
		MessageService:  messageService,
		PositionService: positionService,
	}
}

// GetTopSecretSplit Intena devolver el mensaje que ha sido enviado en partes
func (controller *GenericController) GetTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	total := 2
	var wg sync.WaitGroup
	wg.Add(total)

	var message string
	var getMessageError error

	go func() { message, getMessageError = controller.MessageService.TryGetSplitedMessage(&wg) }()

	var x, y float64
	var getDistanceError error

	go func() { x, y, getDistanceError = controller.PositionService.TryGetSplitedPosition(&wg) }()

	wg.Wait()

	if getMessageError != nil || getDistanceError != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("not enough information")
		return
	}

	json.NewEncoder(w).Encode(model.TopSecretResponse{
		Message:  message,
		Position: model.Position{X: x, Y: y},
	})

	//Esto es una especie de ack del mensaje
	controller.MessageService.ClearParts()
	controller.PositionService.ClearParts()
}

// PostTopSecretSplit Recibe el mensaje en partes e intenta devolver el mensaje original y la posicion
func (controller *GenericController) PostTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	total := 2
	var wg sync.WaitGroup
	wg.Add(total)

	go controller.MessageService.AddMessagePart(&wg, request.Message)
	go controller.PositionService.AddDistancePart(&wg, request.Distance)

	wg.Wait()
	w.WriteHeader(http.StatusOK)
}

// PostTopSecret Recibe el mensaje desde los tres satelites y devuelve el mensaje tal cual se genero en el sender y su posición
func (controller *GenericController) PostTopSecret(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var request model.TopSecretRequest
	json.Unmarshal(reqBody, &request)

	total := 2
	var wg sync.WaitGroup
	wg.Add(total)

	var x, y float64
	var getPositionError error

	go func() {
		x, y, getPositionError = controller.PositionService.GetPosition(&wg, request.Distance)
	}()

	var message string
	var getMessagerror error

	fmt.Printf("*model*", &request.Message)

	go func() {
		message, getMessagerror = controller.MessageService.GetMessage(&wg, request.Message)
	}()

	wg.Wait()

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
