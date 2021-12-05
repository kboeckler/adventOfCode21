package solution

import (
	"math"
	"strconv"
	"strings"
)

type Day4 struct{}

func (d Day4) SolvePart1(input string) int {
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

func (d Day4) SolvePart2(input string) int {
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
