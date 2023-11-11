package search

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/commondatageek/mark/internal/item"
)

const MinimumNGramLength = 3

func Search(items []*item.Item, q string, n int) ([]*item.Item, error) {
	// build and query index
	idx := build(items)
	results := query(idx, q)

	if n == 0 {
		return results, nil
	}
	if n > 0 {
		return results[:min(len(results), n)], nil
	}

	return nil, fmt.Errorf("Search: n must be >= 0 (received %d)", n)
}

type Index map[string][]*item.Item

// ngrams returns all ngrams (possible substrings) of length n in s.
func ngrams(n int, s string) []string {
	count := len(s) - (n - 1)
	grams := make([]string, 0, count)
	for i := 0; i < count; i++ {
		grams = append(grams, s[i:i+n])
	}
	return grams
}

func itemToStrings(i *item.Item) []string {
	wordStrings := make([]string, 0)
	wordStrings = append(wordStrings, stringToWords(i.URL)...)
	wordStrings = append(wordStrings, stringToWords(i.Description)...)
	for _, n := range i.Names {
		wordStrings = append(wordStrings, stringToWords(n)...)
	}
	for _, n := range i.Tags {
		wordStrings = append(wordStrings, stringToWords(n)...)
	}
	return wordStrings
}

func stringToWords(s string) []string {
	wordSep := regexp.MustCompile(`\W+`)
	words := wordSep.Split(s, -1)
	return words
}

func build(items []*item.Item) Index {
	idx := make(Index)
	for _, i := range items {
		var distinctI item.Item = *i
		itemStrings := itemToStrings(&distinctI)
		for _, s := range itemStrings {
			s = strings.ToLower(s)
			for n := MinimumNGramLength; n <= len(s); n++ {
				// for n := MinimumNGramLength; n <= MaximumNGramLength; n++ {
				grams := unique(ngrams(n, s))
				for _, g := range grams {
					idx[g] = append(idx[g], &distinctI)
				}
			}
		}
	}
	return idx
}

func query(idx Index, query string) []*item.Item {
	results := make([]*item.Item, 0)
	words := stringToWords(query)
	for _, s := range words {
		for n := MinimumNGramLength; n <= len(s); n++ {
			// for n := MinimumNGramLength; n <= MaximumNGramLength; n++ {
			grams := unique(ngrams(n, s))
			for _, g := range grams {
				g = strings.ToLower(g)
				if items, ok := idx[g]; ok {
					results = append(results, items...)
				}
			}
		}
	}

	counts := count(results)
	sorted := sortByCount(counts)
	slices.Reverse(sorted)

	return sorted
}

func count(items []*item.Item) map[*item.Item]int {
	counts := make(map[*item.Item]int)
	for _, i := range items {
		var i *item.Item = i
		counts[i]++
	}
	return counts
}

func sortByCount(counts map[*item.Item]int) []*item.Item {
	type ItemCount struct {
		Item  *item.Item
		Count int
	}
	countSlice := make([]ItemCount, 0)
	for k, v := range counts {
		var k *item.Item = k
		var v int = v
		countSlice = append(countSlice, ItemCount{k, v})
	}
	slices.SortFunc(countSlice, func(a, b ItemCount) int {
		if a.Count > b.Count {
			return 1
		}
		if a.Count < b.Count {
			return -1
		}
		if a.Count == b.Count {
			return 0
		}
		panic("sortByCount: a and b aren't comparable")
	})
	result := make([]*item.Item, len(countSlice))
	for i := range countSlice {
		result[i] = countSlice[i].Item
	}
	return result
}

func unique[T comparable](items []T) []T {
	unique := make(map[T]bool)
	for _, item := range items {
		unique[item] = true
	}
	results := make([]T, 0)
	for k := range unique {
		results = append(results, k)
	}
	return results
}
