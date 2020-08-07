package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPosition(t *testing.T) {
	x, y := getPosition(25, 9, 25)

	assert.Equal(t, x, 5)
	assert.Equal(t, y, 7)
}

func TestWhenMessageIsIncomplete(t *testing.T) {
	kenobi := []string{"", "este", "es", "un", "mensaje"}
	skywalker := []string{"este", "", "un", "mensaje"}
	sato := []string{"", "", "es", "", "mensaje"}

	var result string

	result = getMessage(kenobi, skywalker, sato)

	assert.EqualValues(t, "error", result)
}

func TestWhenMessageIsWellFormed(t *testing.T) {
	kenobi := []string{"", "este", "es", "un", "mensaje"}
	skywalker := []string{"este", "", "", "un", "mensaje"}
	sato := []string{"", "", "es", "", "mensaje"}

	var result string

	result = getMessage(kenobi, skywalker, sato)

	assert.EqualValues(t, "este es un mensaje", result)
}

func TestWhenRandomHelloMessage(t *testing.T) {
	kenobi := []string{"Hola1", "", "", "", ""}
	skywalker := []string{"", "", "", "", "Hola5"}
	sato := []string{"", "Hola2", "Hola3", "Hola4", ""}
	var result string

	result = getMessage(kenobi, skywalker, sato)

	assert.EqualValues(t, "Hola1 Hola2 Hola3 Hola4 Hola5", result)
}
