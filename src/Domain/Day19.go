package Domain

import (
	"sort"
)

type D19Color rune

/* const (
	D19_WHITE D19Color = 'w'
	D19_BLUE  D19Color = 'u'
	D19_BLACK D19Color = 'b'
	D19_RED   D19Color = 'r'
	D19_GREEN D19Color = 'g'
) */

/* type D19Towel struct {
	Colors string
} */

type D19Store struct {
	MaxTowelLen int
	Towels      map[string]bool
	Patterns    []string
	CacheValid  map[string]bool
}

func (d *D19Store) findMaxTowel() {
	if d.MaxTowelLen > 0 {
		return
	}

	for towel, _ := range d.Towels {
		if d.MaxTowelLen < len(towel) {
			d.MaxTowelLen = len(towel)
		}
	}
}

func (d *D19Store) recurse(input string, visited map[string]int) int {
	if options, found := visited[input]; found {
		if options > 0 {
			return options
		}

		return 0
	}

	visited[input] = 0

	maxSearch := len(input)
	if maxSearch > d.MaxTowelLen {
		maxSearch = d.MaxTowelLen
	}
	for i := maxSearch; i > 0; i-- {
		if _, ok := d.Towels[input[:i]]; ok {
			if i == len(input) {
				visited[input]++
			} else {
				visited[input] += d.recurse(input[i:], visited)
			}
		}
	}

	return visited[input]
}

func (d *D19Store) FindPatternRecurse(input string) int {
	d.findMaxTowel()

	visited := make(map[string]int)
	return d.recurse(input, visited)
}

func (d *D19Store) FindPattern(input string) bool {
	d.findMaxTowel()

	visited := make(map[string]bool)
	searchColorPatterns := []string{input}
	for len(searchColorPatterns) > 0 {
		colorPattern := searchColorPatterns[0]
		searchColorPatterns = searchColorPatterns[1:]

		if visited[colorPattern] {
			continue
		}
		visited[colorPattern] = true

		maxSearch := len(colorPattern)
		if maxSearch > d.MaxTowelLen {
			maxSearch = d.MaxTowelLen
		}
		for i := maxSearch; i > 0; i-- {
			if _, ok := d.Towels[colorPattern[:i]]; ok {
				if i == len(colorPattern) {
					return true
				} else {
					searchColorPatterns = append(searchColorPatterns, colorPattern[i:])
				}
			}
		}

		sort.Slice(searchColorPatterns, func(i, j int) bool {
			return len(searchColorPatterns[i]) > len(searchColorPatterns[j])
		})
	}

	return false
}
