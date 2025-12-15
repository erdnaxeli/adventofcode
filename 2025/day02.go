package main

import (
	"fmt"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day2p1(input aoc.Input) string {
	productIDRanges := input.Delimiter(",").ToRanges(true)

	sum := 0
	for _, productIDRange := range productIDRanges {
		sum += getSumInvalidProductID(productIDRange)
	}

	return aoc.ResultI(sum)
}

func (s solver) Day2p2(input aoc.Input) string {
	fmt.Println("This part was solven in perl, run day02.pl")
	return ""
}

func getSumInvalidProductID(productIDRange aoc.Range) int {
	sum := 0
	for productID := productIDRange.Start; productID <= productIDRange.End; productID++ {
		if !isProductIDValid(productID) {
			sum += productID
		}
	}

	return sum
}

func isProductIDValid(productID int) bool {
	s := fmt.Sprint(productID)
	return len(s)%2 != 0 || s[0:len(s)/2] != s[len(s)/2:]
}
