package main

import (
	"testing"
)

func TestIsPossibleGame(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	cubesInBag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	result := IsPossibleGame(input, cubesInBag)

	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestIsImpossibleGame(t *testing.T) {
	input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	cubesInBag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	result := IsPossibleGame(input, cubesInBag)

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestSumPossibleGames(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	cubesInBag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	result := SumPossibleGames(input, cubesInBag)

	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}
