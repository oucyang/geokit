package tcalc

import (
	"testing"
)

func TestAllTileCoordsNums(t *testing.T) {
	for i := 0; i <= 16; i++ {
		t.Logf("maxzoom=%d total=%d", i, AllTileCoordsNums(i))
	}
}

func TestTileCoord_Children(t *testing.T) {
	var tile = &TileCoord{Z: 10, X: 833, Y: 424}
	children := tile.Children()
	t.Logf("%+v", children)
}
