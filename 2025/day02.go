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
	product_id_ranges := parseInputD2(input)

	sum := 0
	for _, product_id_range := range product_id_ranges {
		sum += getSumInvalidProductID(product_id_range)
	}

	return aoc.ResultI(sum)
}

func (s solver) Day2p2(input aoc.Input) string {
	fmt.Println("This part was solven in perl, run day02.pl")
	return ""
}

func parseInputD2(input aoc.Input) []ProductIDRange {
	var product_ids_ranges []ProductIDRange

	product_id_ranges_s := input.Delimiter(",").ToStringSlice()
	for _, product_id_range_s := range product_id_ranges_s {
		parts := product_id_range_s.SplitOnAtoi("-")
		product_ids_ranges = append(product_ids_ranges, ProductIDRange{first: parts[0], last: parts[1]})
	}

	return product_ids_ranges
}

func getSumInvalidProductID(product_id_range ProductIDRange) int {
	sum := 0
	for product_id := product_id_range.first; product_id <= product_id_range.last; product_id++ {
		if !isProductIDValid(product_id) {
			fmt.Printf("%d is invalid\n", product_id)
			sum += product_id
		}
	}

	return sum
}

func isProductIDValid(product_id int) bool {
	s := fmt.Sprint(product_id)
	return len(s)%2 != 0 || s[0:len(s)/2] != s[len(s)/2:]
}
