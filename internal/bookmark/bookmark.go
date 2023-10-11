package bookmark

import (
	"fmt"
	"strings"
)

type Bookmark struct {
	Names []string
	URL   string
}

func (b Bookmark) String() string {
	return fmt.Sprintf("%s|%s", strings.Join(b.Names, ","), b.URL)
}
