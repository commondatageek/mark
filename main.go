package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"time"

	"github.com/commondatageek/mark/internal/item"
	"github.com/commondatageek/mark/internal/itemstore"

	"github.com/pkg/browser"
)

const (
	SUCCESS            = 0
	ERR_USAGE          = 1
	ERR_ITEMS_FILE     = 2
	ERR_LOAD_STORE     = 3
	ERR_BROWSER        = 4
	ERR_NAME_NOT_FOUND = 5
	ERR_HOME_DIR       = 6
	ERR_OTHER          = 7

	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

func main() {
	path := itemsPath()

	var items *itemstore.ItemStore

	if f, err := os.Open(path); err != nil {
		// if there is no items file, create an empty ItemStore
		if errors.Is(err, fs.ErrNotExist) {
			items = itemstore.New()
		} else {
			fatal(fmt.Sprintf("could not open items file %s: %s", path, err), ERR_ITEMS_FILE)
		}
	} else {
		// if there is an items file, read it
		defer f.Close()
		items, err = itemstore.Load(f)
		if err != nil {
			fatal(err.Error(), ERR_LOAD_STORE)
		}
	}

	if len(os.Args) == 1 {
		// search by default
		err := search(items)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(ERR_OTHER)
		}
	} else {
		switch os.Args[1] {
		case "go":
			if len(os.Args) != 3 {
				usage()
				os.Exit(ERR_USAGE)
			}
			url := os.Args[2]
			err := open(items, url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(ERR_OTHER)
			}

			if err := save(items); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(ERR_OTHER)
			}
		case "add":
			if len(os.Args) != 3 {
				usage()
				os.Exit(ERR_USAGE)
			}
			url := os.Args[2]
			err := add(items, url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(ERR_OTHER)
			}

			if err := save(items); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(ERR_OTHER)
			}
		default:
			usage()
			os.Exit(ERR_USAGE)
		}
	}

	os.Exit(SUCCESS)
}

func save(items *itemstore.ItemStore) error {
	f, err := os.Create(itemsPath())
	if err != nil {
		return fmt.Errorf("save: %s", err)
	}
	defer f.Close()

	if err := items.Save(f); err != nil {
		return fmt.Errorf("save: %s", err)
	}

	return nil
}

func add(items *itemstore.ItemStore, url string) error {
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

	b := item.Item{
		Names:        names,
		Tags:         tags,
		URL:          url,
		CreatedTime:  timestamp.Format(RFC3339),
		ModifiedTime: timestamp.Format(RFC3339),
	}
	err = items.Add(b)
	if err != nil {
		return fmt.Errorf("add: %s", err)
	}

	return nil
}

func open(items *itemstore.ItemStore, label string) error {
	if b := items.Get(label); b != nil {
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
	fmt.Fprintf(os.Stderr, "search for an item:\n")
	fmt.Fprintf(os.Stderr, "\t%s\n", cmd)
	fmt.Fprintf(os.Stderr, "go to an item:\n")
	fmt.Fprintf(os.Stderr, "\t%s go my/item/name\n", cmd)
}

func search(items *itemstore.ItemStore) error {
	label, err := getInput("query")
	if err != nil {
		return fmt.Errorf("search: %s", err)
	}

	results, err := items.Search(label, 5)
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

func itemsPath() string {
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
