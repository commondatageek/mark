package main

import (
	"log"
	"os"

	store "github.com/commondatageek/mark/internal/bookmarkstore"

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

func getBookmarks() *store.BookmarkStore {
	s := store.New()

	return s
}
