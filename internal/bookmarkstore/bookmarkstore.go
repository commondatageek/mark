package bookmarkstore

import (
	"fmt"

	bmark "github.com/commondatageek/mark/internal/bookmark"
)

type BookmarkStore struct {
	names map[string]*bmark.Bookmark
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
