package service

import (
	"errors"
	"math"

	"github.com/bernnabe/mp/app/repository"
)

type PositionServiceInterface interface {
	GetPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error)
	TryGetSplitedPosition() (x, y float64, err error)
	AddDistancePart(kenobiDistance, skywalkerDistance, satoDistance float64)
	ClearParts()
}

type PositionService struct {
	Repository repository.PositionRepositoryInterface
}

// New : build new Service
func NewPositionService(repository repository.PositionRepositoryInterface) PositionServiceInterface {
	return &PositionService{
		Repository: repository,
	}
}

type satPosition struct {
	X        float64
	Y        float64
	Distance float64
}

// TryGetSplitedPosition Intenta determinar la posición de la nave si es que ya conoce la posición de todos los satellites
func (service *PositionService) TryGetSplitedPosition() (x, y float64, err error) {
	kenobi := service.Repository.Get(kenobiKey)
	skywalker := service.Repository.Get(skywalkerKey)
	sato := service.Repository.Get(satoKey)

	xResult, yResult, err := service.GetPosition(kenobi, skywalker, sato)

	if err == nil {
		return xResult, yResult, nil
	}

	return 0, 0, errors.New("Not enough information")
}

// AddDistancePart Intenta determinar la posición de la nave si es que ya conoce la posición de todos los satellites
func (service *PositionService) AddDistancePart(kenobiDistance, skywalkerDistance, satoDistance float64) {
	kenobi := service.Repository.Get(kenobiKey)
	skywalker := service.Repository.Get(skywalkerKey)
	sato := service.Repository.Get(satoKey)

	if kenobi == 0 {
		kenobi = kenobiDistance
	}
	if skywalker == 0 {
		skywalker = skywalkerDistance
	}
	if sato == 0 {
		sato = satoDistance
	}

	service.Repository.Add(kenobiKey, kenobi)
	service.Repository.Add(skywalkerKey, skywalker)
	service.Repository.Add(satoKey, sato)
}

//GetPosition Determina la posición de un punto en el plano r2 utilizando un sistema de ecuaciones
//El método asume que la posición de la nave está en la intersección exacta de los tres satellites.
//
// Triangulación en el plano
// https://www.wolframalpha.com/input/?i=%28x-3%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B+%28x-6%29%5E2%2B%28y-10%29%5E2%3D3%5E2%3B+%28x-9%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B
//
func (service *PositionService) GetPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error) {
	kenobiPosition := satPosition{3, 3, float64(kenobiDistance)}        //x1. y1. distance r1
	skywalkerPosition := satPosition{6, 10, float64(skywalkerDistance)} //x2. y2. distance r2
	satoPosition := satPosition{9, 3, float64(satoDistance)}            //x2. y3. distance r3

	if kenobiDistance == 0 || skywalkerDistance == 0 || satoDistance == 0 {
		return 0, 0, errors.New("not enough information to determine position")
	}

	xResult, yResult := getXY(kenobiPosition, skywalkerPosition, satoPosition)

	if math.IsNaN(xResult) || math.IsNaN(yResult) {
		return 0, 0, errors.New("the position cannot be determined")
	}

	return xResult, yResult, nil
}

//getXY Determina en base a dos ecuaciones el punto X Y de interseccion con la tercera ecuación
func getXY(kenobiPosition, skywalkerPosition, satoPosition satPosition) (x float64, y float64) {

	//Calculo los puntos que forman la recta resultante kenobi/skywalkerPosition
	k1, k2, k3 := getEqLine(kenobiPosition, skywalkerPosition)
	//Calculo los puntos que forman la recta resultante kenobi/sato
	k4, k5, k6 := getEqLine(kenobiPosition, satoPosition)

	//Usando las posiciones hago los calculos para determinar Y
	yResult := (((k1 * k6) / k4) - k3) /
		(k2 -
			((k1 * k5) /
				k4))

	//Usando las posiciones hago los calculos para determinar X
	xResult := (-k3 - (k2 * yResult)) / k1

	return xResult, yResult
}

func getEqLine(source, target satPosition) (float64, float64, float64) {
	//Igualo la ecuación de la recta de source y target
	k1 := (-2 * source.X) + (2 * target.X)
	k2 := (-2 * source.Y) + (2 * target.Y)

	k3 := math.Pow(source.X, 2) +
		math.Pow(source.Y, 2) -
		math.Pow(source.Distance, 2) -
		math.Pow(target.X, 2) -
		math.Pow(target.Y, 2) +
		math.Pow(target.Distance, 2)

	return k1, k2, k3
}

func (service *PositionService) ClearParts() {
	service.Repository.Clear()
}
