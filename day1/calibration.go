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
	code := calibrate(string(content))
	code2 := calibrateWithWords(string(content))
	fmt.Println(code)
	fmt.Println(code2)
}

func concatNumbers(firstDigit string, lastDigit string) int {
	concatNumbers := firstDigit + lastDigit
	out, _ := strconv.Atoi(concatNumbers)
	return out
}

func replaceWordsWithNumbers(input string) string {
	hack := input
	hack = strings.ReplaceAll(hack, "one", "o1e")
	hack = strings.ReplaceAll(hack, "two", "t2o")
	hack = strings.ReplaceAll(hack, "three", "t3e")
	hack = strings.ReplaceAll(hack, "four", "f4r")
	hack = strings.ReplaceAll(hack, "five", "f5e")
	hack = strings.ReplaceAll(hack, "six", "s6x")
	hack = strings.ReplaceAll(hack, "seven", "s7n")
	hack = strings.ReplaceAll(hack, "eight", "e8t")
	hack = strings.ReplaceAll(hack, "nine", "n9e")
	return hack
}

func calibrate(input string) int {
	numberRegex := regexp.MustCompile(`\d`)
	lines := strings.Fields(input)
	sum := 0
	for _, line := range lines {
		numbers := numberRegex.FindAllString(line, -1)
		sum += concatNumbers(numbers[0], numbers[len(numbers)-1])
	}
	return sum
}

func calibrateWithWords(input string) int {
	numberRegex := regexp.MustCompile(`\d`)
	lines := strings.Fields(input)
	sum := 0
	for _, line := range lines {
		hack := replaceWordsWithNumbers(line)
		numbers := numberRegex.FindAllString(hack, -1)
		sum += concatNumbers(numbers[0], numbers[len(numbers)-1])
	}
	return sum
}
