package main

import (
	"log"
	"regexp"

	"github.com/erdnaxeli/adventofcode/aoc"
)

var (
	IPV7_HYPERNET_RGX = regexp.MustCompile(`\[([a-z]+)\]`)
	IPV7_SUPERNET_RGX = regexp.MustCompile(`([a-z]+)(?:\[[a-z]+\])?`)
)

func (s solver) Day7p1(input aoc.Input) string {
	count := 0

	for _, ip := range input.ToStringSlice() {
		log.Printf("Test ip %s", ip)
		hypernetAbba := false
		for _, match := range IPV7_HYPERNET_RGX.FindAllStringSubmatch(string(ip), -1) {
			log.Printf("Test hypernet %s", match[1])
			if ipv7IsABBA(match[1]) {
				hypernetAbba = true
				break
			}
		}

		if hypernetAbba {
			continue
		}

		for _, match := range IPV7_SUPERNET_RGX.FindAllStringSubmatch(string(ip), -1) {
			log.Printf("Test net %s", match[1])
			if ipv7IsABBA(match[1]) {
				count++
				break
			}
		}
	}

	return aoc.ResultI(count)
}

func (s solver) Day7p2(input aoc.Input) string {
	count := 0

	for _, ip := range input.ToStringSlice() {
		hasSSL := false
		var abas []string
		for _, match := range IPV7_SUPERNET_RGX.FindAllStringSubmatch(string(ip), -1) {
			abas = append(abas, ipv7GetABAs(match[1])...)
		}

		for _, match := range IPV7_HYPERNET_RGX.FindAllStringSubmatch(string(ip), -1) {
			supernet := match[1]
			for i := 0; i <= len(supernet)-3; i++ {
				for _, aba := range abas {
					if supernet[i] == aba[1] && supernet[i+1] == aba[0] && supernet[i] == supernet[i+2] {
						hasSSL = true
						break
					}
				}

				if hasSSL {
					break
				}
			}

			if hasSSL {
				break
			}
		}

		if hasSSL {
			count++
		}
	}

	return aoc.ResultI(count)
}

func ipv7IsABBA(ip string) bool {
	for i := 0; i <= len(ip)-4; i++ {
		if ip[i] != ip[i+1] && ip[i] == ip[i+3] && ip[i+1] == ip[i+2] {
			return true
		}
	}

	return false
}

func ipv7GetABAs(supernet string) []string {
	var abas []string
	for i := 0; i <= len(supernet)-3; i++ {
		if supernet[i] != supernet[i+1] && supernet[i] == supernet[i+2] {
			abas = append(abas, supernet[i:i+3])
		}
	}

	return abas
}
