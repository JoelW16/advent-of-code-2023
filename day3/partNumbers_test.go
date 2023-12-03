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

	if len(result) != 2 || result[0] != 467 || result[1] != 35 {
		t.Errorf("Expected [467 35], got %d", result)
	}
}
