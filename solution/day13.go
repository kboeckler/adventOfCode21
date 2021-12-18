package solution

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day13 struct {
}

func (d Day13) SolvePart1(input string) int {
	folds, width, dots := parseDay13(input)
	dots = foldDots(folds[0], dots, width)
	return len(dots)
}

func (d Day13) SolvePart2(input string) int {
	folds, width, dots := parseDay13(input)
	for _, fold := range folds {
		dots = foldDots(fold, dots, width)
	}
	remainingWidth := 0
	remainingHeight := 0
	for _, dot := range dots {
		if remainingWidth < dot.x+1 {
			remainingWidth = dot.x + 1
		}
		if remainingHeight < dot.y+1 {
			remainingHeight = dot.y + 1
		}
	}

	log.Println("Solution Day 13 - Part2 as Print:")
	for y := 0; y < remainingHeight; y++ {
		for x := 0; x < remainingWidth; x++ {
			_, hasDot := dots[x+y*width]
			if hasDot {
				fmt.Fprintf(os.Stdout, "#")
			} else {
				fmt.Fprintf(os.Stdout, ".")
			}
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
	return len(dots)
}

func foldDots(fold string, dots map[int]Dot, width int) map[int]Dot {
	foldCommand := strings.Split(fold, "=")
	if foldCommand[0] == "x" {
		foldColumn, _ := strconv.Atoi(foldCommand[1])
		newdots := make(map[int]Dot, 0)
		for _, dot := range dots {
			if dot.x < foldColumn {
				newdots[dot.x+dot.y*width] = dot
			} else if dot.x > foldColumn {
				newX := 2*foldColumn - dot.x
				newDot := Dot{newX, dot.y}
				newdots[newDot.x+newDot.y*width] = newDot
			}
		}
		dots = newdots
	} else {
		foldRow, _ := strconv.Atoi(foldCommand[1])
		newdots := make(map[int]Dot, 0)
		for _, dot := range dots {
			if dot.y < foldRow {
				newdots[dot.x+dot.y*width] = dot
			} else if dot.y > foldRow {
				newY := 2*foldRow - dot.y
				newDot := Dot{dot.x, newY}
				newdots[newDot.x+newDot.y*width] = newDot
			}
		}
		dots = newdots
	}
	return dots
}

func parseDay13(input string) ([]string, int, map[int]Dot) {
	rows := strings.Split(input, "\n")
	dotsAsArray := make([]Dot, 0)
	folds := make([]string, 0)
	width := 0
	for _, row := range rows {
		if len(strings.TrimSpace(row)) == 0 {
			continue
		}
		foldReplacement := strings.ReplaceAll(row, "fold along ", "")
		if strings.EqualFold(foldReplacement, row) {
			split := strings.Split(row, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			dot := Dot{x, y}
			dotsAsArray = append(dotsAsArray, dot)
			if width < x+1 {
				width = x + 1
			}
		} else {
			folds = append(folds, foldReplacement)
		}
	}
	dots := make(map[int]Dot, 0)
	for _, dot := range dotsAsArray {
		dots[dot.x+dot.y*width] = dot
	}
	return folds, width, dots
}

type Dot struct {
	x, y int
}
