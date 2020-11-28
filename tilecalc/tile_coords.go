package tilecalc

type TileCoord struct {
	X, Y, Z int
}

func GenerateAllTileCoords(maxZoom int) []*TileCoord {
	if maxZoom < 0 {
		return nil
	}
	if maxZoom > 16 {
		maxZoom = 16
	}

	var coords = make([]*TileCoord, 0)
	for i := 0; i <= maxZoom; i++ {

	}
	return coords
}

func AllTileCoordsNums(maxZoom int) int {
	if maxZoom < 0 {
		return 0
	}
	return int((uint64(1)<<uint64(2*(maxZoom+1)) - 1) / 3)
}

/*
 0: "https://a.tiles.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/{z}/{x}/{y}.vector.pbf?access_token=pk.eyJ1Ijoib3VjaGVuZyIsImEiOiJja2MzM3V6NXUwOGh5MnZwM2Zxcnk3aDVpIn0.SrJTQHglGBRg-kM6TAxOww"
 1: "https://b.tiles.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/{z}/{x}/{y}.vector.pbf?access_token=pk.eyJ1Ijoib3VjaGVuZyIsImEiOiJja2MzM3V6NXUwOGh5MnZwM2Zxcnk3aDVpIn0.SrJTQHglGBRg-kM6TAxOww"
*/
