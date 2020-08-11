package service

import (
	"errors"
	"math"
)

type DistanceInterface interface {
	GetPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error)
}

type DistanceService struct {
}

// New : build new Service
func NewDistanceService() DistanceInterface {
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

	xResult := getX(satoPosition, skywalkerPosition, kenobiPosition)
	yResult := getY(xResult, satoPosition, skywalkerPosition, kenobiPosition)

	if math.IsNaN(xResult) || math.IsNaN(yResult) {
		return 0, 0, errors.New("the position cannot be determined")
	}

	return xResult, yResult, nil
}

func getY(xResult float64, kenobiPosition, skywalkerPosition, satoPosition SatPosition) float64 {

	result := []float64{}

	result = append(result, compareY(xResult, kenobiPosition)...)
	result = append(result, compareY(xResult, skywalkerPosition)...)
	result = append(result, compareY(xResult, satoPosition)...)

	return getPopularElement(result)
}

func compareY(xResult float64, satellite SatPosition) []float64 {
	result := []float64{}

	result = append(result, math.Sqrt(math.Pow(satellite.Distance, 2)-math.Pow(xResult-satellite.X, 2))+satellite.Y)
	result = append(result, satellite.Y-math.Sqrt(math.Pow(satellite.Distance, 2)-math.Pow(xResult-satellite.X, 2)))

	return result
}

func getX(kenobiPosition, skywalkerPosition, satoPosition SatPosition) float64 {
	result := []float64{}

	result = append(result, compareX(kenobiPosition, skywalkerPosition, satoPosition)...)
	result = append(result, compareX(kenobiPosition, satoPosition, skywalkerPosition)...)
	result = append(result, compareX(satoPosition, skywalkerPosition, kenobiPosition)...)

	return getPopularElement(result)
}

func compareX(source, target, reference SatPosition) []float64 {

	//Igualo la ecuación de posición de source y target para determinar uno de los puntos en comun con la tercera ecuación
	k1 := (-2 * source.X) + (2 * target.X)
	k2 := (-2 * source.Y) + (2 * target.Y)

	k3 := math.Pow(source.X, 2) +
		math.Pow(source.Y, 2) -
		math.Pow(source.Distance, 2) -
		math.Pow(target.X, 2) -
		math.Pow(target.Y, 2) +
		math.Pow(target.Distance, 2)

	if k2 == 0 {
		result := -k3 / k1
		return []float64{result, result}
	}

	k4 := k1 / -k2
	k5 := k3 / -k2

	//Resuelvo por ecuación cuadrática
	a := 1 + math.Pow(k4, 2)
	b := (-2 * reference.X) +
		2*k4*k5 -
		2*reference.Y*k4
	c := math.Pow(k5, 2) -
		2*reference.Y*k5 +
		math.Pow(reference.X, 2) +
		math.Pow(reference.Y, 2) -
		math.Pow(reference.Distance, 2)

	eqResult1 := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	eqResult2 := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)

	return []float64{eqResult1, eqResult2}
}

func getPopularElement(a []float64) (r float64) {
	count, tempCount := 1, 0
	popular := a[0]

	for i := 0; i < len(a)-1; i++ {
		temp := a[i]
		tempCount = 0

		for j := 1; j < len(a); j++ {
			if temp == a[j] {
				tempCount++
			}
		}

		if tempCount > count {
			popular = temp
			count = tempCount
		}
	}

	return popular
}
