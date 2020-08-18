package service

import (
	"sync"
	"testing"

	"github.com/bernnabe/mp/app/model"
	"github.com/bernnabe/mp/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestWhenMessageIsIncomplete(t *testing.T) {
	kenobi := []string{"", "este", "es", "un", "mensaje"}
	skywalker := []string{"este", "", "un", "mensaje"}
	sato := []string{"", "", "es", "", "mensaje"}

	model := model.Message{
		Kenobi:    kenobi,
		Skywalker: skywalker,
		Sato:      sato,
	}
	total := 1
	var wg sync.WaitGroup
	wg.Add(total)

	var result string
	var err error

	messageService := NewMessageService(repository.NewMessageRepository())

	go func() { result, err = messageService.GetMessage(&wg, model) }()

	wg.Wait()

	assert.NotNil(t, err)
	assert.EqualValues(t, "", result)
}

func TestWhenMessageIsWellFormed(t *testing.T) {
	kenobi := []string{"", "este", "es", "un", "mensaje"}
	skywalker := []string{"este", "", "", "un", "mensaje"}
	sato := []string{"", "", "es", "", "mensaje"}
	model := model.Message{
		Kenobi:    kenobi,
		Skywalker: skywalker,
		Sato:      sato,
	}
	total := 1
	var wg sync.WaitGroup
	wg.Add(total)

	var result string
	var err error

	messageService := NewMessageService(repository.NewMessageRepository())

	go func() { result, err = messageService.GetMessage(&wg, model) }()
	wg.Wait()

	assert.Nil(t, err)
	assert.EqualValues(t, "este es un mensaje", result)
}

func TestWhenRandomHelloMessage(t *testing.T) {
	kenobi := []string{"Hola1", "", "", "", ""}
	skywalker := []string{"", "", "", "", "Hola5"}
	sato := []string{"", "Hola2", "Hola3", "Hola4", ""}
	model := model.Message{
		Kenobi:    kenobi,
		Skywalker: skywalker,
		Sato:      sato,
	}
	total := 1
	var wg sync.WaitGroup
	wg.Add(total)

	var result string
	var err error

	messageService := NewMessageService(repository.NewMessageRepository())
	go func() { result, err = messageService.GetMessage(&wg, model) }()

	wg.Wait()

	assert.Nil(t, err)
	assert.EqualValues(t, "Hola1 Hola2 Hola3 Hola4 Hola5", result)
}
