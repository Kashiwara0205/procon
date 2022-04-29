package p21

import (
	"math"
)

func CalcMaxPerimeter(a []int) int{
	var maxPerimeter = 0

	if !validateSize(a){
		return maxPerimeter
	}
	
	size := len(a)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if !validateMathematicalConditions(a[i], a[j], a[k]){
					continue
				}

				maxPerimeter = calcPerimeter( a[i], a[j], a[k] )
			}
		}
	}

	return maxPerimeter
}

func validateSize(a []int) bool{
	size := len(a)
	if size < 3{
		return false
	}

	return true
}

func validateMathematicalConditions(a, b, c int) bool{
	return int( math.Abs(float64(a) - float64(b)) ) < c && c < (a + b)
}

func calcPerimeter(a, b, c int) int{
	return a + b + c
}