package bookmarkstore

import (
	"encoding/json"
	"fmt"
	"io"

	bmark "github.com/commondatageek/mark/internal/bookmark"
	search "github.com/commondatageek/mark/internal/bookmarkstore/search"
)

type BookmarkStore struct {
	names map[string]*bmark.Bookmark
}

func Load(r io.Reader) (*BookmarkStore, error) {
	s := New()

	dec := json.NewDecoder(r)
	// TODO: can we just use a `for dec.More()` here instead of checking for io.EOF?
	for {
		var b bmark.Bookmark
		if err := dec.Decode(&b); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, fmt.Errorf("boookmarkstore.Load: cannot decode JSON: %s", err)
			}
		}
		s.Add(b)
	}

	return s, nil
}

func (s *BookmarkStore) Save(w io.Writer) error {
	enc := json.NewEncoder(w)
	for _, b := range s.All() {
		err := enc.Encode(b)
		if err != nil {
			return fmt.Errorf("Save: %s", err)
		}
	}
	return nil
}

func New() *BookmarkStore {
	return &BookmarkStore{
		names: make(map[string]*bmark.Bookmark),
	}
}

func (s *BookmarkStore) Add(b bmark.Bookmark) error {
	// see if any of the names are already taken
	for _, n := range b.Names {
		if b, found := s.names[n]; found {
			return fmt.Errorf("name '%s' already exists: %s", n, b)
		}
	}

	// doesn't exist yet, let's add it
	for _, n := range b.Names {
		s.names[n] = &b
	}

	return nil
}

func (s *BookmarkStore) Get(name string) *bmark.Bookmark {
	b, foundName := s.names[name]
	if !foundName {
		return nil
	}
	return b
}

func (s *BookmarkStore) All() []*bmark.Bookmark {
	set := make(map[*bmark.Bookmark]bool)
	for _, b := range s.names {
		set[b] = true
	}
	list := make([]*bmark.Bookmark, 0, len(s.names))
	for b := range set {
		list = append(list, b)
	}
	return list
}

func (s *BookmarkStore) Search(query string, n int) ([]*bmark.Bookmark, error) {
	results, err := search.Search(s.All(), query, n)
	if err != nil {
		return nil, fmt.Errorf("Search: %s", err)
	}
	return results, nil
}
