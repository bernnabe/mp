package service

import (
	"bytes"
	"errors"
	"strings"
	"sync"

	"github.com/bernnabe/mp/app/model"
	"github.com/bernnabe/mp/app/repository"
)

type MessageServiceInterface interface {
	GetMessage(wg *sync.WaitGroup, messages model.Message) (message string, err error)
	TryGetSplitedMessage(wg *sync.WaitGroup) (message string, err error)
	AddMessagePart(wg *sync.WaitGroup, messages model.Message)
	ClearParts()
}

type MessageService struct {
	Repository repository.MessageRepositoryInterface
}

// New : build new Service
func NewMessageService(repository repository.MessageRepositoryInterface) MessageServiceInterface {
	return &MessageService{
		Repository: repository,
	}
}

//TryGetSplitedMessage Intenta determinar el mensaje si todos los satellites ya informaron su parte
func (service *MessageService) TryGetSplitedMessage(wg *sync.WaitGroup) (m string, err error) {
	defer wg.Done()

	kenobi := service.Repository.Get(kenobiKey)
	skywalker := service.Repository.Get(skywalkerKey)
	sato := service.Repository.Get(satoKey)

	model := model.Message{
		Kenobi:    kenobi,
		Skywalker: skywalker,
		Sato:      sato,
	}

	message, err := processMessageParts(model)

	if err == nil {
		return message, nil
	}

	return "", errors.New("not enough information")
}

//AddMessagePart Agrega las partes de los mensajes
func (service *MessageService) AddMessagePart(wg *sync.WaitGroup, message model.Message) {
	defer wg.Done()

	kenobi := service.Repository.Get(kenobiKey)
	skywalker := service.Repository.Get(skywalkerKey)
	sato := service.Repository.Get(satoKey)

	kenobi = append(kenobi, message.Kenobi...)
	skywalker = append(skywalker, message.Skywalker...)
	sato = append(sato, message.Sato...)

	service.Repository.Add(kenobiKey, kenobi)
	service.Repository.Add(skywalkerKey, skywalker)
	service.Repository.Add(satoKey, sato)
}

// GetMessage Procesa los mensajes recibidos en cada satelite
// input: Mensajes tal cual se reciben en cada satelite
// output: Mensaje tal cual fué enviado desde el emisor.
func (service *MessageService) GetMessage(wg *sync.WaitGroup, message model.Message) (m string, err error) {
	defer wg.Done()

	return processMessageParts(message)
}

func processMessageParts(message model.Message) (string, error) {
	keys := make(map[string]bool)
	var buffer bytes.Buffer

	if !(len(message.Kenobi) == len(message.Skywalker) && len(message.Kenobi) == len(message.Sato)) {
		return "", errors.New("message has not been well received")
	}

	add := func(part string, keys map[string]bool, buffer *bytes.Buffer) {
		if _, value := keys[part]; !value && part != "" {
			keys[part] = true
			buffer.WriteString(part + " ")
		}
	}

	for i := range message.Kenobi {
		add(message.Kenobi[i], keys, &buffer)
		add(message.Skywalker[i], keys, &buffer)
		add(message.Sato[i], keys, &buffer)
	}

	return strings.TrimRight(buffer.String(), " "), nil
}

func (service *MessageService) ClearParts() {
	service.Repository.Clear()
}
