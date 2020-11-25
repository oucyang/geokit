package compute

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/oucyang/geokit/geom"
)

type PointInPolygonCase struct {
	geom.Polygon
	*geom.Point
	Answer bool
}

func (c *PointInPolygonCase) Test(t *testing.T) {
	if PointInPolygon2(c.Point, c.Polygon) != c.Answer {
		t.Fatalf("wrong answer [%s] [%+v] right answer=%t", c.Point, c.Polygon, c.Answer)
	}
}

func parseCoords(coords string) []*geom.Point {
	var points = make([]*geom.Point, 0)
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
		points = append(points, geom.NewPoint(lon, lat))
	}
	return points

}

func newPolygonByCoords(coords string) geom.Polygon {
	return geom.Polygon(parseCoords(coords))
}

func newPointByCoords(coords string) *geom.Point {
	return parseCoords(coords)[0]
}

func TestPointInPolygon(t *testing.T) {
	var case1 = &PointInPolygonCase{Polygon: newPolygonByCoords("0;0;10;0;10;10;0;10"), Point: newPointByCoords("20;20"), Answer: false}
	var case2 = &PointInPolygonCase{Polygon: newPolygonByCoords("0;0;10;0;10;10;0;10"), Point: newPointByCoords("5;5"), Answer: true}
	var case3 = &PointInPolygonCase{Polygon: newPolygonByCoords("0;0;10;0;10;10;0;10"), Point: newPointByCoords("-1;10"), Answer: false}
	var case4 = &PointInPolygonCase{Polygon: newPolygonByCoords("0;0;5;5;5;0"), Point: newPointByCoords("3;3"), Answer: true}
	var case5 = &PointInPolygonCase{Polygon: newPolygonByCoords("0;0;5;5;5;0"), Point: newPointByCoords("5;1"), Answer: true}
	var case6 = &PointInPolygonCase{Polygon: newPolygonByCoords("0;0;5;5;5;0"), Point: newPointByCoords("8;1"), Answer: false}
	var case7 = &PointInPolygonCase{Polygon: newPolygonByCoords("0;0;5;5;5;0"), Point: newPointByCoords("5;5"), Answer: true}
	cases := []*PointInPolygonCase{case1, case2, case3, case4, case5, case6, case7}

	cases[0].Test(t)
	cases[1].Test(t)
	cases[2].Test(t)
	cases[3].Test(t)
	cases[4].Test(t)
	cases[5].Test(t)
	cases[6].Test(t)

	var complex = newPolygonByCoords("141.569972;29.658172;148.513331;28.968448;151.062159;23.048090;161.696925;21.584604;165.652003;26.006137;159.411769;32.961385;164.333644;36.922403;175.319972;33.914825;174.001613;12.349335;144.030909;11.575504")
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("151.062159;23.048090"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("161.696925;21.584604"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("129.880519;21.584604"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("146.403956;21.584604"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("164.333644;36.922403"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("146.403956;36.922403"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("158.181300;26.242872"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("177.341456;26.557769"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("168.552394;17.275851"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("161.609034;7.066764"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("159.499659;39.883351"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("165.652003;26.006137"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("165.652003;33.463046"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("165.652003;39.571789"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("165.652003;8.061966"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("161.696925;27.112752"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("161.696925;7.816239"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("161.696925;37.894588"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("161.696925;32.561515"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("161.696925;15.982095"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("141.569972;29.658172"), Answer: true}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("141.569972;33.394026"), Answer: false}).Test(t)
	(&PointInPolygonCase{Polygon: complex, Point: newPointByCoords("141.569972;10.022043"), Answer: false}).Test(t)
}
