package aoc

import (
	"sort"
)

type Counter struct {
	counter map[rune]int
}

type CounterKey struct {
	Key   rune
	Count int
}

func NewCounter(s String) Counter {
	c := Counter{make(map[rune]int)}
	for _, letter := range string(s) {
		c.counter[letter]++
	}

	return c
}

func (c Counter) GetKeys() []CounterKey {
	var result []CounterKey
	for k, v := range c.counter {
		result = append(result, CounterKey{k, v})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Count == result[j].Count {
			return result[i].Key < result[j].Key
		}

		return result[i].Count > result[j].Count
	})
	return result
}

func (c Counter) Get(k rune) int {
	return c.counter[k]
}
