package main

import (
	"testing"
)

func TestGetDistance(t *testing.T) {
	input := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	result := GetStepsTakenToDestination(input, false)

	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestGetDistanceExample2(t *testing.T) {
	input := `LLR

	AAA = (BBB, BBB)
	BBB = (AAA, ZZZ)
	ZZZ = (ZZZ, ZZZ)`

	result := GetStepsTakenToDestination(input, false)

	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestGetDistanceForGhosts(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	result := GetStepsTakenToDestination(input, true)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}

}
