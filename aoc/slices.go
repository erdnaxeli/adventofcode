package aoc

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
