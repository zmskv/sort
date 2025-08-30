package sorter

import (
	"sort"
	"testing"

	"github.com/zmskv/sort/internal/comparator"
	"github.com/zmskv/sort/internal/options"
)

func TestSorterLexicalAndReverse(t *testing.T) {
	lines := []LineData{
		{"paul", "paul"},
		{"bob", "bob"},
		{"andrew", "andrew"},
	}

	cmp := comparator.Make(options.Options{})
	s := New(lines, cmp, false)
	sort.Sort(s)

	want := []string{"andrew", "bob", "paul"}
	for i, l := range s.lines {
		if l.Original != want[i] {
			t.Errorf("Lexical got %q, want %q", l.Original, want[i])
		}
	}

	s = New(lines, cmp, true)
	sort.Sort(s)
	want = []string{"paul", "bob", "andrew"}
	for i, l := range s.lines {
		if l.Original != want[i] {
			t.Errorf("Reverse got %q, want %q", l.Original, want[i])
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	lines := []LineData{
		{"apple", "apple"},
		{"banana", "banana"},
		{"apple", "apple"},
		{"pear", "pear"},
	}
	lines = RemoveDuplicates(lines)
	want := []string{"apple", "banana", "pear"}

	if len(lines) != len(want) {
		t.Fatalf("RemoveDuplicates length: got %d, want %d", len(lines), len(want))
	}

	for i, l := range lines {
		if l.Original != want[i] {
			t.Errorf("RemoveDuplicates[%d] = %q; want %q", i, l.Original, want[i])
		}
	}
}

func TestSorterByColumn(t *testing.T) {
	lines := []LineData{
		{"andrew\t1", "1"},
		{"paul\t10", "10"},
		{"bob\t2", "2"},
	}

	cmp := comparator.Make(options.Options{Numeric: true})
	s := New(lines, cmp, false)
	sort.Sort(s)

	want := []string{"andrew\t1", "bob\t2", "paul\t10"}
	for i, l := range s.lines {
		if l.Original != want[i] {
			t.Errorf("Column sort[%d] = %q; want %q", i, l.Original, want[i])
		}
	}
}
