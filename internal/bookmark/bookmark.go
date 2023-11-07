package bookmark

import (
	"fmt"
	"strings"
)

type Bookmark struct {
	Names       []string `json:"names"`
	Tags        []string `json:"tags"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
}

func (b Bookmark) String() string {
	namesString := "\"" + strings.Join(b.Names, "\", \"") + "\""
	tagsString := "\"" + strings.Join(b.Tags, "\", \"") + "\""

	return fmt.Sprintf("names: [%s]\n  tags: [%s]\n  url: %s", namesString, tagsString, b.URL)
}
