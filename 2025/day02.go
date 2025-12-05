package main

import (
	"fmt"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day2p1(input aoc.Input) string {
	productIdRanges := input.Delimiter(",").ToRanges(true)

	sum := 0
	for _, productIdRange := range productIdRanges {
		sum += getSumInvalidProductID(productIdRange)
	}

	return aoc.ResultI(sum)
}

func (s solver) Day2p2(input aoc.Input) string {
	fmt.Println("This part was solven in perl, run day02.pl")
	return ""
}

func getSumInvalidProductID(productIdRange aoc.Range) int {
	sum := 0
	for productId := productIdRange.Start; productId <= productIdRange.End; productId++ {
		if !isProductIDValid(productId) {
			sum += productId
		}
	}

	return sum
}

func isProductIDValid(productId int) bool {
	s := fmt.Sprint(productId)
	return len(s)%2 != 0 || s[0:len(s)/2] != s[len(s)/2:]
}
