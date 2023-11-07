package main

import (
	"bufio"
	"fmt"
	"os"
	"path"

	store "github.com/commondatageek/mark/internal/bookmarkstore"

	"github.com/pkg/browser"
)

const (
	SUCCESS            = 0
	ERR_USAGE          = 1
	ERR_BOOKMARKS_FILE = 2
	ERR_LOAD_STORE     = 3
	ERR_BROWSER        = 4
	ERR_NAME_NOT_FOUND = 5
	ERR_HOME_DIR       = 6
)

func main() {
	path := bookmarksPath()
	f, err := os.Open(path)
	if err != nil {
		fatal(fmt.Sprintf("could not open bookmarks file %s: %s", path, err), ERR_BOOKMARKS_FILE)
	}
	defer f.Close()

	bookmarks, err := store.Load(f)
	if err != nil {
		fatal(err.Error(), ERR_LOAD_STORE)
	}

	if len(os.Args) == 1 {
		// search by default
		err := search(bookmarks)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	} else {
		if os.Args[1] == "go" {
			err := open(bookmarks, os.Args[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}
		} else {
			usage()
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func open(bookmarks *store.BookmarkStore, label string) error {
	if b := bookmarks.Get(label); b != nil {
		fmt.Printf("opening: %s\n", b)
		if err := browser.OpenURL(b.URL); err != nil {
			return fmt.Errorf("go: %s", err)
		}
	} else {
		return fmt.Errorf("unable to find a URL for '%s'", label)
	}
	return nil
}

func fatal(msg string, code int) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(code)
}

func usage() {
	cmd := os.Args[0]
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "search for a bookmark:\n")
	fmt.Fprintf(os.Stderr, "\t%s\n", cmd)
	fmt.Fprintf(os.Stderr, "go to a bookmark:\n")
	fmt.Fprintf(os.Stderr, "\t%s my/bookmark/name\n", cmd)
}

func search(bookmarks *store.BookmarkStore) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("query: ")
	label, _ := reader.ReadString('\n')

	results, err := bookmarks.Search(label, 5)
	if err != nil {
		return fmt.Errorf("search: %s", err)
	}

	for _, r := range results {
		fmt.Println(r)
	}

	return nil
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
