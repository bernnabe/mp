package model

//Distance is a model
type Distance struct {
	Kenobi    float64 `json:"kenobi"`
	Skywalker float64 `json:"skywalker"`
	Sato      float64 `json:"sato"`
}

//Message is a model
type Message struct {
	Kenobi    []string `json:"kenobi"`
	Skywalker []string `json:"skywalker"`
	Sato      []string `json:"sato"`
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
