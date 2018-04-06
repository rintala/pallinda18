package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func Average(xy []float64) float64 {
	avg := (xy[0] + xy[1]) / 2
	return avg
}
