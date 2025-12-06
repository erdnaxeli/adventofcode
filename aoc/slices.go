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

func SumSlice(s []int) int {
	result := 0

	for _, v := range s {
		result += v
	}

	return result
}

func MultSlice(s []int) int {
	result := 1

	for _, v := range s {
		result *= v
	}

	return result
}
