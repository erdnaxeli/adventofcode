package main

import (
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type CCCard int

const (
	C2 CCCard = iota
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	T
	J
	Q
	K
	A
)

type CCType int

const (
	HighCard CCType = iota
	Pair1
	Pair2
	Three
	FullHouse
	Four
	Five
)

type CCHand struct {
	Hand []CCCard
	Type CCType
	Bid  int
}

func (s solver) Day7p1(input aoc.Input) string {
	var hands []CCHand
	for _, line := range input.ToStringSlice() {
		parts := line.Split()
		hand := parseCCCards(parts[0].SplitOnS(""))
		hands = append(hands, CCHand{
			Hand: hand,
			Type: getCCHandType(hand),
			Bid:  parts[1].Atoi(),
		})
	}

	slices.SortFunc(hands, func(a, b CCHand) int {
		if a.Type < b.Type {
			return -1
		} else if a.Type > b.Type {
			return 1
		} else {
			for i := range a.Hand {
				if a.Hand[i] < b.Hand[i] {
					return -1
				} else if a.Hand[i] > b.Hand[i] {
					return 1
				}
			}

			return 0
		}
	})

	wh := 0
	for i, hand := range hands {
		wh += (i + 1) * hand.Bid
	}

	return aoc.ResultI(wh)
}

func (s solver) Day7p2(input aoc.Input) string {
	var hands []CCHand
	for _, line := range input.ToStringSlice() {
		parts := line.Split()
		hand := parseCCCards(parts[0].SplitOnS(""))
		hands = append(hands, CCHand{
			Hand: hand,
			Type: getCCHandType2(hand),
			Bid:  parts[1].Atoi(),
		})
	}

	slices.SortFunc(hands, func(a, b CCHand) int {
		if a.Type < b.Type {
			return -1
		} else if a.Type > b.Type {
			return 1
		} else {
			for i := range a.Hand {
				if a.Hand[i] == b.Hand[i] {
					continue
				}

				if a.Hand[i] == J {
					return -1
				} else if b.Hand[i] == J {
					return 1
				} else if a.Hand[i] < b.Hand[i] {
					return -1
				} else if a.Hand[i] > b.Hand[i] {
					return 1
				}
			}

			return 0
		}
	})

	wh := 0
	for i, hand := range hands {
		wh += (i + 1) * hand.Bid
	}

	return aoc.ResultI(wh)
}

func parseCCCards(c []string) []CCCard {
	var cards []CCCard
	for _, cc := range c {
		cards = append(cards, parseCCCard(cc))
	}

	return cards
}

func parseCCCard(s string) CCCard {
	switch s {
	case "A":
		return A
	case "K":
		return K
	case "Q":
		return Q
	case "J":
		return J
	case "T":
		return T
	case "9":
		return C9
	case "8":
		return C8
	case "7":
		return C7
	case "6":
		return C6
	case "5":
		return C5
	case "4":
		return C4
	case "3":
		return C3
	case "2":
		return C2
	default:
		panic("unknown card " + s)
	}
}

func getCCHandType(cards []CCCard) CCType {
	c := make(map[CCCard]int)
	for _, card := range cards {
		c[card]++
	}

	three := false
	pairs := 0
	for _, count := range c {
		if count == 5 {
			return Five
		}

		if count == 4 {
			return Four
		}

		if count == 3 {
			three = true
		} else if count == 2 {
			pairs += 1
		}
	}

	if three {
		if pairs == 1 {
			return FullHouse
		}

		return Three
	}

	if pairs == 2 {
		return Pair2
	}

	if pairs == 1 {
		return Pair1
	}

	return HighCard
}

func getCCHandType2(cards []CCCard) CCType {
	// we now use the joker
	jokers := 0
	c := make(map[CCCard]int)
	for _, card := range cards {
		if card == J {
			jokers += 1
			continue
		}

		c[card]++
	}

	three := false
	pairs := 0
	for _, count := range c {
		if count == 5 {
			return Five
		}

		if count == 4 {
			if jokers == 1 {
				return Five
			}

			return Four
		}

		if count == 3 {
			three = true
		} else if count == 2 {
			pairs += 1
		}
	}

	if three {
		if jokers == 2 {
			return Five
		} else if jokers == 1 {
			return Four
		}

		if pairs == 1 {
			return FullHouse
		}

		return Three
	}

	if pairs == 2 {
		if jokers == 1 {
			return FullHouse
		}

		return Pair2
	}

	if pairs == 1 {
		if jokers == 3 {
			return Five
		} else if jokers == 2 {
			return Four
		} else if jokers == 1 {
			return Three
		}

		return Pair1
	}

	if jokers == 5 || jokers == 4 {
		return Five
	} else if jokers == 3 {
		return Four
	} else if jokers == 2 {
		return Three
	} else if jokers == 1 {
		return Pair1
	}

	return HighCard
}
