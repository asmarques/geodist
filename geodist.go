// Implementation of the Haversine and Vincenty methods for calculating geographical
// distances between points on the surface of the Earth.
package geodist

import "math"

// Point represents the coordinates of a given point on Earth
type Point struct {
	// Latitude of the point in degrees
	Lat float64
	// Longitude of the point in degrees
	Long float64
}

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}
