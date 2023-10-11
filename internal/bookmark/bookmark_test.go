package bookmark

import "testing"

func TestStringMatchesFormat(t *testing.T) {
	b := Bookmark{
		Names: []string{"a", "b", "c"},
		URL:   "https://example.com",
	}

	expected := `Bookmark(Names: ["a", "b", "c"], URL: "https://example.com")`
	received := b.String()
	if received != expected {
		t.Fatalf("'%s' != '%s'", received, expected)
	}
}
