package service

import (
	"bytes"
	"errors"
	"strings"

	"github.com/bernnabe/mp/app/repository"
)

type Message interface {
	GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (message string, err error)
}

type MessageService struct {
	Repository repository.Repository
}

// New : build new Service
func NewMessageService(repository repository.Repository) Message {
	return &MessageService{
		Repository: repository,
	}
}

// GetMessage Procesa los mensajes recibidos en cada satelite
// input: Mensajes tal cual se reciben en cada satelite
// output: Mensaje tal cual fué enviado desde el emisor.
func (service *MessageService) GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (message string, err error) {

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
