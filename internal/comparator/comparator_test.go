package comparator

import (
	"testing"

	"github.com/zmskv/sort/internal/options"
)

func TestComparatorNumeric(t *testing.T) {
	opts := options.Options{Numeric: true}
	cmp := Make(opts)

	tests := []struct {
		a, b string
		want int
	}{
		{"2", "10", -1},
		{"10", "2", 1},
		{"5", "5", 0},
	}

	for _, tt := range tests {
		got := cmp(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Numeric cmp(%q,%q) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestComparatorHuman(t *testing.T) {
	opts := options.Options{Human: true}
	cmp := Make(opts)

	tests := []struct {
		a, b string
		want int
	}{
		{"1K", "2K", -1},
		{"5M", "1M", 1},
		{"10K", "10K", 0},
	}

	for _, tt := range tests {
		got := cmp(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Human cmp(%q,%q) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestComparatorMonth(t *testing.T) {
	opts := options.Options{Month: true}
	cmp := Make(opts)

	tests := []struct {
		a, b string
		want int
	}{
		{"Jan", "Feb", -1},
		{"Dec", "Mar", 1},
		{"Apr", "Apr", 0},
	}

	for _, tt := range tests {
		got := cmp(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Month cmp(%q,%q) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestComparatorIgnoreTrailing(t *testing.T) {
	opts := options.Options{IgnoreB: true}
	cmp := Make(opts)

	tests := []struct {
		a, b string
		want int
	}{
		{"apple   ", "apple", 0},
		{"banana ", "banana", 0},
	}

	for _, tt := range tests {
		got := cmp(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("IgnoreB cmp(%q,%q) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
