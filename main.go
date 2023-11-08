package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	bmark "github.com/commondatageek/mark/internal/bookmark"
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

	RFC3339 = "2006-01-02T15:04:05Z07:00"
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
		switch os.Args[1] {
		case "go":
			if len(os.Args) != 3 {
				usage()
				os.Exit(1)
			}
			url := os.Args[2]
			err := open(bookmarks, url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}

			if err := save(bookmarks); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}
		case "add":
			if len(os.Args) != 3 {
				usage()
				os.Exit(1)
			}
			url := os.Args[2]
			err := add(bookmarks, url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}

			if err := save(bookmarks); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}
		default:
			usage()
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func save(bookmarks *store.BookmarkStore) error {
	f, err := os.Create(bookmarksPath())
	if err != nil {
		return fmt.Errorf("save: %s", err)
	}
	defer f.Close()

	if err := bookmarks.Save(f); err != nil {
		return fmt.Errorf("save: %s", err)
	}

	return nil
}

func add(bookmarks *store.BookmarkStore, url string) error {
	var getCommaSeparatedList = func(prompt string) ([]string, error) {
		inputs, err := getInput(prompt)
		if err != nil {
			return nil, fmt.Errorf("getCommaSeparatedList: %s", err)
		}
		inputsList := strings.Split(inputs, ",")
		for i := range inputsList {
			inputsList[i] = strings.TrimSpace(inputsList[i])
		}
		return inputsList, nil
	}

	names, err := getCommaSeparatedList("names")
	if err != nil {
		return fmt.Errorf("add: %s", err)
	}

	tags, err := getCommaSeparatedList("tags")
	if err != nil {
		return fmt.Errorf("add: %s", err)
	}

	timestamp := time.Now().UTC()

	b := bmark.Bookmark{
		Names:        names,
		Tags:         tags,
		URL:          url,
		CreatedTime:  timestamp.Format(RFC3339),
		ModifiedTime: timestamp.Format(RFC3339),
	}
	err = bookmarks.Add(b)
	if err != nil {
		return fmt.Errorf("add: %s", err)
	}

	return nil
}

func open(bookmarks *store.BookmarkStore, label string) error {
	if b := bookmarks.Get(label); b != nil {
		fmt.Printf("opening: %s\n", b)
		if err := browser.OpenURL(b.URL); err != nil {
			return fmt.Errorf("go: %s", err)
		}
		b.AccessedTime = time.Now().UTC().Format(RFC3339)
		b.AccessCount += 1
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
	label, err := getInput("query")
	if err != nil {
		return fmt.Errorf("search: %s", err)
	}

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

func getInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	result, err := reader.ReadString('\n')
	if err != nil {
		return result, fmt.Errorf("getInput: %s", err)
	}
	return result, nil
}
