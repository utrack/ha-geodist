package pstream

import (
	"github.com/pkg/errors"
	"github.com/utrack/ha-geodist/geo"
)

// Point is a pair of geo coords that has an ID attached.
type Point struct {
	geo.Point
	ID string
}

// Points iterates over the collection of Point.
type Points interface {
	Next() (Point, error)
}

var (
	// ErrNoPoints is returned by Points if there's no more
	// points to scan.
	ErrNoPoints = errors.New("collection is exhausted")
)
