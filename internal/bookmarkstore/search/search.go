package search

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	b "github.com/commondatageek/mark/internal/bookmark"
)

const MinimumNGramLength = 3

func Search(items []*b.Bookmark, q string, n int) ([]*b.Bookmark, error) {
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

type Index map[string][]*b.Bookmark

// ngrams returns all ngrams (possible substrings) of length n in s.
func ngrams(n int, s string) []string {
	count := len(s) - (n - 1)
	grams := make([]string, 0, count)
	for i := 0; i < count; i++ {
		grams = append(grams, s[i:i+n])
	}
	return grams
}

func itemToStrings(item *b.Bookmark) []string {
	wordStrings := make([]string, 0)
	wordStrings = append(wordStrings, stringToWords(item.URL)...)
	wordStrings = append(wordStrings, stringToWords(item.Description)...)
	for _, n := range item.Names {
		wordStrings = append(wordStrings, stringToWords(n)...)
	}
	for _, n := range item.Tags {
		wordStrings = append(wordStrings, stringToWords(n)...)
	}
	return wordStrings
}

func stringToWords(s string) []string {
	wordSep := regexp.MustCompile(`\W+`)
	words := wordSep.Split(s, -1)
	return words
}

func build(items []*b.Bookmark) Index {
	idx := make(Index)
	for _, item := range items {
		var i b.Bookmark = *item
		itemStrings := itemToStrings(&i)
		for _, s := range itemStrings {
			s = strings.ToLower(s)
			for n := MinimumNGramLength; n <= len(s); n++ {
				// for n := MinimumNGramLength; n <= MaximumNGramLength; n++ {
				grams := unique(ngrams(n, s))
				for _, g := range grams {
					idx[g] = append(idx[g], &i)
				}
			}
		}
	}
	return idx
}

func query(idx Index, query string) []*b.Bookmark {
	results := make([]*b.Bookmark, 0)
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

func count(items []*b.Bookmark) map[*b.Bookmark]int {
	counts := make(map[*b.Bookmark]int)
	for _, i := range items {
		var i *b.Bookmark = i
		counts[i]++
	}
	return counts
}

func sortByCount(counts map[*b.Bookmark]int) []*b.Bookmark {
	type BookmarkCount struct {
		Item  *b.Bookmark
		Count int
	}
	countSlice := make([]BookmarkCount, 0)
	for k, v := range counts {
		var k *b.Bookmark = k
		var v int = v
		countSlice = append(countSlice, BookmarkCount{k, v})
	}
	slices.SortFunc(countSlice, func(a, b BookmarkCount) int {
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
	result := make([]*b.Bookmark, len(countSlice))
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
