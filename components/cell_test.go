package components

import (
	"testing"
)

func TestNewCell(t *testing.T) {
	test := NewCell()
	expected := NoMark
	if test.mark != expected {
		t.Error(test.mark, expected)
	}
}

func TestSetMark(t *testing.T) {
	tests := []struct {
		input1   *Cell
		input2   string
		expected bool
	}{
		{&Cell{mark: NoMark}, XMark, true},
		{&Cell{mark: NoMark}, OMark, true},
		{&Cell{mark: XMark}, OMark, false},
		{&Cell{mark: XMark}, XMark, false},
		{&Cell{mark: OMark}, OMark, false},
		{&Cell{mark: OMark}, XMark, false},
	}
	for _, test := range tests {
		actual := Mark(test.input1, test.input2)
		//fmt.Println(actual)
		//fmt.Println(test.expected)
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}
}

func TestGetMark(t *testing.T) {
	tests := []struct {
		input1   *Cell
		expected string
	}{
		{&Cell{mark: NoMark}, NoMark},
		{&Cell{mark: OMark}, OMark},
		{&Cell{mark: XMark}, XMark},
	}
	for _, test := range tests {
		actual := GetMark(test.input1)
		if actual != test.expected {
			t.Error(actual, test.expected)
		}
	}

}
