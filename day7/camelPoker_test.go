package main

import (
	"testing"
)

func TestGetHandRank(t *testing.T) {
	pokerHands := map[string]int{
		"AAAAA": 6,
		"AAAA2": 5,
		"AAA22": 4,
		"AAA12": 3,
		"AA122": 2,
		"A1233": 1,
		"12345": 0,
	}

	for hand, expectedRank := range pokerHands {
		result := GetPokerHandStrength(hand)

		if result != expectedRank {
			t.Errorf("Expected %d, got %d", expectedRank, result)
		}
	}
}

func TestGetHandValue(t *testing.T) {
	pokerHandsInput := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	result := GetPokerHandWinnings(pokerHandsInput)

	if result != 6440 {
		t.Errorf("Expected 6440, got %d", result)
	}
}

func TestGetHandValueWithJokers(t *testing.T) {
	pokerHandsInput := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	result := GetPokerHandWinningsWithJokers(pokerHandsInput)

	if result != 5905 {
		t.Errorf("Expected 5905, got %d", result)
	}
}

func TestGetHandValueWithJokersEdgeCases(t *testing.T) {
	pokerHandsInput := `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41`

	result := GetPokerHandWinningsWithJokers(pokerHandsInput)

	if result != 6839 {
		t.Errorf("Expected 6839, got %d", result)
	}
}
