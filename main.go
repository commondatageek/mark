package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/browser"
)

func main() {
	bookmarks := getBookmarks()

	if len(os.Args) != 2 {
		log.Fatalf("usage: mark (name|alias)")
	}

	label := os.Args[1]

	if b := bookmarks.Get(label); b != nil {
		log.Printf("opening %s ...", b.Names[0])
		if err := browser.OpenURL(b.URL); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Could not find a URL for '%s'", label)
	}
}

func getBookmarks() *BookmarkStore {
	s := NewBookmarkStore()

	return s
}

func NewBookmarkStore() *BookmarkStore {
	return &BookmarkStore{
		names: make(map[string]*Bookmark),
	}
}

type Bookmark struct {
	Names []string
	URL   string
}

func (b Bookmark) String() string {
	return fmt.Sprintf("%s|%s", strings.Join(b.Names, ","), b.URL)
}

type BookmarkStore struct {
	names map[string]*Bookmark
}

func (s *BookmarkStore) Add(b Bookmark) error {
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

func (s *BookmarkStore) Get(name string) *Bookmark {
	b, foundName := s.names[name]
	if !foundName {
		return nil
	}
	return b
}
