package service

import (
	"testing"

	"github.com/bernnabe/mp/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetPosition(t *testing.T) {
	positionService := NewPositionService(repository.NewPositionRepository())
	x, y, err := positionService.GetPosition(5, 3, 5)

	assert.Nil(t, err)
	assert.Equal(t, x, float64(6))
	assert.Equal(t, y, float64(7))
}
