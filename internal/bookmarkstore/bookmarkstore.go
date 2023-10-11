package bookmarkstore

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	bmark "github.com/commondatageek/mark/internal/bookmark"
)

type BookmarkStore struct {
	names map[string]*bmark.Bookmark
}

func Load(path string) (*BookmarkStore, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("bookmarkstore.Load: %s", err)
	}
	defer f.Close()

	s := New()

	dec := json.NewDecoder(f)
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
