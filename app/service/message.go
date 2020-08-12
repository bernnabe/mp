package service

import (
	"bytes"
	"errors"
	"strings"

	"github.com/bernnabe/mp/app/repository"
)

type MessageServiceInterface interface {
	GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (message string, err error)
	TryGetSplitedMessage() (message string, err error)
	AddMessagePart(kenobiMessages, skywalkerMessages, satoMessages []string)
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
func (service *MessageService) TryGetSplitedMessage() (m string, err error) {
	const (
		kenobiKey    = "kenobi"
		skywalkerKey = "skywalker"
		satoKey      = "sato"
	)

	kenobi := service.Repository.Get(kenobiKey)
	skywalker := service.Repository.Get(skywalkerKey)
	sato := service.Repository.Get(satoKey)

	message, err := service.GetMessage(kenobi, skywalker, sato)

	if err == nil {
		return message, nil
	}

	return "", errors.New("Not enoght information")
}

//AddMessagePart
func (service *MessageService) AddMessagePart(kenobiMessages, skywalkerMessages, satoMessages []string) {
	const (
		kenobiKey    = "kenobi"
		skywalkerKey = "skywalker"
		satoKey      = "sato"
	)

	kenobi := service.Repository.Get(kenobiKey)
	skywalker := service.Repository.Get(skywalkerKey)
	sato := service.Repository.Get(satoKey)

	kenobi = append(kenobi, kenobiMessages...)
	skywalker = append(skywalker, skywalkerMessages...)
	sato = append(sato, satoMessages...)

	service.Repository.Add(kenobiKey, kenobi)
	service.Repository.Add(skywalkerKey, skywalker)
	service.Repository.Add(satoKey, sato)
}

// GetMessage Procesa los mensajes recibidos en cada satelite
// input: Mensajes tal cual se reciben en cada satelite
// output: Mensaje tal cual fué enviado desde el emisor.
func (service *MessageService) GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (m string, err error) {

	if !(len(kenobiMessages) == len(skywalkerMessages) && len(kenobiMessages) == len(satoMessages)) {
		return "", errors.New("message isn't well formed")
	}

	return processMessageParts(kenobiMessages, skywalkerMessages, satoMessages), nil
}

func processMessageParts(kenobiMessages, skywalkerMessages, satoMessages []string) string {
	keys := make(map[string]bool)
	var buffer bytes.Buffer

	add := func(part string, keys map[string]bool, buffer *bytes.Buffer) {
		if _, value := keys[part]; !value && part != "" {
			keys[part] = true
			buffer.WriteString(part + " ")
		}
	}

	for i := range kenobiMessages {
		add(kenobiMessages[i], keys, &buffer)
		add(skywalkerMessages[i], keys, &buffer)
		add(satoMessages[i], keys, &buffer)
	}

	return strings.TrimRight(buffer.String(), " ")
}

func (service *MessageService) ClearParts() {
	service.Repository.Clear()
}
