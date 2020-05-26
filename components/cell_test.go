package components

import "testing"

func TestNewCell(t *testing.T) {
	want := NoMark
	got := NewCell()
	if want != got.mark {
		t.Error(want, got.mark)
	}
}

func TestMark(t *testing.T) {
	tests := []struct {
		cell *Cell
		mark string
		want bool
	}{
		{&Cell{mark: NoMark}, XMark, true},
		{&Cell{mark: NoMark}, OMark, true},
		{&Cell{mark: XMark}, OMark, false},
		{&Cell{mark: XMark}, XMark, false},
		{&Cell{mark: OMark}, OMark, false},
		{&Cell{mark: OMark}, XMark, false},
	}
	for _, test := range tests {
		got, _ := test.cell.SetMark(test.mark)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}

func TestGetMark(t *testing.T) {
	tests := []struct {
		cell *Cell
		want string
	}{
		{&Cell{mark: NoMark}, NoMark},
		{&Cell{mark: OMark}, OMark},
		{&Cell{mark: XMark}, XMark},
	}
	for _, test := range tests {
		got := test.cell.GetMark()
		if test.want != got {
			t.Error(test.want, got)
		}
	}

}
