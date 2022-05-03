package p23

type Result struct {
	cost       int
	directions []int
}

func NewResult(cost int, directions []int) Result {
	return Result{cost: cost, directions: directions}
}

func (a *Result) Cost() int {
	return a.cost
}

func (a *Result) Directions() []int {
	return a.directions
}
