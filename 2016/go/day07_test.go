package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpv7IsAbba(t *testing.T) {
	tests := []struct {
		test     string
		expected bool
	}{
		{"abba", true},
		{"abab", false},
		{"aaaabbaaa", true},
		{"aabaaaaba", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, ipv7IsABBA(test.test))
	}
}
