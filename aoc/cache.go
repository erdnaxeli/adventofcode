package aoc

type Cache interface {
	GetInput(year int, day int, part int) string
}

type DefaultCache struct{}

func NewDefaultCache() DefaultCache {
	return DefaultCache{}
}

func (c DefaultCache) GetInput(year int, day int, part int) string {
	return ""
}
