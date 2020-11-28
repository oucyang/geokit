package tilecalc

import "math"

func Tile2LonLat(x, y, zoom int) (float64, float64) {
	n := float64(int64(0x1) << zoom)
	lon := 360.0*float64(x)/float64(n) - 180.0
	lat := math.Atan(math.Sinh(math.Pi*(1-2.0*float64(y)/n))) * 180.0 / math.Pi
	return lon, lat
}

func LonLat2Tile(lon, lat float64, zoom int) (int, int) {
	lat_rad := lat * math.Pi / 180
	n := float64(int64(0x1) << zoom)
	llx := n * ((lon + 180) / 360)
	lly := n * (1 - (math.Log(math.Tan(lat_rad)+1/math.Cos(lat_rad)) / math.Pi)) / 2
	return int(llx), int(lly)
}
