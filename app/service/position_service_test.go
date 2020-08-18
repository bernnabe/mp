package service

import (
	"sync"
	"testing"

	"github.com/bernnabe/mp/app/model"
	"github.com/bernnabe/mp/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetPosition(t *testing.T) {
	positionService := NewPositionService(repository.NewPositionRepository())

	model := model.Distance{
		Kenobi:    5,
		Skywalker: 3,
		Sato:      5,
	}

	total := 1
	var wg sync.WaitGroup
	wg.Add(total)

	x, y, err := positionService.GetPosition(&wg, model)

	wg.Done()

	assert.Nil(t, err)
	assert.Equal(t, x, float64(6))
	assert.Equal(t, y, float64(7))
}
