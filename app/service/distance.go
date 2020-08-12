package service

import (
	"errors"
	"math"
)

type Distance interface {
	GetPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error)
}

type DistanceService struct {
}

// New : build new Service
func NewDistanceService() Distance {
	return &DistanceService{}
}

type SatPosition struct {
	X        float64
	Y        float64
	Distance float64
}

//GetPosition Determina la posición de un punto en el plano r2 utilizando un sistema de ecuaciones
//El método asume que la posición de la nave está en la intersección exacta de los tres satellites.
//
// Triangulación en el plano
// https://www.wolframalpha.com/input/?i=%28x-3%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B+%28x-6%29%5E2%2B%28y-10%29%5E2%3D3%5E2%3B+%28x-9%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B
//
func (service *DistanceService) GetPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error) {
	kenobiPosition := SatPosition{3, 3, float64(kenobiDistance)}        //x1. y1. distance r1
	skywalkerPosition := SatPosition{6, 10, float64(skywalkerDistance)} //x2. y2. distance r2
	satoPosition := SatPosition{9, 3, float64(satoDistance)}            //x2. y3. distance r3

	xResult, yResult := getXY(kenobiPosition, skywalkerPosition, satoPosition)

	if math.IsNaN(xResult) || math.IsNaN(yResult) {
		return 0, 0, errors.New("the position cannot be determined")
	}

	return xResult, yResult, nil
}

//getXY Determina en base a dos ecuaciones el punto X de interseccion con la tercera ecuación
func getXY(kenobiPosition, skywalkerPosition, satoPosition SatPosition) (p1 float64, p2 float64) {
	result := []float64{}

	k1, k2, k3 := 0, 1, 2
	k4, k5, k6 := 3, 4, 5

	//Primera Recta
	result = append(result, getEqLine(kenobiPosition, skywalkerPosition)...)

	//Segunda Recta
	result = append(result, getEqLine(kenobiPosition, satoPosition)...)

	//Punto x de la tercera recta
	x := (((result[k1] * result[k6]) / result[k4]) - result[k3]) /
		(result[k2] -
			((result[k1] * result[k5]) /
				result[k4]))

	//Punto y de la tercera recta
	y := (-result[k3] - (result[k2] * p1)) / result[k1]

	return x, y
}

func getEqLine(source, target SatPosition) []float64 {
	//Igualo la ecuación de la recta de source y target
	k1 := (-2 * source.X) + (2 * target.X)
	k2 := (-2 * source.Y) + (2 * target.Y)

	k3 := math.Pow(source.X, 2) +
		math.Pow(source.Y, 2) -
		math.Pow(source.Distance, 2) -
		math.Pow(target.X, 2) -
		math.Pow(target.Y, 2) +
		math.Pow(target.Distance, 2)

	return []float64{k1, k2, k3}
}
