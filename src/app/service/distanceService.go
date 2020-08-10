package service

import (
	"errors"
	"math"
)

type SatPosition struct {
	X        float64
	Y        float64
	Distance float64
}

func GetPosition(kenobiDistance, skywalkerDistance, satoDistance float64) (x, y float64, err error) {

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

	result := []SatPosition{}

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
					result = append(result, SatPosition{X: float64(ix), Y: float64(iy), Distance: position.Distance})
					println("x:" + string(ix))
				}

				println(int(zero))
			}
		}
	}

	if len(result) != 2 {
		return 0, 0, errors.New("the position can't be determined")
	}

	return result[0].X, result[1].Y, nil
}
