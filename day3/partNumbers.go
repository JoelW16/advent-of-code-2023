package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
}

func CreatePartMap(input string) []string {
	return strings.Split(input, "\n")
}

func CheckAdjacentParts(partMap []string, x int, y int) bool {
	isTop := partMap[y][x-1] == '*'
	isBottom := partMap[y][x+1] == '*'
	isRight := partMap[y-1][x] == '*'
	isLeft := partMap[y+1][x] == '*'

	isTopRight := partMap[y-1][x-1] == '*'
	isBottomRight := partMap[y+1][x-1] == '*'
	isTopLeft := partMap[y-1][x+1] == '*'
	isBottomLeft := partMap[y+1][x+1] == '*'

	return isTop || isBottom || isRight || isLeft || isTopRight || isBottomRight || isTopLeft || isBottomLeft
}

func FindPartNumbers(input string) []int {
	partMap := CreatePartMap(input)
	fmt.Println(partMap)

	for y, row := range partMap {
		fmt.Println(y, row)
		for x, part := range row {
			if _, err := strconv.Atoi(string(part)); err == nil {
				fmt.Printf("%q looks like a number.\n", string(part))
				check := CheckAdjacentParts(partMap, x, y)
				fmt.Println(check)
			}
		}
	}
	return []int{1, 12}
}
