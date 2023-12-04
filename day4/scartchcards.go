package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := FindSumOfWinningValues(string(content))
	fmt.Println(part1)
}

func IsWinningNumber(number string, winningNumbers []string) bool {
	for _, winningNumber := range winningNumbers {
		if number == winningNumber {
			return true
		}
	}
	return false
}

func FindWinningValue(input string) int {
	card := strings.Split(input, ":")
	numbers := strings.Split(card[1], "|")
	numberRegex := regexp.MustCompile(`[0-9]+`)
	playerNumbers := numberRegex.FindAllString(numbers[0], -1)
	winningNumbers := numberRegex.FindAllString(numbers[1], -1)
	score := 0
	for _, playerNumber := range playerNumbers {
		if IsWinningNumber(playerNumber, winningNumbers) {
			if score == 0 {
				score++
			} else {
				score = score * 2
			}
		}
	}
	return score
}

func FindSumOfWinningValues(input string) int {
	cards := strings.Split(input, "\n")
	totalScore := 0
	for _, card := range cards {
		score := FindWinningValue(card)
		if score > 0 {
			totalScore += score
		}
	}
	return totalScore
}
