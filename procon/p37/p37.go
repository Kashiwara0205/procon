package p37

const (
	START = iota
	GOAL
	WALL
	PATHWAY
)

const INF = 999999

type Point struct {
	x int
	y int
}

func CreatePoint(x int, y int) Point {
	return Point{x: x, y: y}
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

func Slove(maze [][]int) int {
	col := len(maze)
	row := len(maze[0])

	distanceTable := createDistanceTable(col, row)

	sx, sy := selectPoint(col, row, maze, START)
	startPoint := CreatePoint(sx, sy)

	gx, gy := selectPoint(col, row, maze, GOAL)
	goalPoint := CreatePoint(gx, gy)

	return bfs(col, row, distanceTable, maze, startPoint, goalPoint)
}

func createDistanceTable(col int, row int) [][]int {
	distanceTable := make([][]int, col)

	for i := 0; i < col; i++ {
		distanceTable[i] = make([]int, row)
		for j := 0; j < row; j++ {
			distanceTable[i][j] = INF
		}
	}

	return distanceTable
}

func selectPoint(col int, row int, maze [][]int, pointType int) (int, int) {
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {

			if pointType == maze[i][j] {
				return j, i
			}
		}
	}

	return 0, 0
}

func bfs(col int, row int, distanceTable [][]int, maze [][]int, startPoint Point, goalPoint Point) int {

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	var queue []Point
	queue = append(queue, startPoint)

	distanceTable[startPoint.Y()][startPoint.X()] = 0

	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]

		if point.X() == goalPoint.X() && point.Y() == goalPoint.Y() {
			break
		}

		for i := 0; i < 4; i++ {
			nx := point.X() + dx[i]
			ny := point.Y() + dy[i]

			if 0 <= nx && nx < row && 0 <= ny && ny < col && maze[ny][nx] != WALL && distanceTable[ny][nx] == INF {
				queue = append(queue, CreatePoint(nx, ny))
				distanceTable[ny][nx] = distanceTable[point.Y()][point.X()] + 1
			}

		}
	}

	return distanceTable[goalPoint.Y()][goalPoint.X()]
}
