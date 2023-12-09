package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isEnd(history []int) bool {
	for _, i := range history {
		if i != 0 {
			return false
		}
	}
	return true
}

func GetDiff(history []int) []int {
	diff := make([]int, len(history)-1)
	for i := 0; i < len(history)-1; i++ {
		diff[i] = history[i+1] - history[i]
	}
	return diff
}

func FindNextValue(history []int) int {
	if history[len(history)-1] == 0 && isEnd(history) {
		return 0
	}

	diff := GetDiff(history)

	res := FindNextValue(diff)
	return history[len(history)-1] + res
}

func FindPreviousValue(history []int) int {
	if history[len(history)-1] == 0 && isEnd(history) {
		return 0
	}

	diff := GetDiff(history)

	res := FindPreviousValue(diff)
	return history[0] - res
}

func ConvertToIntArray(input string) []int {
	row := strings.Split(input, " ")
	numbers := []int{}
	for _, str := range row {
		num, _ := strconv.Atoi(str)
		numbers = append(numbers, num)
	}
	return numbers
}

func FindSum(history string, isNext bool) int {
	rows := strings.Split(history, "\n")
	sum := 0
	for _, row := range rows {
		numbers := ConvertToIntArray(row)
		if isNext {
			sum += FindNextValue(numbers)
		} else {
			sum += FindPreviousValue(numbers)
		}
	}
	return sum
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := FindSum(string(content), true)
	fmt.Println("part 1:", part1)

	part2 := FindSum(string(content), false)
	fmt.Println("part 2:", part2)

}
