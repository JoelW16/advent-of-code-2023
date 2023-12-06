package main

import (
	"testing"
)

func TestGetTimeRangeProductToBeatRecord(t *testing.T) {
	raceTime := 7
	recordDistance := 9

	result := GetTimeRangeCountToBeatRecord(raceTime, recordDistance)

	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestGetProductOfTimeCount(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	result := GetProductOfTimeCount(input)

	if result != 288 {
		t.Errorf("Expected 288, got %d", result)
	}
}
