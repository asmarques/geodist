package geodist

import "math"

const r = 6371

func haversin(theta float64) float64 {
	return 0.5 * (1 - math.Cos(theta))
}

// HaversineDistance returns the geographical distance in km between the points p1 and p2 using the Haversine formula.
// The surface of the Earth is approximated by a sphere with a radius of 6371 km.
func HaversineDistance(p1 Point, p2 Point) float64 {
	phi1 := toRadians(p1.Lat)
	phi2 := toRadians(p2.Lat)
	lambda1 := toRadians(p1.Long)
	lambda2 := toRadians(p2.Long)

	return 2 * r * math.Asin(math.Sqrt(haversin(phi2-phi1)+math.Cos(phi1)*math.Cos(phi2)*haversin(lambda2-lambda1)))
}
