package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenMessageIsIncomplete(t *testing.T) {
	kenobi := []string{"", "este", "es", "un", "mensaje"}
	skywalker := []string{"este", "", "un", "mensaje"}
	sato := []string{"", "", "es", "", "mensaje"}

	var result string
	messageService := NewMessageService()
	result, err := messageService.GetMessage(kenobi, skywalker, sato)

	assert.NotNil(t, err)
	assert.EqualValues(t, "", result)
}

func TestWhenMessageIsWellFormed(t *testing.T) {
	kenobi := []string{"", "este", "es", "un", "mensaje"}
	skywalker := []string{"este", "", "", "un", "mensaje"}
	sato := []string{"", "", "es", "", "mensaje"}

	var result string
	messageService := NewMessageService()
	result, err := messageService.GetMessage(kenobi, skywalker, sato)

	assert.Nil(t, err)
	assert.EqualValues(t, "este es un mensaje", result)
}

func TestWhenRandomHelloMessage(t *testing.T) {
	kenobi := []string{"Hola1", "", "", "", ""}
	skywalker := []string{"", "", "", "", "Hola5"}
	sato := []string{"", "Hola2", "Hola3", "Hola4", ""}
	var result string
	messageService := NewMessageService()
	result, err := messageService.GetMessage(kenobi, skywalker, sato)

	assert.Nil(t, err)
	assert.EqualValues(t, "Hola1 Hola2 Hola3 Hola4 Hola5", result)
}
