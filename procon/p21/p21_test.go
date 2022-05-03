package p21

import "testing"

func TestCalcMaxPerimeter(t *testing.T) {
	var a []int
	if 0 != CalcMaxPerimeter(a) {
		t.Errorf("配列の数が0なのに0以外の数字が返却されている")
	}

	a = []int{0}
	if 0 != CalcMaxPerimeter(a) {
		t.Errorf("配列の数が1なのに0以外の数字が返却されている")
	}

	a = []int{0, 1}
	if 0 != CalcMaxPerimeter(a) {
		t.Errorf("配列の数が2なのに0以外の数字が返却されている")
	}

	a = []int{2, 3, 4, 5, 10}
	if 12 != CalcMaxPerimeter(a) {
		t.Errorf("最大周長が12ではない: %v", CalcMaxPerimeter(a))
	}

	a = []int{2, -3, -4, -5, -10}
	if 0 != CalcMaxPerimeter(a) {
		t.Errorf("最大周長が0ではない: %v", CalcMaxPerimeter(a))
	}
}
