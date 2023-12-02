package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	cubesInBag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	task1 := SumPossibleGames(string(content), cubesInBag)
	fmt.Println(task1)
}

func Split(r rune) bool {
	return r == ',' || r == ';'
}

func IsPossibleGame(input string, cubesInBag map[string]int) int {
	game := strings.Split(input, ":")
	cubes := strings.FieldsFunc(game[1], Split)

	for _, cube := range cubes {
		cubeKV := strings.Fields(cube)
		cubeValue, _ := strconv.Atoi(cubeKV[0])
		if cubeValue > cubesInBag[cubeKV[1]] {
			return 0
		}
	}

	gameNumber, _ := strconv.Atoi(strings.Fields(game[0])[1])
	return gameNumber
}

func SumPossibleGames(input string, cubesInBag map[string]int) int {
	games := strings.Split(input, "\n")
	sum := 0
	for _, game := range games {
		sum += IsPossibleGame(game, cubesInBag)
	}
	return sum
}
