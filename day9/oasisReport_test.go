package main

import "testing"

func TestFindNextValue(t *testing.T) {
	history := map[int][]int{
		18: {0, 3, 6, 9, 12, 15},
		28: {1, 3, 6, 10, 15, 21},
		68: {10, 13, 16, 21, 30, 45},
	}

	for expected, input := range history {

		report := FindNextValue(input)

		if report != expected {
			t.Errorf("Expected next value to be %d, got %d", expected, report)
		}
	}
}

func TestFindPreviousValue(t *testing.T) {
	history := []int{10, 13, 16, 21, 30, 45}

	report := FindPreviousValue(history)

	if report != 5 {
		t.Errorf("Expected previous value to be 5, got %d", report)
	}
}

func TestFindSum(t *testing.T) {
	history := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	result := FindSumNext(history)

	if result != 114 {
		t.Errorf("Expected sum to be 114, got %d", result)
	}
}
