package bookmarkstore

import (
	"encoding/json"
	"fmt"
	"io"

	bmark "github.com/commondatageek/mark/internal/bookmark"
)

type BookmarkStore struct {
	names map[string]*bmark.Bookmark
}

func Load(r io.Reader) (*BookmarkStore, error) {
	s := New()

	dec := json.NewDecoder(r)
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
