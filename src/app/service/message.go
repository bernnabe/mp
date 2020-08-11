package service

import (
	"bytes"
	"errors"
	"strings"
)

type MessageInterface interface {
	GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (message string, err error)
}

type MessageService struct {
}

// New : build new Service
func NewMessageService() MessageInterface {
	return &MessageService{}
}

// GetMessage Procesa los mensajes recibidos en cada satelite
// input: Mensajes tal cual se reciben en cada satelite
// output: Mensaje tal cual fué enviado desde el emisor.
func (service *MessageService) GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (message string, err error) {
	validStructures := len(kenobiMessages) == len(skywalkerMessages) && len(kenobiMessages) == len(satoMessages)

	if !validStructures {
		return "", errors.New("message isn't well formed")
	}

	return processMessageParts(kenobiMessages, skywalkerMessages, satoMessages), nil
}

func processMessageParts(kenobiMessages, skywalkerMessages, satoMessages []string) string {
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
