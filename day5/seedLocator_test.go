package main

import (
	"testing"
)

func TestFindMappedValueInRange(t *testing.T) {
	seedValue := 79
	seedToSoilMap := `50 98 2
52 50 48`
	result := FindMappedValue(seedValue, seedToSoilMap)

	if result != 81 {
		t.Errorf("Expected 81, got %d", result)
	}
}

func TestFindMappedValueOutOfRange(t *testing.T) {
	seedValue := 100
	seedToSoilMap := `50 98 2
52 50 48`
	result := FindMappedValue(seedValue, seedToSoilMap)

	if result != 100 {
		t.Errorf("Expected 100, got %d", result)
	}
}

func TestFindMappedValueLowerBound(t *testing.T) {
	seedValue := 50
	seedToSoilMap := `50 98 2
52 50 48`
	result := FindMappedValue(seedValue, seedToSoilMap)

	if result != 52 {
		t.Errorf("Expected 52, got %d", result)
	}
}

func TestFindMappedValueUpperBound(t *testing.T) {
	seedValue := 99
	seedToSoilMap := `50 98 2
52 50 48`
	result := FindMappedValue(seedValue, seedToSoilMap)

	if result != 51 {
		t.Errorf("Expected 51, got %d", result)
	}
}

func TestGetSeedValues(t *testing.T) {
	input := `seeds: 79 14 55 13`
	result := GetSeedValues(input)

	if len(result) != 4 {
		t.Errorf("Expected list of 4 ints, got %d", len(result))
	}
}

func TestFindMinLocationValue(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	result := FindMinLocationValue(input)
	if result != 35 {
		t.Errorf("Expected 35, got %d", result)
	}
}

// func TestFindMinLocationValueFromRange(t *testing.T) {
// 	input := `seeds: 79 14 55 13

// seed-to-soil map:
// 50 98 2
// 52 50 48

// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4

// water-to-light map:
// 88 18 7
// 18 25 70

// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4`

// 	result := FindMinLocationValueFromRange(input)

// 	if result != 46 {
// 		t.Errorf("Expected 46, got %d", result)
// 	}
// }
