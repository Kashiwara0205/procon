package p35

const (
	NONE = iota
	WET
)

func Slove(field [][]int) int {
	colSize := len(field)
	rowSize := len(field[0])

	var rakeCount = 0

	for i := 0; i < colSize; i++ {
		for j := 0; j < rowSize; j++ {
			if WET == field[i][j] {
				updateField(field, i, j, colSize, rowSize)
				rakeCount += 1
			}
		}
	}

	return rakeCount
}

func updateField(field [][]int, i int, j int, colSize int, rowSize int) [][]int {
	field[i][j] = NONE

	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			row := j + dx
			col := i + dy

			if row < rowSize && col < colSize && -1 < row && -1 < col && WET == field[col][row] {
				updateField(field, col, row, colSize, rowSize)
			}

		}
	}

	return field
}
