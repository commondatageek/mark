package itemstore

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/commondatageek/mark/internal/item"
	"github.com/commondatageek/mark/internal/itemstore/search"
)

type ItemStore struct {
	names map[string]*item.ItemV1
}

func Load(r io.Reader) (*ItemStore, error) {
	s := New()

	dec := json.NewDecoder(r)
	// TODO: can we just use a `for dec.More()` here instead of checking for io.EOF?
	for {
		var b item.ItemV1
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

func (s *ItemStore) Save(w io.Writer) error {
	enc := json.NewEncoder(w)
	for _, b := range s.All() {
		err := enc.Encode(b)
		if err != nil {
			return fmt.Errorf("Save: %s", err)
		}
	}
	return nil
}

func New() *ItemStore {
	return &ItemStore{
		names: make(map[string]*item.ItemV1),
	}
}

func (s *ItemStore) Add(b item.ItemV1) error {
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

func (s *ItemStore) Get(name string) *item.ItemV1 {
	b, foundName := s.names[name]
	if !foundName {
		return nil
	}
	return b
}

func (s *ItemStore) All() []*item.ItemV1 {
	set := make(map[*item.ItemV1]bool)
	for _, b := range s.names {
		set[b] = true
	}
	list := make([]*item.ItemV1, 0, len(s.names))
	for b := range set {
		list = append(list, b)
	}
	return list
}

func (s *ItemStore) Search(query string, n int) ([]*item.ItemV1, error) {
	results, err := search.Search(s.All(), query, n)
	if err != nil {
		return nil, fmt.Errorf("Search: %s", err)
	}
	return results, nil
}
