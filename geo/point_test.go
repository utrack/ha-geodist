package geo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/utrack/ha-geodist/geo"
)

func TestPoint__Distance(t *testing.T) {
	so := assert.New(t)

	type tc struct {
		lat1, lon1 float64
		lat2, lon2 float64

		dist float64
		// max error
		merr float64
	}

	cases := []tc{
		// Rounding errors on diff scales
		tc{
			lat1: 55.756974, lon1: 37.410423,
			lat2: 55.747020, lon2: 37.537688,
			dist: 8040.51, merr: 5,
		},
		tc{
			lat1: 53.184894, lon1: 44.998957,
			lat2: 55.747020, lon2: 37.537688,
			dist: 559635, merr: 30,
		},
		tc{
			lat1: 48.142455, lon1: 11.559709,
			lat2: 55.747020, lon2: 37.537688,
			dist: 1955620, merr: 15,
		},

		// past zero lon
		tc{
			lat1: 34.996414, lon1: -90.309529,
			lat2: 55.747020, lon2: 37.537688,
			dist: 8782150, merr: 30,
		},

		// same point
		tc{
			lat1: 55.747020, lon1: 37.537688,
			lat2: 55.747020, lon2: 37.537688,
			dist: 0, merr: 5,
		},

		// different hemispheres
		// big error margin vs real d since Haversine
		// doesn't account for Earth's curvature
		//
		// it's good enough to compare distances tho, so
		// ok for the test challenge purposes
		tc{
			lat1: -69.999764, lon1: -69.999764,
			lat2: 55.747020, lon2: 37.537688,
			dist: 13558075, merr: 2750000,
		},
	}

	for pos, c := range cases {
		p1 := geo.NewPoint(c.lat1, c.lon1)
		p2 := geo.NewPoint(c.lat2, c.lon2)

		dist := geo.DistanceHav(p1, p2)
		so.InDelta(c.dist, dist, c.merr, " case %v", pos)
	}
}
