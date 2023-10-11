package bookmarkstore

import (
	"strings"
	"testing"

	"github.com/commondatageek/mark/internal/bookmark"
)

func TestLoad(t *testing.T) {
	t.Skip("not imiplemented")
}

func TestLoadWarnMalformed(t *testing.T) {
	t.Skip("not imiplemented")
}

func TestGetMultipleNamesSingleBookmark(t *testing.T) {
	s := New()
	b := bookmark.Bookmark{
		Names: []string{"a", "b"},
		URL:   "https://example.com",
	}
	s.Add(b)

	if s.Get("a") == nil || s.Get("b") == nil {
		t.Fatal("could not find each of multiple names")
	}
}

func TestAddDuplicateNamesFails(t *testing.T) {
	s := New()
	b1 := bookmark.Bookmark{
		Names: []string{"a", "b"},
		URL:   "https://example.com",
	}
	b2 := bookmark.Bookmark{
		Names: []string{"b", "c"},
		URL:   "https://otherexample.com",
	}

	s.Add(b1)
	err := s.Add(b2)
	if err == nil || !strings.Contains(err.Error(), "already exists") {
		t.Fatal("did not detect naming collision")
	}
}

func TestGet(t *testing.T) {
	t.Skip("not imiplemented")
}
