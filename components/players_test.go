package components

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		name string
		mark string
	}{
		{"Tusshar", XMark},
		{"Opponent", OMark},
	}
	for _, test := range tests {
		wantName := test.name
		wantMark := test.mark
		testPlayer := NewPlayer(wantName, wantMark)
		gotName := testPlayer.name
		gotMark := testPlayer.mark
		if wantName != gotName {
			t.Error(wantName, gotName)
		}
		if wantMark != gotMark {
			t.Error(wantMark, gotMark)
		}
	}
}
