package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day4p1(input aoc.Input) string {
	sum := 0
	for _, line := range input.ToStringSlice() {
		room := NewRoom(string(line))
		if room.IsValid() {
			sum += room.SectorID
		}
	}

	return aoc.ResultI(sum)
}

func (s solver) Day4p2(input aoc.Input) string {
	for _, line := range input.ToStringSlice() {
		room := NewRoom(string(line))
		log.Print(room.DecipherName())
		if room.DecipherName() == "northpole object storage" {
			return aoc.ResultI(room.SectorID)
		}
	}

	panic("not found")
}

type Room struct {
	Name     string
	SectorID int
	Checksum string
}

var RoomRgx = regexp.MustCompile(`((?:[a-z]+-)+)(\d+)\[([a-z]+)\]`)

func NewRoom(desc string) Room {
	match := RoomRgx.FindStringSubmatch(desc)
	if match == nil {
		return Room{}
	}

	return Room{
		Name:     match[1][:len(match[1])-1],
		SectorID: aoc.Atoi(match[2]),
		Checksum: match[3],
	}
}

func (r Room) IsValid() bool {
	return r.ComputeChecksum() == r.Checksum
}

func (r Room) ComputeChecksum() string {
	count := make(map[rune]int)
	for _, l := range r.Name {
		if l == '-' {
			continue
		}

		count[l]++
	}

	var letters []rune
	for k := range count {
		letters = append(letters, k)
	}
	sort.Slice(letters, func(i, j int) bool {
		if count[letters[i]] == count[letters[j]] {
			return letters[i] < letters[j]
		}

		return count[letters[i]] >= count[letters[j]]
	})

	var checksum strings.Builder
	for i := 0; i < 5; i++ {
		fmt.Fprint(&checksum, string(letters[i]))
	}
	return checksum.String()
}

func (r Room) DecipherName() string {
	shift := r.SectorID % 26
	var name strings.Builder
	for _, l := range r.Name {
		if l == '-' {
			fmt.Fprint(&name, " ")
			continue
		}

		r := (((l + rune(shift)) - 'a') % 26) + 'a'
		fmt.Fprint(&name, string(r))
	}

	return strings.ToLower(name.String())
}
