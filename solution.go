package square

import "math"
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle snum =  3
	SidesSquare  snum = 4 
	SidesCircle snum = 0
)
type snum int

func CalcSquare(sideLen float64, sidesNum snum) float64 {
	switch sidesNum  {
	case SidesTriangle:
		return math.Sqrt(3) * sideLen* sideLen / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return sideLen / (math.Pi * 4)
	default:
		return sideLen
	}
}
