package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/topsecret", receiveMessage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Api")
	fmt.Println("Endpoint Hit: homePage")
}

//Distance is a model
type Distance struct {
	Kenobi    float64 `json:kenobi`
	Skywalker float64 `json:skywalker`
	Sato      float64 `json:soto`
}

//Message is a model
type Message struct {
	Kenobi    []string `json:kenobi`
	Skywalker []string `json:skywalker`
	Sato      []string `json:soto`
}

//TopSecretRequest is a model
type TopSecretRequest struct {
	Distance Distance `json:"distance"`
	Message  Message  `json:"message"`
}

//Position is a model
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

//TopSecretResponse is a model
type TopSecretResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

func receiveMessage(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	var error error
	var request TopSecretRequest
	json.Unmarshal(reqBody, &request)

	x, y, error := getPosition(request.Distance.Kenobi, request.Distance.Skywalker, request.Distance.Sato)
	position := Position{X: x, Y: y}

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
		return
	}

	message, error := GetMessage(request.Message.Kenobi, request.Message.Skywalker, request.Message.Sato)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var response TopSecretResponse
	response.Message = message
	response.Position = position

	json.NewEncoder(w).Encode(response)
}

// GetMessage Procesa los mensajes recibidos en cada satelite
// input: Mensajes tal cual se reciben en cada satelite
// output: Mensaje tal cual fué enviado desde el emisor.
func GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (message string, err error) {
	validStructures := len(kenobiMessages) == len(skywalkerMessages) && len(kenobiMessages) == len(satoMessages)

	if !validStructures {
		return "", errors.New("message isn't well formed")
	}

	return processPartsMessage(kenobiMessages, skywalkerMessages, satoMessages), nil
}

func processPartsMessage(kenobiMessages, skywalkerMessages, satoMessages []string) string {
	keys := make(map[string]bool)
	var buffer bytes.Buffer

	var addPart = func(part string, keys map[string]bool, buffer *bytes.Buffer) {
		if _, value := keys[part]; !value && part != "" {
			keys[part] = true
			buffer.WriteString(part + " ")
		}
	}

	for i := range kenobiMessages {
		addPart(kenobiMessages[i], keys, &buffer)
		addPart(skywalkerMessages[i], keys, &buffer)
		addPart(satoMessages[i], keys, &buffer)
	}

	return strings.TrimRight(buffer.String(), " ")
}

// type Position struct {
// 	X        float64 `json:"x"`
// 	Y        float64 `json:"y"`
// 	Distance float64
// }

func getPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error) {
	return 0, 0, nil

	/*
		Triangulación en el plano
		https://www.wolframalpha.com/input/?i=%28x-3%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B+%28x-6%29%5E2%2B%28y-10%29%5E2%3D3%5E2%3B+%28x-9%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B
	*/

	/*
		y := position.Y - math.Sqrt(
			float64(kenobiDistance)-
				math.Pow(float64(ix), 2)+
				2*float64(ix)*float64(position.X)-
				math.Pow(float64(position.X), 2))


							zero := position.Distance -
						((math.Pow(x, 2) -
							2*x*position.X +
							math.Pow(position.X, 2)) +
							(math.Pow(y, 2) -
								2*y*position.Y +
								math.Pow(position.Y, 2)))

	*/

	var satellitesPosition = []Position{}

	satellitesPosition = append(satellitesPosition,
		Position{3, 3},
	)
	// satellitesPosition = append(satellitesPosition,
	// 	Position{3, 3, float64(kenobiDistance)},
	// )

	// Position{6, 10, float64(skywalkerDistance)},
	// Position{9, 3, float64(satoDistance)})

	result := []Position{}

	// for _, position := range satellitesPosition {
	// 	for ix := 1; ix <= 15; ix++ {
	// 		for iy := 1; iy <= 15; iy++ {
	// 			// x := float64(ix)
	// 			// y := float64(iy)

	// 			zero := 0

	// 			if zero == 0 {
	// 				result = append(result, Position{X: float64(ix), Y: float64(iy), Distance: position.Distance})
	// 				println("x:" + string(ix))
	// 			}

	// 			println(int(zero))
	// 		}
	// 	}
	// }

	if len(result) != 2 {
		return 0, 0, errors.New("the position can't be determined")
	}

	return result[0].X, result[1].Y, nil
}
