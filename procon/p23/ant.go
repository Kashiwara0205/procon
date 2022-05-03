package p23

type Ant struct {
	cost      int
	position  int
	direction int
}

func NewAnt(position int, direction int) Ant {
	return Ant{position: position, direction: direction}
}

func (a *Ant) Cost() int {
	return a.cost
}

func (a *Ant) Position() int {
	return a.position
}

func (a *Ant) Direction() int {
	return a.direction
}

func (a *Ant) SetDirection(direction int) {
	a.direction = direction
}

func (a *Ant) MoveLeft() {
	a.cost += 1
	a.position -= 1
}

func (a *Ant) MoveRight() {
	a.cost += 1
	a.position += 1
}
