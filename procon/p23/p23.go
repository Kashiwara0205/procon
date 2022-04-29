package p23

import (
	"math"
)

func RunAnts(antPositions []int, stickSize int, calcType int) Result{
	antSize := len(antPositions)
	ants := createAnts(antPositions, stickSize, antSize, calcType)

	for i := 0; i < antSize; i++ {
		for j := 0; j < stickSize; j++{
			if isDrop( ants[i].Position(), stickSize ){
				continue
			}

			if RIGHT == ants[i].Direction(){
				ants[i].MoveRight()
			}else{
				ants[i].MoveLeft()
			}
		}
	}

	cost := calcCost(ants)
	directions := createDirections(ants)

	return NewResult(cost, directions)
}

func calcCost(ants []Ant) int{
	var cost = 0

	for _, ant := range ants {
		if cost < ant.Cost() {
			cost = ant.Cost()
		}
	}

	return cost
}

func createDirections(ants []Ant) []int{
	var directions []int
	for _, ant := range ants {
		directions = append( directions, ant.Direction() )
	}

	return directions
}

func isDrop(antPosition int, stickSize int) bool{
	// 右に落下
	if antPosition >= stickSize{
		return true 
	}

	// 左に落下
	if antPosition <= 0 {
		return true
	}

	return false
}

func createAnts(antPositions []int, stickSize int, antSize int, calcType int)[]Ant{
	var ants []Ant
	for i := 0; i < antSize; i++ {
		direction := LEFT
		if antPositions[i] > int( math.Abs(float64(stickSize) / 2) ){
			direction = RIGHT
		}

		if MAX == calcType {
			direction ^= 1
		}

		ants = append(ants, NewAnt(antPositions[i], direction))
	}

	return ants
}