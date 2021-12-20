package solution

import (
	"fmt"
	"strings"
)

type Day14 struct {
}

func (d Day14) SolvePart1(input string) int {
	return solveDay14WithSteps(input, 10)
}

func (d Day14) SolvePart2(input string) int {
	return solveDay14WithSteps(input, 40)
}

func solveDay14WithSteps(input string, steps int) int {
	rows := strings.Split(input, "\n")
	template := ""
	rulesByNeedle := make(map[string]Rule)
	for _, row := range rows {
		pairSplit := strings.Split(row, " -> ")
		if len(pairSplit) == 1 {
			if strings.TrimSpace(pairSplit[0]) != "" {
				template = pairSplit[0]
			}
		} else {
			left := rune(pairSplit[0][0])
			right := rune(pairSplit[0][1])
			ins := rune(pairSplit[1][0])
			rule := Rule{left, right, ins}
			rulesByNeedle[pairSplit[0]] = rule
		}
	}
	for i := 0; i < steps; i++ {
		template = runInsertion(template, rulesByNeedle)
	}
	countOfChar := make(map[rune]int)
	for _, char := range template {
		count, _ := countOfChar[char]
		countOfThis := count + 1
		countOfChar[char] = countOfThis
	}
	highestCount := 0
	lowestCount := 9999999999
	for _, counts := range countOfChar {
		if counts > highestCount {
			highestCount = counts
		}
		if counts < lowestCount {
			lowestCount = counts
		}
	}
	return highestCount - lowestCount
}

func runInsertion(template string, rulesByNeedle map[string]Rule) string {
	insertions := make([]Insertion, 0)
	for index := 0; index < len(template)-1; index++ {
		rule, _ := rulesByNeedle[fmt.Sprintf("%c%c", template[index], template[index+1])]
		insertions = append(insertions, Insertion{index + 1, rule.insertChar})
	}
	insertionsList := InsertionList{insertions}
	builder := strings.Builder{}
	lastIndex := len(template) - 1
	builder.WriteRune(rune(template[lastIndex]))
	for j := len(insertionsList.insertions) - 1; j >= 0; j-- {
		insertion := insertionsList.insertions[j]
		index := insertion.index
		if index < lastIndex {
			for k := lastIndex - 1; k >= index; k-- {
				builder.WriteRune(rune(template[k]))
			}
			lastIndex = index
		}
		builder.WriteRune(insertion.char)
		lastIndex = index
	}
	if 0 < lastIndex {
		for k := lastIndex - 1; k >= 0; k-- {
			builder.WriteRune(rune(template[k]))
		}
		lastIndex = 0
	}
	reversedNewTemplate := builder.String()
	reverseBuilder := strings.Builder{}
	for j := len(reversedNewTemplate) - 1; j >= 0; j-- {
		reverseBuilder.WriteRune(rune(reversedNewTemplate[j]))
	}
	template = reverseBuilder.String()
	return template
}

type Rule struct {
	left, right rune
	insertChar  rune
}

type Insertion struct {
	index int
	char  rune
}

type InsertionList struct {
	insertions []Insertion
}

func (list InsertionList) Len() int {
	return len(list.insertions)
}

func (list InsertionList) Less(i, j int) bool {
	return list.insertions[i].index < list.insertions[j].index
}

func (list InsertionList) Swap(i, j int) {
	temp := list.insertions[i]
	list.insertions[i] = list.insertions[j]
	list.insertions[j] = temp
}
