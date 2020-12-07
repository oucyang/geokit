package tcalc

import (
	"fmt"
	"github.com/oucyang/geokit/mvt"
	"io/ioutil"
	"math/rand"

	"github.com/golang/protobuf/proto"
)

type TileCoord struct {
	X, Y, Z int
}

func (t *TileCoord) TopLeftChild() *TileCoord {
	return &TileCoord{Z: t.Z + 1, X: t.X << 1, Y: t.Y << 1}
}

func (t *TileCoord) TopRightChild() *TileCoord {
	return &TileCoord{Z: t.Z + 1, X: (t.X << 1) + 1, Y: t.Y << 1}
}

func (t *TileCoord) BottomLeftChild() *TileCoord {
	return &TileCoord{Z: t.Z + 1, X: t.X << 1, Y: (t.Y << 1) + 1}
}

func (t *TileCoord) BottomRightChild() *TileCoord {
	return &TileCoord{Z: t.Z + 1, X: (t.X << 1) + 1, Y: (t.Y << 1) + 1}
}

func (t *TileCoord) Children() []*TileCoord {
	z, x, y := t.Z+1, t.X<<1, t.Y<<1
	return []*TileCoord{
		&TileCoord{Z: z, X: x, Y: y},         // top left
		&TileCoord{Z: z, X: x + 1, Y: y},     // top right
		&TileCoord{Z: z, X: x, Y: y + 1},     // bottom left
		&TileCoord{Z: z, X: x + 1, Y: y + 1}, // bottom right
	}
}

func (t *TileCoord) String() string {
	return fmt.Sprintf("%d_%d_%d", t.Z, t.X, t.Y)
}

func (t *TileCoord) MapboxUrl(accessToken string) string {
	return fmt.Sprintf(mapboxUrl, t.Z, t.X, t.Y, accessToken)
}

/*
 0: "https://a.tiles.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/{z}/{x}/{y}.vector.pbf?access_token=?"
 1: "https://b.tiles.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/{z}/{x}/{y}.vector.pbf?access_token=?"
    "https://api.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/{z}/{x}/{y}.vector.pbf?sku=101Z5P5w9GWb7&access_token=?"
*/
const (
	mapboxUrl0 = "https://a.tiles.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/%d/%d/%d.vector.pbf?access_token=%s"
	mapboxUrl1 = "https://a.tiles.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/%d/%d/%d.vector.pbf?access_token=%s"
	mapboxUrl  = "https://api.mapbox.com/v4/mapbox.mapbox-streets-v8,mapbox.mapbox-terrain-v2/%d/%d/%d.vector.pbf?&access_token=%s"
)

func (t *TileCoord) MapboxTileUrl(accessToke string) string {
	if rand.Intn(2) == 0 {
		return fmt.Sprintf(mapboxUrl0, t.Z, t.X, t.Y, accessToke)
	} else {
		return fmt.Sprintf(mapboxUrl1, t.Z, t.X, t.Y, accessToke)
	}
}

func AllTileCoordsNums(maxZoom int) int {
	if maxZoom < 0 {
		return 0
	}
	return int((uint64(1)<<uint64(2*(maxZoom+1)) - 1) / 3)
}

func DecodeTile(filename string) error {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	vect := &mvt.Tile{}
	if err := proto.Unmarshal(body, vect); err != nil {
		return err
	}
	for _, layer := range vect.Layers {
		fmt.Printf("name=%s\n", *layer.Name)
	}
	return nil
}
