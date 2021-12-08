package solution

import (
	"math"
	"sort"
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
	sort.Ints(numbers)
	middleNumber := numbers[len(numbers)/2-1]
	fuel := 0
	for _, number := range numbers {
		delta := int(math.Abs(float64(number - middleNumber)))
		fuel = fuel + delta
	}
	return fuel
}

func (d Day7) SolvePart2(input string) int {
	numbersAsString := strings.Split(input, ",")
	numbers := make([]int, 0, len(numbersAsString))
	for _, numberAsString := range numbersAsString {
		number, _ := strconv.Atoi(numberAsString)
		numbers = append(numbers, number)
	}
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	average := int(math.Round(float64(sum) / float64(len(numbers))))
	fuel := 0
	for _, number := range numbers {
		delta := int(math.Abs(float64(number - average)))
		fuel = fuel + delta
	}
	return fuel
}
