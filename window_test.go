package main

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/utrack/ha-geodist/geo"
	"github.com/utrack/ha-geodist/pstream"
)

func TestWindow__UnsortedInsert_GetClosest(t *testing.T) {
	so := assert.New(t)

	w := newWindow(5, false, geo.NewPoint(1, 0))

	w.Add(pstream.NewPoint("0", 10, 0))
	w.Add(pstream.NewPoint("1", -20, 0))
	w.Add(pstream.NewPoint("2", 4, 0))
	w.Add(pstream.NewPoint("3", 8, 0))
	w.Add(pstream.NewPoint("4", 1, 0))
	w.Add(pstream.NewPoint("5", 1, 0))
	w.Add(pstream.NewPoint("6", 1000, 0))
	w.Add(pstream.NewPoint("7", 25, 0))
	w.Add(pstream.NewPoint("8", 6, 0))

	got := w.GetAccum()

	expect := strings.Split("4,5,2,8,3", ",")

	for pos := range got {
		so.Equal(expect[pos], got[pos].ID)
	}
}

func TestWindow__UnsortedInsert_GetFarthest(t *testing.T) {
	so := assert.New(t)

	w := newWindow(5, true, geo.NewPoint(1, 0))

	w.Add(pstream.NewPoint("0", 10, 0))
	w.Add(pstream.NewPoint("1", -20, 0))
	w.Add(pstream.NewPoint("2", 4, 0))
	w.Add(pstream.NewPoint("3", 8, 0))
	w.Add(pstream.NewPoint("4", 1, 0))
	w.Add(pstream.NewPoint("5", 1, 0))
	w.Add(pstream.NewPoint("6", 1000, 0))
	w.Add(pstream.NewPoint("7", 25, 0))
	w.Add(pstream.NewPoint("8", 6, 0))

	got := w.GetAccum()

	expect := strings.Split("6,7,1,0,3", ",")

	for pos := range got {
		so.Equal(expect[pos], got[pos].ID)
	}
}
