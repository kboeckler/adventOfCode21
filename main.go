package main

import (
	"github.com/kboeckler/adventOfCode21/solution"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, sol := range allSolutions() {
		nameOfSolution := reflect.TypeOf(sol).Name()
		data, err := ioutil.ReadFile(workingDir + "/src/main/resources/com/github/kboeckler/adventOfCode/" + nameOfSolution + ".txt")
		if err != nil {
			log.Fatal("File reading error", err)
		}
		input := strings.ReplaceAll(string(data), "\r\n", "\n")
		result1 := sol.SolvePart1(input)
		result2 := sol.SolvePart2(input)
		log.Printf("Solution %v - Part1: %d Part2: %d\n", nameOfSolution, result1, result2)
	}
}

func allSolutions() []solution.Solution {
	var solutions = make([]solution.Solution, 0)
	solutions = append(solutions, solution.Day3{})
	solutions = append(solutions, solution.Day4{})
	solutions = append(solutions, solution.Day5{})
	solutions = append(solutions, solution.Day6{})
	solutions = append(solutions, solution.Day7{})
	//solutions = append(solutions, solution.Day8{})
	return solutions
}
