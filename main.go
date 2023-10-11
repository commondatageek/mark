package main

import (
	"fmt"
	"os"
	"path"

	store "github.com/commondatageek/mark/internal/bookmarkstore"

	"github.com/pkg/browser"
)

const (
	SUCCESS            = 0
	ERR_USAGE          = 1
	ERR_LOAD_STORE     = 2
	ERR_BROWSER        = 3
	ERR_NAME_NOT_FOUND = 4
	ERR_HOME_DIR       = 5
)

func main() {
	if len(os.Args) != 2 {
		fatal("usage: mark LABEL", ERR_USAGE)
	}

	bookmarks, err := store.Load(bookmarksPath())
	if err != nil {
		fatal(err.Error(), ERR_LOAD_STORE)
	}

	label := os.Args[1]

	if b := bookmarks.Get(label); b != nil {
		fmt.Printf("opening: %s\n", b)
		if err := browser.OpenURL(b.URL); err != nil {
			fatal(err.Error(), ERR_BROWSER)
		}
	} else {
		fatal(fmt.Sprintf("unable to find a URL for '%s'", label), ERR_NAME_NOT_FOUND)
	}

	os.Exit(SUCCESS)
}

func fatal(msg string, code int) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(code)
}

func homeDir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		fatal("could not get user home directory: "+err.Error(), ERR_HOME_DIR)
	}
	return h
}

func bookmarksPath() string {
	return path.Join(homeDir(), ".bookmarks.jsonl")
}
