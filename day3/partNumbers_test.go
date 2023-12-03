package main

import (
	"testing"
)

func TestFindPartNumbersSmall(t *testing.T) {
	input :=
		`467..114..
...*......
..35..633.`
	result := FindPartNumbers(input)

	if result != 502 {
		t.Errorf("Expected 502, got %d", result)
	}
}

func TestFindPartNumbers(t *testing.T) {
	input :=
		`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	result := FindPartNumbers(input)

	if result != 4361 {
		t.Errorf("Expected 4361, got %d", result)
	}
}

func TestFindGearRatioSmall(t *testing.T) {
	input :=
		`467..114..
...*......
..35..633.`
	result := FindGearRatio(input)

	if result != 16345 {
		t.Errorf("Expected 16345, got %d", result)
	}
}

func TestFindGearRatio(t *testing.T) {
	input :=
		`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	result := FindGearRatio(input)

	if result != 467835 {
		t.Errorf("Expected 467835, got %d", result)
	}
}
