package p34

import "testing"

func TestCheckPartialSum(t *testing.T) {
	var a = []int{1}
	if !Slove(a, 1) {
		t.Errorf("与えた数: [1], 合計値: 1 で失敗")
	}

	a = []int{1, 2}
	if !Slove(a, 3) {
		t.Errorf("与えた数: [1, 2], 合計値: 3 で失敗")
	}

	a = []int{1, 2}
	if !Slove(a, 1) {
		t.Errorf("与えた数: [1, 2], 合計値: 1 で失敗")
	}
}
