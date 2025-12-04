package aoc

// FirstMaxIndex returns the index of the first max value found, and the value.
func FirstMaxIndex(s []int) (int, int) {
	max := s[0]
	maxI := 0
	for i, e := range s {
		if e > max {
			max = e
			maxI = i
		}
	}

	return maxI, max
}

// FirstMaxIndex returns the index of the first min value found, and the value.
func FirstMinIndex(s []int) (int, int) {
	max := s[0]
	maxI := 0
	for i, e := range s {
		if e < max {
			max = e
			maxI = i
		}
	}

	return maxI, max
}
