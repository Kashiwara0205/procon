package p37

import (
	"testing"
)

func TestSlove (t *testing.T) {
	var maze = [][]int{
		{START, PATHWAY, PATHWAY, GOAL}, 
	}

	if 3 != Slove(maze){
		t.Errorf("最短経路コストに誤りがあります: %v", Slove(maze))
	}

	maze = [][]int{
		{START, PATHWAY, PATHWAY, PATHWAY, GOAL}, 
	}

	if 4 != Slove(maze){
		t.Errorf("最短経路コストに誤りがあります: %v", Slove(maze))
	}

	maze = [][]int{
		{START,   PATHWAY, PATHWAY, WALL, WALL}, 
		{PATHWAY, PATHWAY, PATHWAY, PATHWAY, GOAL}, 
	}

	if 5 != Slove(maze){
		t.Errorf("最短経路コストに誤りがあります: %v", Slove(maze))
	}

	maze = [][]int{
		{START,   WALL,    PATHWAY, PATHWAY, PATHWAY}, 
		{PATHWAY, WALL,    PATHWAY, PATHWAY, PATHWAY}, 
		{PATHWAY, PATHWAY, PATHWAY, PATHWAY, PATHWAY}, 
		{GOAL,    PATHWAY, PATHWAY, PATHWAY, PATHWAY}, 
	}

	if 3 != Slove(maze){
		t.Errorf("最短経路コストに誤りがあります: %v", Slove(maze))
	}
}