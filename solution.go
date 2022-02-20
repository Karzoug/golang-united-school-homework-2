package square

import "math"

type SidesNumber int

const (
	SidesTriangle SidesNumber = 3
	SidesSquare               = 4
	SidesCircle               = 0
)

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {
	switch sidesNum {
	case SidesTriangle:
		return ((sideLen * sideLen) * math.Sqrt(3)) / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	default:
		return 0
	}
}
