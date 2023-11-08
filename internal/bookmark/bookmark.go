package bookmark

import (
	"fmt"
	"strings"
)

type Bookmark struct {
	Names       []string `json:"names"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags,omitempty"`
	Description string   `json:"description,omitempty"`
}

func (b Bookmark) String() string {
	namesString := "\"" + strings.Join(b.Names, "\", \"") + "\""
	tagsString := "\"" + strings.Join(b.Tags, "\", \"") + "\""

	return fmt.Sprintf("names: [%s]\n  tags: [%s]\n  url: %s", namesString, tagsString, b.URL)
}
