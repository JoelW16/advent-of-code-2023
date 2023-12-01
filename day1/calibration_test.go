package main

import (
	"testing"
)

func TestSimpleCalibrate(t *testing.T) {
	testInput := `1abc2`

	result := calibrate(testInput)

	if result != 12 {
		t.Errorf("Expected 12, got %d", result)
	}
}

func TestCalibrate(t *testing.T) {
	testInput := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	result := calibrate(testInput)

	if result != 142 {
		t.Errorf("Expected 142, got %d", result)
	}
}

func TestSimpleCalibrateWithWrittenDigits(t *testing.T) {
	testInput := `two1nine`

	result := calibrateWithWords(testInput)

	if result != 29 {
		t.Errorf("Expected 29, got %d", result)
	}
}

func TestSimple18CalibrateWithWrittenDigits(t *testing.T) {
	testInput := `oneight`

	result := calibrateWithWords(testInput)

	if result != 18 {
		t.Errorf("Expected 18, got %d", result)
	}
}

func TestCalibrateWithWrittenDigits(t *testing.T) {
	testInput := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	result := calibrateWithWords(testInput)

	if result != 281 {
		t.Errorf("Expected 281, got %d", result)
	}
}
