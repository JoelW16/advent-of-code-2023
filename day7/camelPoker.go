package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type pokerHand struct {
	hand     string
	bet      int
	strength int
}

type card struct {
	cardType string
	count    int
}

var cardStrengthLookup = map[string]int{
	"5":         6,
	"1,4":       5,
	"2,3":       4,
	"1,1,3":     3,
	"1,2,2":     2,
	"1,1,1,2":   1,
	"1,1,1,1,1": 0,
}

var cardValue = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var cardValueWithJokers = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

func GetPokerHandStrength(pokerHand string) int {
	cardCounts := make(map[string]int)
	for _, card := range pokerHand {
		cardCounts[string(card)]++
	}
	values := make([]int, 0, len(cardCounts))
	for index := range cardCounts {
		values = append(values, cardCounts[index])
	}
	slices.Sort(values)
	valueString := strings.Trim(strings.Replace(fmt.Sprint(values), " ", ",", -1), "[]")
	return cardStrengthLookup[valueString]
}

func SortKeysByHandCount(cardCounts map[string]int) []string {
	keys := make([]string, 0, len(cardCounts))
	for key := range cardCounts {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return cardCounts[keys[i]] > cardCounts[keys[j]]
	})

	return keys
}

func FindJokerSubstitution(cardCounts map[string]int, hand string) string {
	numberOfJokers := cardCounts["J"]
	sortedKeys := SortKeysByHandCount(cardCounts)
	for i := 0; i < len(sortedKeys); i++ {
		if cardCounts[sortedKeys[i]] <= 5-numberOfJokers {
			if sortedKeys[i] == "J" && cardCounts["J"] <= 2 {
				continue
			}
			return sortedKeys[i]
		}
	}
	return "J"
}

func GetPokerHandStrengthWithJokers(hand pokerHand) int {
	cardCounts := make(map[string]int)
	for _, card := range hand.hand {
		cardCounts[string(card)]++
	}
	if cardCounts["J"] == 0 {
		return GetPokerHandStrength(hand.hand)
	}
	substitution := FindJokerSubstitution(cardCounts, hand.hand)
	newHand := strings.Replace(hand.hand, "J", substitution, -1)
	return GetPokerHandStrength(newHand)
}

func GetHands(input string, hasJokers bool) []pokerHand {
	hands := strings.Split(input, "\n")
	rank := make([]pokerHand, len(hands))
	for i := 0; i < len(hands); i++ {
		hand := strings.Split(hands[i], " ")
		bet, _ := strconv.Atoi(hand[1])

		rank[i].hand = hand[0]
		rank[i].bet = bet
		rank[i].strength = GetPokerHandStrength(rank[i].hand)
		if hasJokers {
			rank[i].strength = GetPokerHandStrengthWithJokers(rank[i])
		}
	}
	return rank
}

func SortHands(rank []pokerHand, hasJokers bool) {
	deck := cardValue
	if hasJokers {
		deck = cardValueWithJokers
	}
	sort.Slice(rank, func(i, j int) bool {
		if rank[i].strength == rank[j].strength {
			for index := 0; index < len(rank[i].hand); index++ {
				if deck[string(rank[i].hand[index])] != deck[string(rank[j].hand[index])] {
					return deck[string(rank[i].hand[index])] < deck[string(rank[j].hand[index])]
				}
			}
		}
		return rank[i].strength < rank[j].strength
	})
}

func CalculateWinnings(rank []pokerHand) int {
	winnings := 0
	for i := 1; i <= len(rank); i++ {
		winnings += rank[i-1].bet * i
	}

	return winnings
}

func GetPokerHandWinnings(input string) int {
	rank := GetHands(input, false)

	SortHands(rank, false)

	return CalculateWinnings(rank)
}

func GetPokerHandWinningsWithJokers(input string) int {
	rank := GetHands(input, true)

	SortHands(rank, true)

	return CalculateWinnings(rank)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := GetPokerHandWinnings(string(content))
	fmt.Println("part 1:", part1)

	part2 := GetPokerHandWinningsWithJokers(string(content))
	fmt.Println("part 2:", part2)

}
