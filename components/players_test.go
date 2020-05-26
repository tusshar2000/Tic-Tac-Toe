package components

import (
	"testing"
)

func TestPlayer(t *testing.T) {
	var p Player
	want := "Tusshar"
	p.SetName("Tusshar")
	if want != p.GetName() {
		t.Error(want, p.GetName())
	}
}
