package area

import "math"

type SquareAreaCalculation struct {
	Side int
}
type CircleAreaCalculation struct {
	Circunference int32
}

// The Circle Are are based in the all sides of a square multiplied by itself that means side power of ^4
func SquareArea(square SquareAreaCalculation) int {
	result := square.Side * square.Side * square.Side * square.Side
	return result
}

// The Circle are based in the  PI * (circunference / 2) ^2
func CircleArea(circle CircleAreaCalculation) float64 {
	radious := math.Pow(float64(circle.Circunference/2), 2)
	result := math.Pi * radious

	return result
}
