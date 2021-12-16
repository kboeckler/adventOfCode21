package solution

import (
	"strconv"
	"strings"
)

type Day11 struct {
}

func (d Day11) SolvePart1(input string) int {
	return solveDay11(input, func(step, flashesInThisRound, flashesTotal int) (bool, int) {
		return step == 99, flashesTotal
	})
}

func (d Day11) SolvePart2(input string) int {
	return solveDay11(input, func(step, flashesInThisRound, flashesTotal int) (bool, int) {
		return flashesInThisRound == 100, step + 1
	})
}

func solveDay11(input string, cond func(int, int, int) (bool, int)) int {
	rows := strings.Split(input, "\n")
	cells := make([]int, 0)
	width := len(rows[0])
	for _, row := range rows {
		for _, cell := range row {
			intval, _ := strconv.Atoi(string(cell))
			cells = append(cells, intval)
		}
	}
	amountOfFlashes := 0
	for i := 0; true; i++ {
		for j := range cells {
			cells[j] = cells[j] + 1
		}
		flashes := make(map[int]bool)
		noMoreFlashes := false
		for {
			if noMoreFlashes {
				break
			}
			noMoreFlashes = true
			for j := range cells {
				if cells[j] > 9 {
					x := j % width
					y := j / width
					flashedAlready, _ := flashes[j]
					if !flashedAlready {
						noMoreFlashes = false
						flashes[y*width+x] = true
						for k := -1; k <= 1; k++ {
							for l := -1; l <= 1; l++ {
								othery := (y + k) * width
								otherx := x + l
								cellpos := othery + otherx
								if cellpos >= 0 && cellpos < len(cells) && otherx >= 0 && otherx < width {
									cells[cellpos] = cells[cellpos] + 1
								}
							}
						}
					}
				}
			}
		}
		amountOfFlashes = amountOfFlashes + len(flashes)
		for j := range cells {
			if cells[j] > 9 {
				cells[j] = 0
			}
		}
		end, result := cond(i, len(flashes), amountOfFlashes)
		if end {
			return result
		}
	}
	return -1
}
