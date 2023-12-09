package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNextOasisValue(t *testing.T) {
	sequence := []int{0, 3, 6, 9, 12, 15}

	result := getNextOasisValue(sequence)

	assert.Equal(t, 18, result)
}

func TestGetPrevOasisValue(t *testing.T) {
	sequence := []int{10, 13, 16, 21, 30, 45}

	result := getPrevOasisValue(sequence)

	assert.Equal(t, 5, result)
}
