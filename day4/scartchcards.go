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

	part1 := FindSumOfWinningValues(string(content))
	fmt.Println(part1)

	part2 := FindSumOfPlayableCards(string(content))
	fmt.Println(part2)
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

func UpdatePlayableCardCount(input string, cardCountMap map[int]int) map[int]int {
	card := strings.Split(input, ":")
	numbers := strings.Split(card[1], "|")
	numberRegex := regexp.MustCompile(`[0-9]+`)
	playerNumbers := numberRegex.FindAllString(numbers[0], -1)
	winningNumbers := numberRegex.FindAllString(numbers[1], -1)
	cardNumber, _ := strconv.Atoi(numberRegex.FindString(card[0]))
	score := 0

	for _, playerNumber := range playerNumbers {
		if IsWinningNumber(playerNumber, winningNumbers) {
			score++
		}
	}
	cardCountMap[cardNumber] += 1
	for i := 1; i <= score; i++ {
		cardCountMap[cardNumber+i] += cardCountMap[cardNumber]
	}

	return cardCountMap
}

func FindSumOfPlayableCards(input string) int {
	cards := strings.Split(input, "\n")
	cardCountMap := make(map[int]int)
	for _, card := range cards {
		cardCountMap = UpdatePlayableCardCount(card, cardCountMap)
	}

	sum := 0
	for _, count := range cardCountMap {
		sum += count
	}
	return sum
}
