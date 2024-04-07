package area_test

import (
	"lazarok09/area/area"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Should calculate the square area", func(t *testing.T) {

		var expected int = 10000

		result := area.SquareArea(
			area.SquareAreaCalculation{
				Side: 10,
			},
		)

		if expected != result {
			t.Fatalf("The expected value %d was different than the result %d", expected, result)
		}

	})
	t.Run("Should calculate the circle area", func(t *testing.T) {

		var expected float64 = 50.265482
		result := area.CircleArea(
			area.CircleAreaCalculation{Circunference: 8},
		)

		if math.Abs(expected-result) > 0.000001 {
			t.Fatalf("The expected value %f was different than the result %f", expected, result)
		}
	})
}
