package geodist

import (
	"fmt"
	"math"
)

// WGS-84 ellipsoid
const (
	a float64 = 6378137
	f float64 = 1 / 298.257223563
	b float64 = 6356752.314245
)

const (
	epsilon       = 1e-12
	maxIterations = 200
)

// VincentyDistance returns the geographical distance in km between the points p1 and p2 using Vincenty's inverse formula.
// The surface of the Earth is approximated by the WGS-84 ellipsoid.
// This method may fail to converge for nearly antipodal points.
func VincentyDistance(p1 Point, p2 Point) (float64, error) {
	if p1.Lat == p2.Lat && p1.Long == p2.Long {
		return 0, nil
	}

	U1 := math.Atan((1 - f) * math.Tan(toRadians(p1.Lat)))
	U2 := math.Atan((1 - f) * math.Tan(toRadians(p2.Lat)))
	L := toRadians(p2.Long - p1.Long)
	sinU1 := math.Sin(U1)
	cosU1 := math.Cos(U1)
	sinU2 := math.Sin(U2)
	cosU2 := math.Cos(U2)
	lambda := L

	result := math.NaN()

	for i := 0; i < maxIterations; i++ {
		curLambda := lambda
		sinSigma := math.Sqrt(math.Pow(cosU2*math.Sin(lambda), 2) +
			math.Pow(cosU1*sinU2-sinU1*cosU2*math.Cos(lambda), 2))
		cosSigma := sinU1*sinU2 + cosU1*cosU2*math.Cos(lambda)
		sigma := math.Atan2(sinSigma, cosSigma)
		sinAlpha := (cosU1 * cosU2 * math.Sin(lambda)) / math.Sin(sigma)
		cosSqrAlpha := 1 - math.Pow(sinAlpha, 2)
		cos2sigmam := 0.0
		if cosSqrAlpha != 0 {
			cos2sigmam = math.Cos(sigma) - ((2 * sinU1 * sinU2) / cosSqrAlpha)
		}
		C := (f / 16) * cosSqrAlpha * (4 + f*(4-3*cosSqrAlpha))
		lambda = L + (1-C)*f*sinAlpha*(sigma+C*sinSigma*(cos2sigmam+C*cosSigma*(-1+2*math.Pow(cos2sigmam, 2))))

		if math.Abs(lambda-curLambda) < epsilon {
			uSqr := cosSqrAlpha * ((math.Pow(a, 2) - math.Pow(b, 2)) / math.Pow(b, 2))
			k1 := (math.Sqrt(1+uSqr) - 1) / (math.Sqrt(1+uSqr) + 1)
			A := (1 + (math.Pow(k1, 2) / 4)) / (1 - k1)
			B := k1 * (1 - (3*math.Pow(k1, 2))/8)

			deltaSigma := B * sinSigma * (cos2sigmam + (B/4)*(cosSigma*(-1+2*math.Pow(cos2sigmam, 2))-
				(B/6)*cos2sigmam*(-3+4*math.Pow(sinSigma, 2))*(-3+4*math.Pow(cos2sigmam, 2))))
			s := b * A * (sigma - deltaSigma)
			result = s / 1000

			break
		}
	}

	if math.IsNaN(result) {
		return result, fmt.Errorf("Failed to converge for %v and %v", p1, p2)
	}

	return result, nil
}
