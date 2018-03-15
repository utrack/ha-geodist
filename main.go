package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/utrack/ha-geodist/geo"
	"github.com/utrack/ha-geodist/pstream"
)

var (
	filePath = flag.String("csv", "", "Path to the points' CSV file. Will read from stdin if empty")
)

func main() {
	// TODO we should insert a help text there

	flag.Parse()

	var r io.Reader

	if *filePath == "" {
		r = os.Stdin
	} else {
		var err error
		r, err = os.Open(*filePath)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	csv := pstream.NewCSVStream(r, true)

	haOffice := geo.NewPoint(51.925146, 4.478617)

	fmt.Printf("Using point at %v, %v as a target.\n", haOffice.Lat(), haOffice.Lon())

	closestPoints := newWindow(5, false, haOffice)
	farthestPoints := newWindow(5, true, haOffice)

	var err error
	for {
		var point pstream.Point
		point, err = csv.Next()
		if err != nil {
			break
		}
		closestPoints.Add(point)
		farthestPoints.Add(point)
	}

	// Ignore ErrNoPoints, it is expected
	if err == pstream.ErrNoPoints {
		err = nil
	}
	if err != nil {
		logrus.Fatal("error when reading coords: ", err)
	}

	fmt.Println("Five closest points are: ")
	for _, point := range closestPoints.GetAccum() {
		fmt.Printf("%v,%v,%v\n", point.ID, point.Lat(), point.Lon())
	}
	fmt.Println("Five farthest points are: ")
	for _, point := range farthestPoints.GetAccum() {
		fmt.Printf("%v,%v,%v\n", point.ID, point.Lat(), point.Lon())
	}
}
