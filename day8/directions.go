package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type networkNode struct {
	currentLocation string
	left            string
	right           string
	index           int
}

var wordRegex = regexp.MustCompile(`([A-Z]*)\w+`)

func ExtractNetwork(input []string) (map[string]networkNode, map[int]networkNode, map[string]networkNode) {
	network := make(map[string]networkNode)
	startNodes := make(map[int]networkNode)
	endNodes := make(map[string]networkNode)
	for index, line := range input {
		networkParts := wordRegex.FindAllString(line, -1)
		network[networkParts[0]] = networkNode{
			currentLocation: networkParts[0],
			left:            networkParts[1],
			right:           networkParts[2],
			index:           index,
		}
		if networkParts[0][2] == 'A' {
			startNodes[index] = network[networkParts[0]]
		}
		if networkParts[0][2] == 'Z' {
			endNodes[networkParts[0]] = network[networkParts[0]]
		}
	}

	return network, startNodes, endNodes
}

func iterateTurn(turnIndex int, turnCount int) int {
	if turnIndex+1 < turnCount {
		return turnIndex + 1
	}
	return 0
}

func FindStepsToEnd(walkingPattern string, network map[string]networkNode) int {
	steps := 0
	currentLocation := "AAA"
	endNode := "ZZZ"
	turnIndex := 0
	for currentLocation != endNode {
		turn := walkingPattern[turnIndex]
		if turn == 'L' {
			currentLocation = network[currentLocation].left

		} else {
			currentLocation = network[currentLocation].right
		}
		steps++
		turnIndex = iterateTurn(turnIndex, len(walkingPattern))
	}

	return steps
}

func GDC(a int, b int) int {
	if b == 0 {
		return a
	}
	return GDC(b, a%b)
}

// Returns LCM of array elements
func FindLCM(arr []int, n int) int {
	ans := arr[0]
	for i := 1; i < n; i++ {

		ans = ((arr[i] * ans) /
			(GDC(arr[i], ans)))
	}

	return ans
}
func FindGhostStepsToEnd(walkingPattern string, network map[string]networkNode,
	ghosts map[int]networkNode, endNodes map[string]networkNode) int {
	ghostStepCounts := make([]int, 0)

	for _, ghost := range ghosts {
		stepCount := 0
		for i := 0; i < len(walkingPattern); i++ {
			if _, exists := endNodes[ghost.currentLocation]; exists {
				ghostStepCounts = append(ghostStepCounts, stepCount)
				break
			}

			if walkingPattern[i] == 'L' {
				ghost.currentLocation = network[ghost.currentLocation].left
			} else {
				ghost.currentLocation = network[ghost.currentLocation].right
			}
			ghosts[ghost.index] = ghost
			stepCount++

			if i == len(walkingPattern)-1 {
				i = -1
			}
		}
	}

	return FindLCM(ghostStepCounts, len(ghostStepCounts))
}

func GetStepsTakenToDestination(input string, isGhost bool) int {
	inputLines := strings.Split(input, "\n")
	walkingPattern := inputLines[0]
	network, startNodes, endNode := ExtractNetwork(inputLines[2:])

	if isGhost {
		return FindGhostStepsToEnd(walkingPattern, network, startNodes, endNode)
	}
	return FindStepsToEnd(walkingPattern, network)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := GetStepsTakenToDestination(string(content), false)
	fmt.Println("part 1:", part1)

	part2 := GetStepsTakenToDestination(string(content), true)
	fmt.Println("part 2:", part2)
}
