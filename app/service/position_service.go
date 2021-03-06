package service

import (
	"errors"
	"math"
	"sync"

	"github.com/bernnabe/mp/app/model"
	"github.com/bernnabe/mp/app/repository"
)

type PositionServiceInterface interface {
	GetPosition(wg *sync.WaitGroup, distance model.Distance) (x, y float64, err error)
	TryGetSplitedPosition(wg *sync.WaitGroup) (x, y float64, err error)
	AddDistancePart(wg *sync.WaitGroup, distance model.Distance)
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
func (service *PositionService) TryGetSplitedPosition(wg *sync.WaitGroup) (x, y float64, err error) {
	defer wg.Done()

	model := model.Distance{
		Kenobi:    service.Repository.Get(kenobiKey),
		Skywalker: service.Repository.Get(skywalkerKey),
		Sato:      service.Repository.Get(satoKey),
	}

	xResult, yResult, err := getXY(model)

	if err == nil {
		return xResult, yResult, nil
	}

	return 0, 0, errors.New("Not enough information")
}

// AddDistancePart Intenta determinar la posición de la nave si es que ya conoce la posición de todos los satellites
func (service *PositionService) AddDistancePart(wg *sync.WaitGroup, distance model.Distance) {
	defer wg.Done()

	kenobi := service.Repository.Get(kenobiKey)
	skywalker := service.Repository.Get(skywalkerKey)
	sato := service.Repository.Get(satoKey)

	if kenobi == 0 {
		kenobi = distance.Kenobi
	}
	if skywalker == 0 {
		skywalker = distance.Skywalker
	}
	if sato == 0 {
		sato = distance.Sato
	}

	service.Repository.Add(kenobiKey, kenobi)
	service.Repository.Add(skywalkerKey, skywalker)
	service.Repository.Add(satoKey, sato)
}

//GetPosition Determina la posición de un punto en el plano r2 utilizando un sistema de ecuaciones
func (service *PositionService) GetPosition(wg *sync.WaitGroup, distance model.Distance) (x, y float64, err error) {
	defer wg.Done()

	return getXY(distance)
}

//getXY Determina en base a dos ecuaciones el punto X Y de interseccion con la tercera ecuación
//El método asume que la posición de la nave está en la intersección exacta de los tres satellites.
//
// Triangulación en el plano
// https://www.wolframalpha.com/input/?i=%28x-3%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B+%28x-6%29%5E2%2B%28y-10%29%5E2%3D3%5E2%3B+%28x-9%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B
func getXY(distance model.Distance) (x float64, y float64, err error) {
	kenobiPosition := satPosition{3, 3, float64(distance.Kenobi)}        //x1. y1. distance r1
	skywalkerPosition := satPosition{6, 10, float64(distance.Skywalker)} //x2. y2. distance r2
	satoPosition := satPosition{9, 3, float64(distance.Sato)}            //x2. y3. distance r3

	if distance.Kenobi == 0 || distance.Skywalker == 0 || distance.Sato == 0 {
		return 0, 0, errors.New("not enough information to determine position")
	}

	//Nombro las variables con K1...K6 ya que son sólo parametros para ser usados en la ultima ecuación.
	//Calculo los puntos que forman la recta resultante kenobi/skywalkerPosition
	k1, k2, k3 := getEqLine(kenobiPosition, skywalkerPosition)
	//Calculo los puntos que forman la recta resultante kenobi/sato
	k4, k5, k6 := getEqLine(kenobiPosition, satoPosition)

	//Usando las posiciones hago los calculos para determinar Y
	yResult := (((k1 * k6) / k4) - k3) /
		(k2 - ((k1 * k5) /
			k4))

	//Usando las posiciones hago los calculos para determinar X en funcion yResult
	xResult := (-k3 - (k2 * yResult)) / k1

	if math.IsNaN(xResult) || math.IsNaN(yResult) {
		return 0, 0, errors.New("the position cannot be determined")
	}

	//Devuelvo las coordenadas donde se intersectan las tres rectas
	return xResult, yResult, nil
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
