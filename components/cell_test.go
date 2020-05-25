package components

import "testing"

func TestNewCell(t *testing.T) {
	want := NewCell()
	got := NoMark
	if want.mark != got {
		t.Error(want.mark, got)
	}
}

func TestMark(t *testing.T) {
	tests := []struct {
		input1 *Cell
		input2 string
		got    bool
	}{
		{&Cell{mark: NoMark}, XMark, true},
		{&Cell{mark: NoMark}, OMark, true},
		{&Cell{mark: XMark}, OMark, false},
		{&Cell{mark: XMark}, XMark, false},
		{&Cell{mark: OMark}, OMark, false},
		{&Cell{mark: OMark}, XMark, false},
	}
	for _, test := range tests {
		want := Mark(test.input1, test.input2)
		if want != test.got {
			t.Error(want, test.got)
		}
	}
}

func TestGetMark(t *testing.T) {
	tests := []struct {
		input *Cell
		got   string
	}{
		{&Cell{mark: NoMark}, NoMark},
		{&Cell{mark: OMark}, OMark},
		{&Cell{mark: XMark}, XMark},
	}
	for _, test := range tests {
		want := GetMark(test.input)
		if want != test.got {
			t.Error(want, test.got)
		}
	}

}
