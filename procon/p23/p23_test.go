package p23

import "testing"

func TestRunAnts (t *testing.T) {
	var l = 10
	var a = []int{1}

	var result = RunAnts(a, l, MIN)
	var directions = result.Directions()
	if 1 != result.Cost(){
		t.Errorf("コストに誤りがあります: %v", result.Cost())
	}

	if LEFT != directions[0]{
		t.Errorf("向きが右になっています")
	}

	result = RunAnts(a, l, MAX)
	directions = result.Directions()
	if 9 != result.Cost(){
		t.Errorf("コストに誤りがあります: %v", result.Cost())
	}

	if RIGHT != directions[0]{
		t.Errorf("向きが左になっています")
	}

	a = []int{1, 2}
	result = RunAnts(a, l, MIN)
	directions = result.Directions()
	if 2 != result.Cost(){
		t.Errorf("コストに誤りがあります: %v", result.Cost())
	}

	if LEFT != directions[0] && LEFT != directions[1]{
		t.Errorf("ありの向きに誤りがあります")
	}

	result = RunAnts(a, l, MAX)
	directions = result.Directions()
	if 9 != result.Cost(){
		t.Errorf("コストに誤りがあります: %v", result.Cost())
	}

	if RIGHT != directions[0] && RIGHT != directions[1]{
		t.Errorf("ありの向きに誤りがあります")
	}

	a = []int{2, 6, 7}
	result = RunAnts(a, l, MIN)
	directions = result.Directions()
	if 4 != result.Cost(){
		t.Errorf("コストに誤りがあります: %v", result.Cost())
	}

	if LEFT != directions[0] && RIGHT != directions[1] && RIGHT != directions[2]{
		t.Errorf("ありの向きに誤りがあります")
	}

	result = RunAnts(a, l, MAX)
	directions = result.Directions()
	if 8 != result.Cost(){
		t.Errorf("コストに誤りがあります: %v", result.Cost())
	}

	if RIGHT != directions[0] && RIGHT != directions[1] && RIGHT != directions[2]{
		t.Errorf("ありの向きに誤りがあります")
	}
}