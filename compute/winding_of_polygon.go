package compute

/*
other urls:
https://en.wikipedia.org/wiki/Green's_theorem
https://en.wikipedia.org/wiki/Curve_orientation
*/

// https://www.element84.com/blog/determining-the-winding-of-a-polygon-given-as-a-set-of-ordered-points
// return true clockwise, false counter-clockwise(anti-clockwise)
func WindingOfPolygon(poly Polygon) bool {
	var area = 0.0
	for last := 0; last < len(poly); last++ {
		next := (last + 1) % len(poly)
		area += (poly[next].X - poly[last].X) * (poly[next].Y + poly[last].Y)
	}
	return area > 0.0
}
