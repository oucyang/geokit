package tilecalc

import "testing"

func TestLonLat2Tile(t *testing.T) {
	lon, lat := 112.523205, 40.976342
	zoom := 4
	x, y := LonLat2Tile(lon, lat, zoom)
	t.Logf("x=%d y=%d", x, y)
}

func TestTile2LonLat(t *testing.T) {
	x, y := 13, 6
	zoom := 4
	lon, lat := Tile2LonLat(x, y, zoom)
	t.Logf("lon;lat=%.6f;%.6f", lon, lat)
}
