package service

import (
	"Tic-Tac-Toe/components"
	_ "fmt"
	"testing"
)

func TestCheckRows(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
			Size: 3,
		},
		},
		}, components.OMark, true},
	}

	for _, test := range tests {
		got := test.rs.CheckRow(test.mark)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}

func TestCheckColumns(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
			Size: 3,
		},
		},
		}, components.OMark, true},
	}

	for _, test := range tests {
		got := test.rs.CheckColumn(test.mark)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}

func TestCheckFirstDiagonal(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
			},
			Size: 3,
		},
		},
		}, components.XMark, true},
	}

	for _, test := range tests {
		got := test.rs.CheckFirstDiagonal(test.mark)
		test.rs.CheckFirstDiagonal(test.mark)
		// fmt.Println(test.want,got)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}

func TestCheckSecondDiagonal(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
			},
			Size: 3,
		},
		},
		}, components.XMark, true},
	}

	for _, test := range tests {
		got := test.rs.CheckSecondDiagonal(test.mark)
		test.rs.CheckSecondDiagonal(test.mark)
		// fmt.Println(test.want,got)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}
