package solution

import (
	"strconv"
	"strings"
)

type Day7 struct {
}

func (d Day7) SolvePart1(input string) int {
	numbersAsString := strings.Split(input, ",")
	numbers := make([]int, 0, len(numbersAsString))
	for _, numberAsString := range numbersAsString {
		number, _ := strconv.Atoi(numberAsString)
		numbers = append(numbers, number)
	}
	
	return 0
}

func (d Day7) SolvePart2(input string) int {
	//TODO implement me
	return 0
}
