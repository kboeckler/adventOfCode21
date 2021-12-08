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
	maxNumber := -1
	for _, numberAsString := range numbersAsString {
		number, _ := strconv.Atoi(numberAsString)
		numbers = append(numbers, number)
		if number > maxNumber {
			maxNumber = number
		}
	}
	bestFuel := 999999999
	for i := 0; i <= maxNumber; i++ {
		fuel := 0
		for _, number := range numbers {
			dist := math.Abs(float64(number - i))
			fuelPart := int((dist * (dist + 1)) / 2)
			fuel = fuel + fuelPart
		}
		if fuel < bestFuel {
			bestFuel = fuel
		}
	}
	return bestFuel
}
