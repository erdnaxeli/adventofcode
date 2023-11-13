package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day5p1(input aoc.Input) string {
	index := 0
	var password strings.Builder

	for password.Len() < 8 {
		hash := md5.New()
		fmt.Fprintf(hash, "%s%d", input.Content(), index)
		sum := fmt.Sprintf("%x", hash.Sum(nil))
		if sum[:5] == "00000" {
			fmt.Fprint(&password, string(sum[5]))
		}

		index++
	}

	return password.String()
}

func (s solver) Day5p2(input aoc.Input) string {
	index := 0
	found := 0
	password := make([]byte, 8)

	for found < 8 {
		hash := md5.New()
		fmt.Fprintf(hash, "%s%d", input.Content(), index)
		sum := fmt.Sprintf("%x", hash.Sum(nil))
		index++

		if sum[:5] == "00000" {
			p := int(sum[5] - '0')
			log.Printf("%s%d: %s, %s", input.Content(), index, sum, password)

			if p < 0 || p > 7 || password[p] != 0 {
				continue
			}

			password[p] = sum[6]
			found++
		}
	}

	return string(password)
}
