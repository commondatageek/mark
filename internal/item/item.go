package item

import (
	"fmt"
	"strings"
)

type ItemV1 struct {
	Names        []string `json:"names"`
	URL          string   `json:"url"`
	Tags         []string `json:"tags,omitempty"`
	Description  string   `json:"description,omitempty"`
	CreatedTime  string   `json:"created_time,omitempty"`
	ModifiedTime string   `json:"modified_time,omitempty"`
	AccessedTime string   `json:"accessed_time,omitempty"`
	AccessCount  int      `json:"access_count,omitempty"`
}

func (i ItemV1) String() string {
	namesString := "\"" + strings.Join(i.Names, "\", \"") + "\""
	tagsString := "\"" + strings.Join(i.Tags, "\", \"") + "\""

	return fmt.Sprintf("names: [%s]\n  tags: [%s]\n  url: %s", namesString, tagsString, i.URL)
}

type ItemV2 struct {
	Version      int      `json:"version"`
	Path         string   `json:"path"`
	Title        string   `json:"title"`
	URL          string   `json:"url"`
	Aliases      []string `json:"aliases,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Description  string   `json:"description,omitempty"`
	CreatedTime  string   `json:"created_time,omitempty"`
	ModifiedTime string   `json:"modified_time,omitempty"`
	AccessedTime string   `json:"accessed_time,omitempty"`
	AccessCount  int      `json:"access_count,omitempty"`
}
