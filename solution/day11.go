package solution

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day11 struct {
}

func (d Day11) SolvePart1(input string) int {
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
	for i := 0; i < 100; i++ {
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
								cellpos := (y+k)*width + (x + l)
								if cellpos >= 0 && cellpos < len(cells) {
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
		printCells(cells, width)
	}
	return amountOfFlashes
}

func printCells(cells []int, width int) {
	for i, c := range cells {
		fmt.Fprintf(os.Stderr, "%d", c)
		if i > 0 && i%width == 0 {
			fmt.Fprintf(os.Stderr, "\n")
		}
	}
	fmt.Fprintf(os.Stderr, "\n---\n")
}

func (d Day11) SolvePart2(input string) int {
	//TODO implement me
	return 0
}
