package components

import (
	"testing"
)

func TestPlayerName(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Tusshar"},
		{"Opponent"},
	}
	for _, test := range tests {
		var p Player
		p.SetName(test.name)
		gotName := p.GetName()
		if test.name != gotName {
			t.Error(test.name, gotName)
		}
	}
}

func TestPlayerMark(t *testing.T) {
	tests := []struct {
		mark string
	}{
		{XMark},
		{OMark},
	}
	for _, test := range tests {
		var p Player
		p.SetMark(test.mark)
		gotMark := p.GetMark()
		if test.mark != gotMark {
			t.Error(test.mark, gotMark)
		}
	}
}
