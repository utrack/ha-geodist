package pstream

import (
	"encoding/csv"
	"io"
	"sync"

	"strconv"

	"github.com/pkg/errors"
	"github.com/utrack/ha-geodist/geo"
)

// csvIterator iterates over the comma-separated list of ID,Lat,Lon values;
// accepting one value per line.
type csvIterator struct {
	r         *csv.Reader
	mu        sync.Mutex
	lastError error
}

var _ Points = &csvIterator{}

// NewCSVStream creates a stream of Points using a Reader that
// reads comma-separated values.
func NewCSVStream(r io.Reader, skipHeader bool) Points {
	cr := csv.NewReader(r)
	cr.ReuseRecord = true
	cr.FieldsPerRecord = 3

	// TODO detect ID/Lat/Lon column positions
	// if needed
	if skipHeader {
		_, _ = cr.Read()
	}
	return &csvIterator{
		r: cr,
	}
}

// Next implements Points.
func (i *csvIterator) Next() (Point, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.lastError != nil {
		return Point{}, errors.Wrap(i.lastError, "reader had an error previously")
	}

	row, err := i.r.Read()
	if err != nil {
		if err == io.EOF {
			return Point{}, ErrNoPoints
		}
		err = errors.Wrap(err, "error when reading CSV-row from the reader")
		i.lastError = err
		return Point{}, err
	}

	lat, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		err = errors.Wrapf(err, "couldn't parse value %v to lat", row[1])
		i.lastError = err
		return Point{}, err
	}

	lon, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		err = errors.Wrapf(err, "couldn't parse value %v to lon", row[2])
		i.lastError = err
		return Point{}, err
	}
	p := geo.NewPoint(lat, lon)

	return Point{
		ID:    row[0],
		Point: p,
	}, nil
}
