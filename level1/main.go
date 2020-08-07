package main

import (
	"bytes"
	"errors"
	"strings"
)

func main() {

}

type Position struct {
	X        float64
	Y        float64
	Distance float64
}

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

func getPosition(kenobiDistance, skywalkerDistance, satoDistance float32) (x, y float32) {
	/*
		Triangulación en el plano
		https://www.wolframalpha.com/input/?i=%28x-3%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B+%28x-6%29%5E2%2B%28y-10%29%5E2%3D3%5E2%3B+%28x-9%29%5E2%2B%28y-3%29%5E2%3D5%5E2%3B
	*/

	var satellitesPosition = []Position{}

	satellitesPosition = append(satellitesPosition,
		Position{3, 3, float64(kenobiDistance)},
	)

	// Position{6, 10, float64(skywalkerDistance)},
	// Position{9, 3, float64(satoDistance)})

	result := []Position{}

	for _, position := range satellitesPosition {
		for ix := 1; ix <= 15; ix++ {
			for iy := 1; iy <= 15; iy++ {
				// x := float64(ix)
				// y := float64(iy)

				zero := 0

				if zero == 0 {
					result = append(result, Position{X: float64(ix), Y: float64(iy), Distance: position.Distance})
					println("x:" + string(ix))
				}

				println(int(zero))
			}
		}
	}

	if len(result) != 2 {
		return 0, 0
	}

	return float32(result[0].X), float32(result[1].Y)
}

// GetMessage Procesa los mensajes recibidos en cada satelite
// input: Mensajes tal cual se reciben en cada satelite
// output:  Mensaje tal cual fué enviado desde el emisor.
func GetMessage(kenobiMessages, skywalkerMessages, satoMessages []string) (message string, err error) {
	validStructures := len(kenobiMessages) == len(skywalkerMessages) && len(kenobiMessages) == len(satoMessages)

	if !validStructures {
		return "", errors.New("Message bad formed")
	}

	return processPartsMessage(kenobiMessages, skywalkerMessages, satoMessages), nil
}

func processPartsMessage(kenobiMessages, skywalkerMessages, satoMessages []string) string {
	keys := make(map[string]bool)
	var buffer bytes.Buffer

	var addPart = func(part string, keys map[string]bool, buffer *bytes.Buffer) {
		if part != "" {
			if _, value := keys[part]; !value {
				keys[part] = true
				buffer.WriteString(part + " ")
			}
		}
	}

	for i := range kenobiMessages {
		addPart(kenobiMessages[i], keys, &buffer)
		addPart(skywalkerMessages[i], keys, &buffer)
		addPart(satoMessages[i], keys, &buffer)
	}

	return strings.TrimRight(buffer.String(), " ")
}
