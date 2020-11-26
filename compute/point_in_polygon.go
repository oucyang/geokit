package compute

import (
	"math"
)

/*
 其它连接
https://en.wikipedia.org/wiki/Point_in_polygon
https://observablehq.com/@tmcw/understanding-point-in-polygon
https://wrf.ecse.rpi.edu//Research/Short_Notes/pnpoly.html
http://graphics.stanford.edu/courses/cs368-00-spring/TA/manuals/LEDA/geo_alg.html
https://geomalgorithms.com/
https://automating-gis-processes.github.io/CSC18/course-info/Installing_Anacondas_GIS.html
https://github.com/sromku/polygon-contains-point/blob/master/src/main/java/com/snatik/polygon/Polygon.java
https://codeforces.com/blog/entry/48868
https://math.stackexchange.com/questions/2273108/polygon-in-polygon-testing
https://en.wikipedia.org/wiki/Bentley%E2%80%93Ottmann_algorithm
https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection
*/

// http://www.eecs.umich.edu/courses/eecs380/HANDOUTS/PROJ2/InsidePoly.html
func PointInPolygon(p *Point, polygon Polygon) bool {
	if len(polygon) < 3 {
		return false
	}
	var (
		inside  = false
		xinters float64
		p1, p2  *Point
	)
	p1 = polygon[0]
	for i := 1; i <= len(polygon); i++ {
		if p.X == p1.X && p.Y == p1.Y {
			return true
		}
		p2 = polygon[i%len(polygon)]
		if p.Y > math.Min(p1.Y, p2.Y) && p.Y <= math.Max(p1.Y, p2.Y) && p.X <= math.Max(p1.X, p2.X) {
			if p1.Y != p2.Y { // 可以优化掉
				xinters = (p.Y-p1.Y)*(p2.X-p1.X)/(p2.Y-p1.Y) + p1.X
				if p.X == xinters {
					return true
				}
				if p1.X == p2.X || p.X < xinters {
					inside = !inside
				}
			}
		}
		p1 = p2
	}
	return inside
}

func PointInPolygon2(p *Point, polygon Polygon) bool {
	if len(polygon) < 3 {
		return false
	}
	var (
		inside  = false
		xinters float64
		p1, p2  *Point
	)
	p1 = polygon[0]
	for i := 1; i <= len(polygon); i++ {
		if p.X == p1.X && p.Y == p1.Y {
			return true
		}
		p2 = polygon[i%len(polygon)]
		if p.Y > math.Min(p1.Y, p2.Y) && p.Y <= math.Max(p1.Y, p2.Y) && p.X <= math.Max(p1.X, p2.X) {
			xinters = (p.Y-p1.Y)*(p2.X-p1.X)/(p2.Y-p1.Y) + p1.X
			if p.X == xinters {
				return true
			}
			if p1.X == p2.X || p.X < xinters {
				inside = !inside
			}
		}
		p1 = p2
	}
	return inside
}
