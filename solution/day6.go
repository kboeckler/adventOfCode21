package solution

import (
	"math"
	"strconv"
	"strings"
)

type Day6 struct {
}

func (d Day6) SolvePart1(input string) int {
	timersAsStrings := strings.Split(input, ",")
	timers := make([]int, 0, len(timersAsStrings))
	for _, timerAsString := range timersAsStrings {
		timer, _ := strconv.Atoi(timerAsString)
		timers = append(timers, timer)
	}
	for i := 0; i < 80; i++ {
		for j, timer := range timers {
			if timer == 0 {
				timers[j] = 6
				timers = append(timers, 8)
			} else {
				timers[j] = timer - 1
			}
		}
	}
	return len(timers)
}

func (d Day6) SolvePart2(input string) int {
	timersAsStrings := strings.Split(input, ",")
	fishesPerTimer := make(map[int]int)
	for _, timerAsString := range timersAsStrings {
		timer, _ := strconv.Atoi(timerAsString)
		fishesPerThisTimer, _ := fishesPerTimer[timer]
		fishesPerTimer[timer] = fishesPerThisTimer + 1
	}
	daysLeft := 256
	for {
		nextTimer := 999
		for timer := range fishesPerTimer {
			if timer < daysLeft && timer < nextTimer {
				nextTimer = timer
			}
		}
		if nextTimer == 999 {
			break
		}
		fishes := fishesPerTimer[nextTimer]
		// TODO
		_ = fishes
	}

	fishesAfter80Days := 0
	for timer, fishes := range fishesPerTimer {
		reproductionTimes := math.Floor(float64((80 - timer) / 7))
		fishesAfter80Days = fishesAfter80Days + fishes*int(math.Pow(2, reproductionTimes))
	}
	return fishesAfter80Days
}
