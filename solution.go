package square

import (
	"math"
)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	case >0:
		return (sidesNum * math.Pow(sideLen, 2))/(4*math.Tan(360/2*sidesNum))
	default:
		return 0
	}
}