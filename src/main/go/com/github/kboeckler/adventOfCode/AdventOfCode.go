package main

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, solution := range allSolutions() {
		nameOfSolution := reflect.TypeOf(solution).Name()
		data, err := ioutil.ReadFile(workingDir + "/src/main/resources/com/github/kboeckler/adventOfCode/" + nameOfSolution + ".txt")
		if err != nil {
			log.Fatal("File reading error", err)
		}
		input := strings.ReplaceAll(string(data), "\r\n", "\n")
		result1 := solution.solvePart1(input)
		result2 := solution.solvePart2(input)
		log.Printf("Solution %v - Part1: %d Part2: %d\n", nameOfSolution, result1, result2)
	}
}

func allSolutions() []Solution {
	var solutions = make([]Solution, 0)
	solutions = append(solutions, Day3{})
	solutions = append(solutions, Day4{})
	solutions = append(solutions, Day5{})
	return solutions
}

type Solution interface {
	solvePart1(input string) int
	solvePart2(input string) int
}

type Day3 struct {
}

func (d Day3) solvePart1(input string) int {
	rows := strings.Split(input, "\n")
	width := len(rows[0])
	height := len(rows)

	countOfOnesByPosition := make([]int, 0, width)
	for i := 0; i < width; i++ {
		countOfOnesByPosition = append(countOfOnesByPosition, 0)
	}

	for _, row := range rows {
		for i := range row {
			if row[i] == '1' {
				countOfOnesByPosition[i] = countOfOnesByPosition[i] + 1
			}
		}
	}

	gamma := 0
	epsilon := 0
	for i := width - 1; i >= 0; i-- {
		exponentOf2 := width - 1 - i
		if countOfOnesByPosition[i] >= height/2 {
			gamma = gamma + int(math.Pow(2, float64(exponentOf2)))
		} else {
			epsilon = epsilon + int(math.Pow(2, float64(exponentOf2)))
		}
	}
	return gamma * epsilon
}

func (d Day3) solvePart2(input string) int {
	rows := strings.Split(input, "\n")

	oxygen := determineRating(rows, 1)
	co2 := determineRating(rows, -1)
	oxygenDecimal, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Decimal, _ := strconv.ParseInt(co2, 2, 64)

	return int(oxygenDecimal * co2Decimal)
}

func determineRating(rows []string, comparatorFactor int) string {
	width := len(rows[0])
	workingRows := rows
	for bitPosition := 0; bitPosition < width; bitPosition = bitPosition + 1 {
		if len(workingRows) <= 1 {
			break
		}
		rowsWithOnes := make([]string, 0)
		rowsWithZeroes := make([]string, 0)
		for _, row := range workingRows {
			if row[bitPosition] == '1' {
				rowsWithOnes = append(rowsWithOnes, row)
			} else {
				rowsWithZeroes = append(rowsWithZeroes, row)
			}
		}
		diffOnesZeroes := len(rowsWithOnes) - len(rowsWithZeroes)
		if diffOnesZeroes == 0 && comparatorFactor == 1 {
			workingRows = rowsWithOnes
		} else if diffOnesZeroes == 0 && comparatorFactor == -1 {
			workingRows = rowsWithZeroes
		} else if diffOnesZeroes*comparatorFactor >= 0 {
			workingRows = rowsWithOnes
		} else {
			workingRows = rowsWithZeroes
		}
	}
	return workingRows[0]
}

type Day4 struct{}

func (d Day4) solvePart1(input string) int {
	numberSequence, boards := parseInputToNumbersAndBoards(input)

	var winningBoard *Board
	var winningNumber int
	for _, numberAsString := range numberSequence {
		number, _ := strconv.Atoi(numberAsString)
		for _, board := range boards {
			board.cross(number)
			if board.isWinning() {
				winningBoard = board
				winningNumber = number
				break
			}
		}
		if winningBoard != nil {
			break
		}
	}
	uncrossedCellSum := 0
	for _, cell := range winningBoard.getUncrossedCells() {
		uncrossedCellSum = uncrossedCellSum + cell
	}

	return winningNumber * uncrossedCellSum
}

func (d Day4) solvePart2(input string) int {
	numberSequence, boards := parseInputToNumbersAndBoards(input)

	var lastWinningBoard *Board
	var lastWinningNumber int
	alreadyWonBoards := make(map[*Board]bool)
	for _, numberAsString := range numberSequence {
		number, _ := strconv.Atoi(numberAsString)
		for _, board := range boards {
			_, boardAlreadyWon := alreadyWonBoards[board]
			if boardAlreadyWon {
				continue
			}
			board.cross(number)
			if board.isWinning() {
				alreadyWonBoards[board] = true
				lastWinningBoard = board
				lastWinningNumber = number
			}
		}
		if len(alreadyWonBoards) == len(boards) {
			break
		}
	}
	uncrossedCellSum := 0
	for _, cell := range lastWinningBoard.getUncrossedCells() {
		uncrossedCellSum = uncrossedCellSum + cell
	}

	return lastWinningNumber * uncrossedCellSum
}

func parseInputToNumbersAndBoards(input string) ([]string, []*Board) {
	numberSequenceAndTheRest := strings.SplitAfterN(input, "\n", 2)
	numberSequence := strings.Split(numberSequenceAndTheRest[0], ",")
	boardsRaw := strings.Split(numberSequenceAndTheRest[1], "\n")
	boards := make([]*Board, 0)

	builder := CreateBoardBuilder()
	for _, rowRaw := range boardsRaw {
		if len(rowRaw) > 0 {
			builder.withRow(rowRaw)
		} else {
			board, success := builder.build()
			if success {
				boards = append(boards, board)
			}
			builder.reset()
		}
	}
	return numberSequence, boards
}

type Board struct {
	width              int
	height             int
	cells              []int
	crossedCells       map[int]bool
	horizontalsCrossed []int
	verticalsCrossed   []int
}

func (board *Board) cross(number int) {
	indexInCells := -1
	for i, cell := range board.cells {
		if number == cell {
			indexInCells = i
			break
		}
	}
	if indexInCells >= 0 {
		board.crossedCells[number] = true
		posX := indexInCells % board.width
		posY := int(math.Floor(float64(indexInCells) / float64(board.width)))
		board.verticalsCrossed[posX] = board.verticalsCrossed[posX] + 1
		board.horizontalsCrossed[posY] = board.horizontalsCrossed[posY] + 1
	}
}

func (board *Board) getUncrossedCells() []int {
	uncrossedCells := make([]int, 0)
	for _, cell := range board.cells {
		_, numberExists := board.crossedCells[cell]
		if !numberExists {
			uncrossedCells = append(uncrossedCells, cell)
		}
	}
	return uncrossedCells
}

func (board *Board) isWinning() bool {
	for _, value := range board.horizontalsCrossed {
		if value == board.width {
			return true
		}
	}
	for _, value := range board.verticalsCrossed {
		if value == board.height {
			return true
		}
	}
	return false
}

type BoardBuilder struct {
	rows []string
}

func CreateBoardBuilder() *BoardBuilder {
	return &BoardBuilder{make([]string, 0)}
}

func (builder *BoardBuilder) withRow(row string) {
	builder.rows = append(builder.rows, row)
}

func (builder *BoardBuilder) reset() {
	builder.rows = make([]string, 0)
}

func (builder *BoardBuilder) build() (*Board, bool) {
	cells := make([]int, 0)
	for _, row := range builder.rows {
		for _, numberAsString := range strings.Split(row, " ") {
			if len(numberAsString) > 0 {
				number, _ := strconv.Atoi(numberAsString)
				cells = append(cells, number)
			}
		}
	}
	height := len(builder.rows)
	if height == 0 {
		return nil, false
	}
	width := len(cells) / height
	return &Board{width, height, cells, make(map[int]bool, 0), make([]int, height, height), make([]int, width, width)}, true
}

type Day5 struct {
}

func (d Day5) solvePart1(input string) int {
	return solveDay5ConsideringDiagonals(input, false)
}

func (d Day5) solvePart2(input string) int {
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
