package solution

import (
	"strconv"
	"strings"
)

type Day6 struct {
}

func (d Day6) SolvePart1(input string) int {
	return solveDay6WithDays(input, 80)
}

func (d Day6) SolvePart2(input string) int {
	return solveDay6WithDays(input, 256)
}

func solveDay6WithDays(input string, days int) int {
	timersAsStrings := strings.Split(input, ",")
	fishesPerTimer := make(map[int]int)
	for _, timerAsString := range timersAsStrings {
		timer, _ := strconv.Atoi(timerAsString)
		fishesPerThisTimer, _ := fishesPerTimer[timer]
		fishesPerTimer[timer] = fishesPerThisTimer + 1
	}
	for i := 0; i < days; i++ {
		nextFishesPerTimer := make(map[int]int)
		for timer, fishes := range fishesPerTimer {
			nextTimer := timer - 1
			if nextTimer == -1 {
				nextTimer = 6
				nextFishesPerTimer[8] = fishes
			}
			nextFishes, _ := nextFishesPerTimer[nextTimer]
			nextFishesPerTimer[nextTimer] = nextFishes + fishes
		}
		fishesPerTimer = nextFishesPerTimer
	}
	allFishes := 0
	for _, fishes := range fishesPerTimer {
		allFishes = allFishes + fishes
	}
	return allFishes
}
