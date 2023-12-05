package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FindMappedValue(sourceValue int, sourceMap string) int {
	sourceMapRows := strings.Split(sourceMap, "\n")

	for _, sourceMapRow := range sourceMapRows {
		values := strings.Fields(sourceMapRow)
		location, _ := strconv.Atoi(values[0])
		rangeStart, _ := strconv.Atoi(values[1])
		rangeLength, _ := strconv.Atoi(values[2])

		if sourceValue < rangeStart {
			continue
		}

		if sourceValue >= rangeStart && sourceValue < rangeStart+rangeLength {
			return location + (sourceValue - rangeStart)
		}
	}

	return sourceValue
}

func GetSeedValues(input string) []int {
	numberRegex := regexp.MustCompile(`[0-9]+`)
	seedRow := strings.Split(input, "\n")[0]
	seeds := numberRegex.FindAllString(seedRow, -1)
	seedValues := []int{}

	for _, seed := range seeds {
		seedValue, _ := strconv.Atoi(seed)
		seedValues = append(seedValues, seedValue)
	}

	return seedValues
}

func GetSourceMaps(almanac string) map[int]string {
	maps := make(map[int]string)
	almanacRows := strings.Split(almanac, "\n")[1:]
	mapIndex := 0
	numberRegex := regexp.MustCompile(`[0-9]+`)
	for _, row := range almanacRows {
		if strings.Contains(row, "map:") {
			mapIndex++
			continue
		}
		if numberRegex.MatchString(row) {
			if value, found := maps[mapIndex]; found {
				maps[mapIndex] = value + "\n" + row
			} else {
				maps[mapIndex] = row
			}
		}
	}
	return maps
}

func FindMinLocationValue(almanac string) int {
	seeds := GetSeedValues(almanac)
	maps := GetSourceMaps(almanac)
	min := 0
	for index, seed := range seeds {
		sourceValue := seed
		for mapIndex, _ := range maps {
			sourceValue = FindMappedValue(sourceValue, maps[mapIndex])
		}
		if index == 0 || sourceValue < min {
			min = sourceValue
		}
	}
	return min
}

func FindMinMappedValueWithRange(sourceValue int, sourceMap string) int {
	sourceMapRows := strings.Split(sourceMap, "\n")
	min := 0

	for _, sourceMapRow := range sourceMapRows {
		values := strings.Fields(sourceMapRow)
		location, _ := strconv.Atoi(values[0])
		rangeStart, _ := strconv.Atoi(values[1])
		rangeLength, _ := strconv.Atoi(values[2])
		rangeEnd := rangeStart + rangeLength

		if sourceValue < rangeStart {
			continue
		}

		if sourceValue >= rangeStart && sourceValue < rangeEnd {
			min = location + (sourceValue - rangeStart)
		}

	}

	return min
}

func FindMinLocationValueFromRange(almanac string) int {
	seeds := GetSeedValues(almanac)
	maps := GetSourceMaps(almanac)
	min := 0
	for seedSetIndex := 0; seedSetIndex <= len(seeds)-2; seedSetIndex = seedSetIndex + 2 {
		sourceValue := seeds[seedSetIndex]
		for mapIndex := 1; mapIndex <= len(maps); mapIndex++ {
			sourceValue = FindMinMappedValueWithRange(sourceValue, maps[mapIndex])
		}

		if min == 0 || sourceValue < min {
			min = sourceValue
		}

	}
	return min
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := FindMinLocationValue(string(content))
	fmt.Println("part 1:", part1)

	part2 := FindMinLocationValueFromRange(string(content))
	fmt.Println("part 2:", part2)
}
