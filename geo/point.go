package geo

import "math"

// Point represents a single set of coordinates on Earth.
type Point struct {
	lat, lon float64
}

// NewPoint returns a new Point with specified lat/lon coordinates in degrees.
func NewPoint(lat, lon float64) Point {
	return Point{
		lat: lat,
		lon: lon,
	}
}

// RadLat returns point's Latitude in radians.
func (p Point) RadLat() float64 {
	return p.lat * radPerDegree
}

// RadLon returns point's Longtitude in radians.
func (p Point) RadLon() float64 {
	return p.lon * radPerDegree
}

func (p Point) Lat() float64 {
	return p.lat
}

func (p Point) Lon() float64 {
	return p.lon
}

const (
	// one degree in radians
	radPerDegree = math.Pi / 180.0
	// Earth's radius in metres
	earthRadius = 6371e3
)

// DistanceHav calculates the distance between two points in metres.
// Haversine distance formula is used to calculate the distance.
func DistanceHav(p1, p2 Point) float64 {
	// http://www.movable-type.co.uk/scripts/latlong.html

	lat1 := p1.RadLat()
	lat2 := p2.RadLat()

	deltaLat := (lat2 - lat1)
	deltaLon := (p2.RadLon() - p1.RadLon())

	sqSinDLat := math.Pow(math.Sin(deltaLat/2), 2)
	sqSinDLon := math.Pow(math.Sin(deltaLon/2), 2)

	// left and right-hand sides of an eq for Haversine
	a := sqSinDLat + sqSinDLon*math.Cos(lat1)*math.Cos(lat2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
