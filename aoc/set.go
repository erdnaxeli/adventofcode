package aoc

import "iter"

type Set[E comparable] map[E]struct{}

func NewSet[E comparable](elems ...E) Set[E] {
	s := make(map[E]struct{})
	for _, e := range elems {
		s[e] = struct{}{}
	}

	return s
}

func NewSetFromIter[E comparable](elems iter.Seq[E]) Set[E] {
	s := make(map[E]struct{})
	for e := range elems {
		s[e] = struct{}{}
	}

	return s
}

func (s Set[E]) Add(e E) {
	s[e] = struct{}{}
}

func (s Set[E]) Contains(e E) bool {
	_, ok := s[e]
	return ok
}

// First actually returns an arbitrary element of the set.
func (s Set[E]) First() E {
	for e := range s {
		return e
	}

	// if the set is empty, we return the zero value for the type E
	var e E
	return e
}

func (s Set[E]) Len() int {
	return len(s)
}

func (s Set[E]) Values() iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range s {
			if !yield(e) {
				return
			}
		}
	}
}
