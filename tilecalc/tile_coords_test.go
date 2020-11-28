package tilecalc

import "testing"

func TestAllTileCoordsNums(t *testing.T) {
	for i := 0; i <= 16; i++ {
		t.Logf("maxzoom=%d total=%d", i, AllTileCoordsNums(i))
	}
}
