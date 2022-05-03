package p35

import (
	"testing"
)

func TestSlove(t *testing.T) {
	var field = [][]int{
		{0, 0, 0},
		{0, 1, 0},
	}

	if 1 != Slove(field) {
		t.Errorf("水たまりのカウントに誤りがあります")
	}

	field = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	if 2 != Slove(field) {
		t.Errorf("水たまりのカウントに誤りがあります")
	}

	field = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 0},
	}

	if 1 != Slove(field) {
		t.Errorf("水たまりのカウントに誤りがあります")
	}

	field = [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
		{1, 1, 0},
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 0},
		{1, 1, 0},
	}

	if 3 != Slove(field) {
		t.Errorf("水たまりのカウントに誤りがあります")
	}
}
