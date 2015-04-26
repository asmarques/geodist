package geodist

import (
	"math"
	"testing"
)

func TestHaversine(t *testing.T) {
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
			111.194927,
		},
		{
			Point{0, 0},
			Point{0.5, 179.5},
			19936.460608,
		},
		{
			Point{38.781311, -9.135918},
			Point{37.618817, -122.375427},
			9121.094071,
		},
		{
			Point{-43.489443, 172.534444},
			Point{43.302061, -8.377255},
			19938.534556,
		},
		{
			Point{0, 0},
			Point{0.5, 179.7},
			19950.249787,
		},
		{
			Point{38.781389, -9.135833},
			Point{-38.781389, 170.864167},
			20015.086796,
		},
	}

	e := 0.000001

	for _, test := range tests {
		d := HaversineDistance(test.a, test.b)

		if math.Abs(d-test.d) > e {
			t.Errorf("Expected distance between %v and %v to be %.6f, got %.6f", test.a, test.b, test.d, d)
		}
	}
}
