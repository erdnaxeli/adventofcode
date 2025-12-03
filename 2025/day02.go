package main

import (
	"fmt"

	"github.com/erdnaxeli/adventofcode/aoc"
)

type ProductIDRange struct {
	first int
	last  int
}

func (s solver) Day2p1(input aoc.Input) string {
	productIdRanges := parseInputD2(input)

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

func parseInputD2(input aoc.Input) []ProductIDRange {
	var productIdsRanges []ProductIDRange

	productIdRangesS := input.Delimiter(",").ToStringSlice()
	for _, productIdRangeS := range productIdRangesS {
		parts := productIdRangeS.SplitOnAtoi("-")
		productIdsRanges = append(productIdsRanges, ProductIDRange{first: parts[0], last: parts[1]})
	}

	return productIdsRanges
}

func getSumInvalidProductID(productIdRange ProductIDRange) int {
	sum := 0
	for productId := productIdRange.first; productId <= productIdRange.last; productId++ {
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
