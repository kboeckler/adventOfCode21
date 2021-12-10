package solution

import (
	"sort"
	"strings"
)

type Day9 struct {
}

func (d Day9) SolvePart1(input string) int {
	rows := strings.Split(input, "\n")
	width := len(rows[0])
	height := len(rows)
	cells := make([]int, 0, width*height)
	for _, row := range rows {
		for _, cell := range row {
			number := int(cell - '0')
			cells = append(cells, number)
		}
	}
	lowCellIndices := make([]int, 0)
	for i, number := range cells {
		topIndex := i - width
		bottomIndex := i + width
		leftIndex := i - 1
		rightIndex := i + 1
		if isCellHigherThenMe(topIndex, cells, number) && isCellHigherThenMe(bottomIndex, cells, number) && isCellHigherThenMe(leftIndex, cells, number) && isCellHigherThenMe(rightIndex, cells, number) {
			lowCellIndices = append(lowCellIndices, i)
		}
	}
	sum := 0
	for _, index := range lowCellIndices {
		height := cells[index]
		risk := height + 1
		sum = sum + risk
	}
	return sum
}

func isCellHigherThenMe(index int, cells []int, number int) bool {
	return index < 0 || index >= len(cells) || cells[index] > number
}

func (d Day9) SolvePart2(input string) int {
	rows := strings.Split(input, "\n")
	width := len(rows[0])
	height := len(rows)
	cells := make([]int, 0, width*height)
	for _, row := range rows {
		for _, cell := range row {
			number := int(cell - '0')
			cells = append(cells, number)
		}
	}
	lowCellIndices := make([]int, 0)
	for i, number := range cells {
		topIndex := i - width
		bottomIndex := i + width
		leftIndex := i - 1
		rightIndex := i + 1
		if isCellHigherThenMe(topIndex, cells, number) && isCellHigherThenMe(bottomIndex, cells, number) && isCellHigherThenMe(leftIndex, cells, number) && isCellHigherThenMe(rightIndex, cells, number) {
			lowCellIndices = append(lowCellIndices, i)
		}
	}
	basins := make([]Basin, 0)
	for _, index := range lowCellIndices {
		cellsInBasin := make(map[int]bool)
		openCells := make([]int, 0)
		openCells = append(openCells, index)
		for {
			if len(openCells) == 0 {
				break
			}
			currentCell := openCells[0]
			openCells = openCells[1:]
			visited, _ := cellsInBasin[currentCell]
			if visited {
				continue
			}
			cellsInBasin[currentCell] = true
			openCells = expandOpenCellsIfCellRisingToMe(currentCell-width, cells, cellsInBasin, currentCell, openCells)
			openCells = expandOpenCellsIfCellRisingToMe(currentCell+width, cells, cellsInBasin, currentCell, openCells)
			openCells = expandOpenCellsIfCellRisingToMe(currentCell-1, cells, cellsInBasin, currentCell, openCells)
			openCells = expandOpenCellsIfCellRisingToMe(currentCell+1, cells, cellsInBasin, currentCell, openCells)
		}
		basins = append(basins, Basin{cellsInBasin})
	}
	allBasins := Basins{basins}
	sort.Sort(allBasins)
	product := 1
	for i := 0; i < 3; i++ {
		basin := allBasins.basins[len(allBasins.basins)-i-1]
		product = product * len(basin.cellsInBasin)
	}
	return product
}

func expandOpenCellsIfCellRisingToMe(index int, cells []int, cellsInBasin map[int]bool, currentCell int, openCells []int) []int {
	if index >= 0 && index < len(cells) {
		visited, _ := cellsInBasin[index]
		if !visited && cells[index] > cells[currentCell] && cells[index] < 9 {
			openCells = append(openCells, index)
		}
	}
	return openCells
}

type Basin struct {
	cellsInBasin map[int]bool
}

type Basins struct {
	basins []Basin
}

func (b Basins) Len() int {
	return len(b.basins)
}

func (b Basins) Less(i int, j int) bool {
	return len(b.basins[i].cellsInBasin) < len(b.basins[j].cellsInBasin)
}

func (b Basins) Swap(i int, j int) {
	temp := b.basins[i]
	b.basins[i] = b.basins[j]
	b.basins[j] = temp
}
