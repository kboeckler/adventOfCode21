package solution

import (
	"fmt"
	"strconv"
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
	sum := 0
	for _, display := range displays {
		for _, letter := range display.two {
			display.Extend(letter, "cf")
		}
		for _, letter := range display.three {
			display.Extend(letter, "acf")
		}
		for _, letter := range display.four {
			display.Extend(letter, "bcdf")
		}
		for _, letter := range display.seven {
			display.Extend(letter, "abcdefg")
		}
		var validAssignment *Assignment = nil
		for _, assignment := range display.assignments {
			if display.fulfillsAll(assignment) {
				validAssignment = &assignment
			}
		}
		fourDigitAsString := ""
		for _, resolvedNumber := range display.resolvedOutvals(*validAssignment) {
			fourDigitAsString = fmt.Sprintf("%s%d", fourDigitAsString, resolvedNumber)
		}
		fourDigit, _ := strconv.ParseInt(fourDigitAsString, 10, 0)
		sum = sum + int(fourDigit)
	}
	return sum
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
	assignments             []Assignment
}

func (d Display) Extend(letter rune, values string) {
	// TODO
}

func (d Display) fulfillsAll(assignment Assignment) bool {
	// TODO
	return false
}

func (d Display) resolvedOutvals(assignment Assignment) []int {
	// TODO
	return make([]int, 0)
}

func CreateDisplay(invals, outvals []string) *Display {
	var two string
	var three string
	var four string
	var seven string
	others := make([]string, 0)
	for _, val := range invals {
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
	display := Display{invals, outvals, two, three, four, seven, others, make([]Assignment, 0)}
	return &display
}

type Assignment struct {
	vals map[rune]rune
}

func CreateAssignment() Assignment {
	return Assignment{make(map[rune]rune)}
}
