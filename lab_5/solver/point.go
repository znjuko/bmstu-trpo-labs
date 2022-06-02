package solver

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	Name string
	X    float64
	Y    float64
	Z    float64
}

func ParsePoint(point string) (p Point) {
	split := strings.Split(strings.TrimSpace(point), "")

	values := []float64{0, 0, 0}
	pos := 0
	for i := 2; i < len(split); i++ {
		var currentCoord string
		for split[i] != "," && split[i] != ")" {
			currentCoord += split[i]
			i++
		}

		values[pos], _ = strconv.ParseFloat(currentCoord, 64)
		pos++
	}

	return Point{
		Name: split[0],
		X:    values[0],
		Y:    values[1],
		Z:    values[2],
	}
}

func CountLength(first, second Point) float64 {
	return math.Sqrt(
		(first.X-second.X)*(first.X-second.X) +
			(first.Y-second.Y)*(first.Y-second.Y) +
			(first.Z-second.Z)*(first.Z-second.Z),
	)
}
