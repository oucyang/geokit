package geom

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("%f;%f", p.X, p.Y)
}

type Line []*Point

type Polygon Line

func (p Polygon) String() string {
	var ss = make([]string, len(p))
	for i, pt := range p {
		ss[i] = pt.String()
	}
	return strings.Join(ss, ";")
}

func NewPoint(lng, lat float64) *Point {
	return &Point{lng, lat}
}

func NewLine(points []*Point) Line {
	return Line(points)
}

func NewPolygon(points []*Point) Polygon {
	return Polygon(points)
}
