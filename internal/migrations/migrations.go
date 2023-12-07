package migrations

import (
	"strings"

	"github.com/commondatageek/mark/internal/item"
)

func MigrateV1toV2(i item.ItemV1) item.ItemV2 {
	path := i.Names[0]
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = strings.ReplaceAll(path, ":", "-")

	var aliases []string
	if len(i.Names) > 1 {
		aliases = i.Names[1:]
	} else {
		aliases = nil
	}

	return item.ItemV2{
		Version:      2,
		Path:         path,
		Title:        i.Names[0],
		URL:          i.URL,
		Aliases:      aliases,
		Tags:         i.Tags,
		Description:  i.Description,
		CreatedTime:  i.CreatedTime,
		ModifiedTime: i.ModifiedTime,
		AccessedTime: i.AccessedTime,
		AccessCount:  i.AccessCount,
	}
}
