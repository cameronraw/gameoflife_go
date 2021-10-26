package cell

import (
	"testing"
)

func TestCellInitializedWithState(t *testing.T) {
	var cell cell = NewCell("ALIVE")
	got := cell.state
	want := "ALIVE"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
