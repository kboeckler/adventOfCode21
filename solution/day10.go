package solution

import (
	"sort"
	"strings"
)

type Day10 struct {
}

func (d Day10) SolvePart1(input string) int {
	rows := strings.Split(input, "\n")
	illegalCharactersCount := make(map[rune]int)
	for _, row := range rows {
		index, _ := getIndexOfCorruptedCharacterOrRemainingClosingBrackets(row)
		if index >= 0 {
			char := row[index]
			countBefore, _ := illegalCharactersCount[rune(char)]
			illegalCharactersCount[rune(char)] = countBefore + 1
		}
	}
	result := 0
	for char, count := range illegalCharactersCount {
		var points int
		switch char {
		case ')':
			points = 3
		case ']':
			points = 57
		case '}':
			points = 1197
		case '>':
			points = 25137
		default:
			panic("Invalid char")
		}
		result = result + count*points
	}
	return result
}

func (d Day10) SolvePart2(input string) int {
	rows := strings.Split(input, "\n")
	scores := make([]int, 0)
	for _, row := range rows {
		index, brackets := getIndexOfCorruptedCharacterOrRemainingClosingBrackets(row)
		if index == -1 && len(*brackets) > 0 {
			score := 0
			for i := len(*brackets) - 1; i >= 0; i-- {
				bracket := (*brackets)[i]
				score = score * 5
				switch bracket {
				case ')':
					score = score + 1
				case ']':
					score = score + 2
				case '}':
					score = score + 3
				case '>':
					score = score + 4
				default:
					panic("Invalid char")
				}
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func getIndexOfCorruptedCharacterOrRemainingClosingBrackets(row string) (int, *[]rune) {
	closingBrackets := make([]rune, 0)
	for i, char := range row {
		if char == '(' {
			closingBrackets = append(closingBrackets, ')')
		} else if char == '[' {
			closingBrackets = append(closingBrackets, ']')
		} else if char == '{' {
			closingBrackets = append(closingBrackets, '}')
		} else if char == '<' {
			closingBrackets = append(closingBrackets, '>')
		} else {
			nextClosing := closingBrackets[len(closingBrackets)-1]
			if char == nextClosing {
				closingBrackets = closingBrackets[0 : len(closingBrackets)-1]
			} else {
				return i, nil
			}
		}
	}
	return -1, &closingBrackets
}
