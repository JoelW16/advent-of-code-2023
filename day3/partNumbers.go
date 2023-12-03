package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	part1 := FindPartNumbers(string(content))
	fmt.Println(part1)
}

func CreatePartMap(input string) []string {
	return strings.Split(input, "\n")
}

func SearchIndexes(partMap []string, startIndex int, endIndex int, row int) bool {
	symbolRegex := regexp.MustCompile(`[^0-9.]`)
	search := partMap[row][startIndex : endIndex+1]
	return symbolRegex.MatchString(search)
}

func GetStartIndex(numberXIndexes []int) int {
	startIndex := numberXIndexes[0]
	if startIndex > 0 {
		startIndex--
	}
	return startIndex
}

func GetEndIndex(numberXIndexes []int, partMap []string, row int) int {
	endIndex := numberXIndexes[len(numberXIndexes)-1] - 1
	if endIndex < len(partMap[row])-1 {
		endIndex++
	}
	return endIndex
}

func isSymbolAbove(partMap []string, numberXIndexes []int, row int) bool {
	if row == 0 {
		return false
	}
	startIndex := GetStartIndex(numberXIndexes)
	endIndex := GetEndIndex(numberXIndexes, partMap, row)
	return SearchIndexes(partMap, startIndex, endIndex, row-1)
}

func isSymbolBelow(partMap []string, numberXIndexes []int, row int) bool {
	if row == len(partMap)-1 {
		return false
	}
	startIndex := GetStartIndex(numberXIndexes)
	endIndex := GetEndIndex(numberXIndexes, partMap, row)
	return SearchIndexes(partMap, startIndex, endIndex, row+1)
}

func isSymbolOnLeft(partMap []string, numberXIndexes []int, row int) bool {
	if numberXIndexes[0] == 0 {
		return false
	}
	startIndex := GetStartIndex(numberXIndexes)
	return SearchIndexes(partMap, startIndex, startIndex, row)
}

func isSymbolOnRight(partMap []string, numberXIndexes []int, row int) bool {
	if numberXIndexes[len(numberXIndexes)-1] == len(partMap[row])-1 {
		return false
	}
	endIndex := GetEndIndex(numberXIndexes, partMap, row)
	return SearchIndexes(partMap, endIndex, endIndex, row)
}

func CheckAdjacentSymbols(partMap []string, numberXIndexes []int, row int) bool {
	return isSymbolAbove(partMap, numberXIndexes, row) || isSymbolBelow(partMap, numberXIndexes, row) || isSymbolOnLeft(partMap, numberXIndexes, row) || isSymbolOnRight(partMap, numberXIndexes, row)
}

func FindPartNumbers(input string) int {
	partMap := CreatePartMap(input)
	numberRegex := regexp.MustCompile(`\d+`)
	sum := 0
	for y, row := range partMap {
		numberIndexes := numberRegex.FindAllStringIndex(row, -1)
		numbers := numberRegex.FindAllString(row, -1)
		for i, numberIndex := range numberIndexes {
			if CheckAdjacentSymbols(partMap, numberIndex, y) {
				num, _ := strconv.Atoi(numbers[i])
				sum += num
			}
		}
	}
	return sum
}
