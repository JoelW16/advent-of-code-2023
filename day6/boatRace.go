package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetTimeRangeCountToBeatRecord(raceTime int, recordDistance int) int {
	count := 0
	for speed := 0; speed <= raceTime; speed++ {
		remainingTime := raceTime - speed
		distance := speed * remainingTime
		if distance > recordDistance {
			count++
		}
	}
	return count
}

func GetProductOfTimeCount(input string) int {
	numberRegex := regexp.MustCompile(`[0-9]+`)
	inputRows := strings.Split(input, "\n")

	raceTimes := numberRegex.FindAllString(inputRows[0], -1)
	recordDistances := numberRegex.FindAllString(inputRows[1], -1)
	product := 1

	for i := 0; i < len(raceTimes); i++ {
		raceTime, _ := strconv.Atoi(raceTimes[i])
		recordDistance, _ := strconv.Atoi(recordDistances[i])
		product *= GetTimeRangeCountToBeatRecord(raceTime, recordDistance)
	}
	return product
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := GetProductOfTimeCount(string(content))
	fmt.Println("part 1:", part1)
}
