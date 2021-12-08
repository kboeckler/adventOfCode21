package solution

import (
	"strings"
)

type Day8 struct {
}

func (d Day8) SolvePart1(input string) int {
	displays := parseDisplays(input)
	uniqueNumbers := 0
	for _, display := range displays {
		for _, val := range display.outvals {
			if len(val) == 2 || len(val) == 3 || len(val) == 4 || len(val) == 7 {
				uniqueNumbers = uniqueNumbers + 1
			}
		}
	}
	return uniqueNumbers
}

func (d Day8) SolvePart2(input string) int {
	displays := parseDisplays(input)
	for _, display := range displays {
		_ = display
		for _, part := range display.two {
			_ = part
		}
		_ = Setting{nil, nil, nil, nil, nil, nil, nil}
	}
	return 0
}

func parseDisplays(input string) []*Display {
	rows := strings.Split(input, "\n")
	displays := make([]*Display, 0)
	for _, row := range rows {
		invals := make([]string, 0)
		outvals := make([]string, 0)
		rawrows := strings.Split(row, " | ")
		inrow := strings.Split(rawrows[0], " ")
		outrow := strings.Split(rawrows[1], " ")
		for _, val := range inrow {
			invals = append(invals, val)
		}
		for _, val := range outrow {
			outvals = append(outvals, val)
		}
		display := CreateDisplay(invals, outvals)
		displays = append(displays, display)
	}
	return displays
}

type Display struct {
	invals, outvals         []string
	two, three, four, seven string
	others                  []string
}

func CreateDisplay(invals, outvals []string) *Display {
	var two string
	var three string
	var four string
	var seven string
	others := make([]string, 0)
	for _, val := range outvals {
		if len(val) == 2 {
			two = val
		} else if len(val) == 3 {
			three = val
		} else if len(val) == 4 {
			four = val
		} else {
			others = append(others, val)
		}
	}
	display := Display{invals, outvals, two, three, four, seven, others}
	return &display
}

type Setting struct {
	a, b, c, d, e, f, g *rune
}
