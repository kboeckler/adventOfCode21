package main

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(workingDir + "/src/main/resources/com/github/kboeckler/adventOfCode/" + "Day3.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	input := string(data)

	for _, solution := range allSolutions() {
		result1 := solution.solvePart1(input)
		result2 := solution.solvePart2(input)
		log.Printf("Solution %v - Part1: %d Part2: %d\n", reflect.TypeOf(solution).Name(), result1, result2)
	}
}

func allSolutions() []Solution {
	var solutions = make([]Solution, 0)
	solutions = append(solutions, Day3{})
	return solutions
}

type Solution interface {
	solvePart1(input string) int
	solvePart2(input string) int
}

type Day3 struct {
}

func (d Day3) solvePart1(input string) int {
	rows := strings.Split(input, "\n")
	width := len(rows[0])
	height := len(rows)

	countOfOnesByPosition := make([]int, 0, width)
	for i := 0; i < width; i++ {
		countOfOnesByPosition = append(countOfOnesByPosition, 0)
	}

	for _, row := range rows {
		for i := range row {
			if row[i] == '1' {
				countOfOnesByPosition[i] = countOfOnesByPosition[i] + 1
			}
		}
	}

	gamma := 0
	epsilon := 0
	for i := width - 1; i >= 0; i-- {
		exponentOf2 := width - 1 - i
		if countOfOnesByPosition[i] >= height/2 {
			gamma = gamma + int(math.Pow(2, float64(exponentOf2)))
		} else {
			epsilon = epsilon + int(math.Pow(2, float64(exponentOf2)))
		}
	}
	return gamma * epsilon
}

func (d Day3) solvePart2(input string) int {
	rows := strings.Split(input, "\n")

	oxygen := determineRating(rows, 1)
	co2 := determineRating(rows, -1)
	oxygenDecimal, _ := strconv.ParseInt(oxygen, 2, 0)
	co2Decimal, _ := strconv.ParseInt(co2, 2, 0)

	return int(oxygenDecimal * co2Decimal)
}

func determineRating(rows []string, comparatorFactor int) string {
	width := len(rows[0])
	workingRows := rows
	for bitPosition := 0; bitPosition < width; bitPosition = bitPosition + 1 {
		if len(workingRows) <= 1 {
			break
		}
		rowsWithOnes := make([]string, 0)
		rowsWithZeroes := make([]string, 0)
		for _, row := range workingRows {
			if row[bitPosition] == '1' {
				rowsWithOnes = append(rowsWithOnes, row)
			} else {
				rowsWithZeroes = append(rowsWithZeroes, row)
			}
		}
		diffOnesZeroes := len(rowsWithOnes) - len(rowsWithZeroes)
		if diffOnesZeroes == 0 && comparatorFactor == 1 {
			workingRows = rowsWithOnes
		} else if diffOnesZeroes == 0 && comparatorFactor == -1 {
			workingRows = rowsWithZeroes
		} else if diffOnesZeroes*comparatorFactor >= 0 {
			workingRows = rowsWithOnes
		} else {
			workingRows = rowsWithZeroes
		}
	}
	return workingRows[0]
}
