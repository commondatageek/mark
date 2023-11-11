package itemstore

import (
	"strings"
	"testing"

	"github.com/commondatageek/mark/internal/item"
)

func TestLoad(t *testing.T) {
	t.Skip("not imiplemented")
}

func TestLoadWarnMalformed(t *testing.T) {
	t.Skip("not imiplemented")
}

func TestGetMultipleNamesSingleItem(t *testing.T) {
	s := New()
	b := item.Item{
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
	b1 := item.Item{
		Names: []string{"a", "b"},
		URL:   "https://example.com",
	}
	b2 := item.Item{
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
