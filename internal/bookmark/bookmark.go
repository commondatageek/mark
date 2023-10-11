package bookmark

import (
	"fmt"
	"strings"
)

type Bookmark struct {
	Names []string `json:"names"`
	URL   string   `json:"url"`
}

func (b Bookmark) String() string {
	nameStrs := make([]string, len(b.Names))
	for i := 0; i < len(nameStrs); i++ {
		nameStrs[i] = fmt.Sprintf(`"%s"`, b.Names[i])
	}
	namesString := strings.Join(nameStrs, ", ")

	return fmt.Sprintf(`Bookmark(Names: [%s], URL: "%s")`, namesString, b.URL)
}
