package compute

import (
	"fmt"
	"strconv"
	"strings"
)

func parseCoords(coords string) []*Point {
	var points = make([]*Point, 0)
	valueStrs := strings.Split(coords, ";")
	length := len(valueStrs) / 2
	for i := 0; i < length; i++ {
		lon, err := strconv.ParseFloat(valueStrs[i*2], 64)
		if err != nil {
			panic(fmt.Errorf("bad float '%s' -- %v", valueStrs[i*2], err))
		}
		lat, err := strconv.ParseFloat(valueStrs[i*2+1], 64)
		if err != nil {
			panic(fmt.Errorf("bad float '%s' - %v", valueStrs[i*2+1], err))
		}
		points = append(points, NewPoint(lon, lat))
	}
	return points

}

func newPolygonByCoords(coords string) Polygon {
	return Polygon(parseCoords(coords))
}

func newPointByCoords(coords string) *Point {
	return parseCoords(coords)[0]
}

func reversePoints(points []*Point) []*Point {
	var reverse = make([]*Point, len(points))
	for i := 0; i < len(points); i++ {
		reverse[i] = points[len(points)-i-1]
	}
	return reverse
}
