package solution

import (
	"math"
	"strconv"
	"strings"
)

type Day5 struct {
}

func (d Day5) SolvePart1(input string) int {
	return solveDay5ConsideringDiagonals(input, false)
}

func (d Day5) SolvePart2(input string) int {
	return solveDay5ConsideringDiagonals(input, true)
}

func solveDay5ConsideringDiagonals(input string, withDiagonals bool) int {
	rows := strings.Split(input, "\n")
	lines := make([]Line, 0)
	width := 0
	for _, row := range rows {
		line := CreateLine(row)
		if withDiagonals || !line.isDiagonal() {
			lines = append(lines, line)
			width = int(math.Max(float64(width), math.Max(float64(line.x1), float64(line.x2))))
		}
	}
	coveredCells := make(map[int]int)
	for _, line := range lines {
		for _, vertex := range line.getVertices() {
			cellIndex := vertex.y*width + vertex.x
			coverage := coveredCells[cellIndex]
			coveredCells[cellIndex] = coverage + 1
		}
	}
	cellsWithCoverage2Plus := 0
	for _, coverage := range coveredCells {
		if coverage >= 2 {
			cellsWithCoverage2Plus = cellsWithCoverage2Plus + 1
		}
	}
	return cellsWithCoverage2Plus
}

type Line struct {
	x1, y1, x2, y2 int
}

type Vertex struct {
	x, y int
}

func CreateLine(row string) Line {
	vertices := strings.Split(row, " -> ")
	a := strings.Split(vertices[0], ",")
	b := strings.Split(vertices[1], ",")
	x1, _ := strconv.Atoi(a[0])
	y1, _ := strconv.Atoi(a[1])
	x2, _ := strconv.Atoi(b[0])
	y2, _ := strconv.Atoi(b[1])
	return Line{x1, y1, x2, y2}
}

func (line Line) isDiagonal() bool {
	return line.x1 != line.x2 && line.y1 != line.y2
}

func (line Line) getVertices() []Vertex {
	vertices := make([]Vertex, 0)
	x := line.x1
	y := line.y1
	vertices = append(vertices, Vertex{x, y})
	for {
		if x == line.x2 && y == line.y2 {
			break
		}
		if line.x2 > x {
			x = x + 1
		} else if line.x2 < x {
			x = x - 1
		}
		if line.y2 > y {
			y = y + 1
		} else if line.y2 < y {
			y = y - 1
		}
		vertices = append(vertices, Vertex{x, y})
	}
	return vertices
}
