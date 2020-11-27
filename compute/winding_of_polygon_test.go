package compute

import "testing"

type WindingOfPolygonCase struct {
	Polygon
	Answer bool
}

func (c *WindingOfPolygonCase) Test(t *testing.T) {
	if WindingOfPolygon(c.Polygon) != c.Answer {
		t.Fatalf("wrong answer polygon=[%s] right answer=%t", c.Polygon, c.Answer)
	}
}

func TestWindingOfPolygon(t *testing.T) {
	var case1 = &WindingOfPolygonCase{Polygon: newPolygonByCoords("-148.552568;0;-143.502121;6.808066;-137.580907;0;-143.502121;-6.808066"), Answer: true}
	var case2 = &WindingOfPolygonCase{Polygon: newPolygonByCoords("-143.502121;-6.808066;-137.580907;0;-143.502121;6.808066;-148.552568;0"), Answer: false}
	var case3 = &WindingOfPolygonCase{Polygon: newPolygonByCoords("-153.658943;7.074890;-146.614745;8.897338;-138.206129;9.961586;-137.159019;2.586310;-140.998424;4.043437;-141.125347;6.759894;-144.330140;6.255477;-144.234948;2.554612;-148.835888;1.761915"), Answer: true}
	var case4 = &WindingOfPolygonCase{Polygon: NewPolygon(reversePoints(parseCoords("-153.658943;7.074890;-146.614745;8.897338;-138.206129;9.961586;-137.159019;2.586310;-140.998424;4.043437;-141.125347;6.759894;-144.330140;6.255477;-144.234948;2.554612;-148.835888;1.761915"))), Answer: false}
	case1.Test(t)
	case2.Test(t)
	case3.Test(t)
	case4.Test(t)
}
