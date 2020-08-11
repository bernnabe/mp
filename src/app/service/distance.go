package service

import (
	"errors"
	"fmt"
	"math"
)

type DistanceInterface interface {
	GetPositionEq(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error)
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

func (service *DistanceService) GetPositionEq(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error) {

	/*
		Triangulación en el plano
		https://www.wolframalpha.com/input/?i=%28x-3%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B+%28x-6%29%5E2%2B%28y-10%29%5E2%3D3%5E2%3B+%28x-9%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B
	*/

	//Resuelvo por sistema de 3 ecuaciones.
	//Asumo que la posición de la nave está en la intersección exacta de los tres satellites.

	kenobiPosition := SatPosition{3, 3, float64(kenobiDistance)}        //x1. y1. distance r1
	skywalkerPosition := SatPosition{6, 10, float64(skywalkerDistance)} //x2. y2. distance r2
	satoPosition := SatPosition{9, 3, float64(satoDistance)}            //x2. y3. distance r3

	x1, y1 := compare(kenobiPosition, skywalkerPosition, satoPosition)
	x2, y2 := compare(kenobiPosition, satoPosition, skywalkerPosition)
	x3, y3 := compare(satoPosition, skywalkerPosition, kenobiPosition)

	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(x3)
	fmt.Println(y3)

	return 0, 0, nil
}

func compare(source, target, reference SatPosition) (r1, r2 float64) {

	//Igualo la ecuación de posición de kenobi y skaywalker para determinar uno de los puntos en comun con la tercera ecuación
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
		return result, result
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

	return eqResult1, eqResult2
}

func (service *DistanceService) GetPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error) {

	/*
		Triangulación en el plano
		https://www.wolframalpha.com/input/?i=%28x-3%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B+%28x-6%29%5E2%2B%28y-10%29%5E2%3D3%5E2%3B+%28x-9%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B
	*/

	/*
		y := position.Y - math.Sqrt(
			float64(kenobiDistance)-
				math.Pow(float64(ix), 2)+
				2*float64(ix)*float64(position.X)-
				math.Pow(float64(position.X), 2))


							zero := position.Distance -
						((math.Pow(x, 2) -
							2*x*position.X +
							math.Pow(position.X, 2)) +
							(math.Pow(y, 2) -
								2*y*position.Y +
								math.Pow(position.Y, 2)))

	*/

	var satellitesPosition = []SatPosition{}

	satellitesPosition = append(satellitesPosition,
		SatPosition{3, 3, float64(kenobiDistance)},
		SatPosition{6, 10, float64(skywalkerDistance)},
		SatPosition{9, 3, float64(satoDistance)})

	result := map[SatPosition]int{}

	for _, position := range satellitesPosition {
		for ix := 1; ix <= 15; ix++ {
			for iy := 1; iy <= 15; iy++ {
				x := float64(ix)
				y := float64(iy)

				// zero := position.Y - math.Sqrt(
				// 	math.Pow(float64(position.Distance), 2)-
				// 		math.Pow(x, 2)+
				// 		2*x*position.X-
				// 		math.Pow(position.X, 2))

				zero := position.Distance -
					((math.Pow(x, 2) -
						2*x*position.X +
						math.Pow(position.X, 2)) +
						(math.Pow(y, 2) -
							2*y*position.Y +
							math.Pow(position.Y, 2)))

				if zero == 0 {
					key := SatPosition{X: float64(ix), Y: float64(iy), Distance: position.Distance}
					//result = append(result, SatPosition{X: float64(ix), Y: float64(iy), Distance: position.Distance})

					if _, value := result[key]; !value {
						result[key] = 1
					}

					println("x:" + string(ix))
				}

				println(int(zero))
			}
		}
	}

	if len(result) != 2 {
		return 0, 0, errors.New("the position can't be determined")
	}

	return 0, 0, nil
}
