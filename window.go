package main

import (
	"sort"

	"github.com/utrack/ha-geodist/geo"
	"github.com/utrack/ha-geodist/pstream"
)

// pointWithDist is a point that has a distance to some other arbitrary point
// stored inside it.
type pointWithDist struct {
	pstream.Point
	dist float64
}

type window struct {
	refPoint geo.Point
	pts      []pointWithDist
	count    int

	inverseSort bool
}

// newWindow returns a new window that stores <count> closest or farthest points
// relative to the reference point.
func newWindow(count int, farthest bool, ref geo.Point) *window {
	return &window{
		refPoint:    ref,
		pts:         make([]pointWithDist, 0, count),
		count:       count,
		inverseSort: farthest,
	}
}

// Add inserts another point to the window.
func (w *window) Add(p pstream.Point) {

	dist := geo.DistanceHav(p.Point, w.refPoint)

	if len(w.pts) >= w.count {
		lastAccPointDist := w.pts[len(w.pts)-1].dist
		// this point is farther than last and we've already filled all slots
		if !w.inverseSort && lastAccPointDist < dist {
			return
		}
		// or closer than last if we're storing farthest ones
		if w.inverseSort && lastAccPointDist > dist {
			return
		}
	}

	pWithDist := newPointDist(p, dist)
	w.pts = pointDistInsertSort(w.pts, pWithDist, w.inverseSort)
	if len(w.pts) > w.count {
		w.pts = w.pts[:w.count]
	}
}

// GetAccum returns this window's accumulator.
func (w *window) GetAccum() []pointWithDist {
	ret := make([]pointWithDist, len(w.pts))
	copy(ret, w.pts)
	return ret
}

func newPointDist(p pstream.Point, d float64) pointWithDist {
	return pointWithDist{
		Point: p,
		dist:  d,
	}
}

// pointDistInsertSort inserts the new point into a sorted list of points,
// preserving the sorting.
// inverse is true if points are sorted from farthest to closest.
func pointDistInsertSort(data []pointWithDist, el pointWithDist, inverse bool) []pointWithDist {
	index := sort.Search(len(data), func(i int) bool {
		if inverse {
			return data[i].dist < el.dist
		}
		return data[i].dist > el.dist
	})

	data = append(data, pointWithDist{})

	copy(data[index+1:], data[index:])
	data[index] = el

	return data
}
