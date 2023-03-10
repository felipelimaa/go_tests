package geometrics_test

import (
	"math"
	"testing"
	. "tests/geometrics"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{10, 12}
		expectedArea := float64(120)
		resultArea := r.Area()

		if resultArea != expectedArea {
			t.Fatalf("A área recebica %f é diferente da esperada %f.", resultArea, expectedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{10}
		expectedArea := float64(math.Pi * 100)
		resultArea := c.Area()

		if resultArea != expectedArea {
			t.Fatalf("A área recebica %f é diferente da esperada %f.", resultArea, expectedArea)
		}
	})
}
