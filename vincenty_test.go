package geodist

import (
	"math"
	"testing"
)

func TestVincentySuccess(t *testing.T) {
	tests := []struct {
		a, b Point
		d    float64
	}{
		{
			Point{0, 0},
			Point{0, 0},
			0,
		},
		{
			Point{0, 0},
			Point{0, 1},
			111.319491,
		},
		{
			Point{0, 0},
			Point{0.5, 179.5},
			19936.288579,
		},
		{
			Point{38.781311, -9.135918},
			Point{37.618817, -122.375427},
			9142.358572,
		},
		{
			Point{-43.489443, 172.534444},
			Point{43.302061, -8.377255},
			19942.780939,
		},
	}

	e := 0.000001

	for _, test := range tests {
		d, err := VincentyDistance(test.a, test.b)

		if err != nil {
			t.Errorf("Expected distance between %v and %v to converge to %.6f", test.a, test.b, test.d)
		}

		if math.Abs(d-test.d) > e {
			t.Errorf("Expected distance between %v and %v to be %.6f, got %.6f", test.a, test.b, test.d, d)
		}
	}
}

func TestVincentyFail(t *testing.T) {
	tests := []struct {
		a, b Point
	}{
		{
			Point{0, 0},
			Point{0.5, 179.7},
		},
		{
			Point{38.781389, -9.135833},
			Point{-38.781389, 170.864167},
		},
	}

	for _, test := range tests {
		_, err := VincentyDistance(test.a, test.b)

		if err == nil {
			t.Errorf("Expected distance between %v and %v to not converge", test.a, test.b)
		}
	}
}
