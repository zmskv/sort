package comparator

import (
	"strconv"
	"strings"

	"github.com/zmskv/sort/internal/options"
)

var monthMap = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
	"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
	"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

var humanSuffixes = map[string]float64{
	"K": 1024,
	"M": 1024 * 1024,
	"G": 1024 * 1024 * 1024,
	"T": 1024 * 1024 * 1024 * 1024,
}

// Comparator - функция сравнения
type Comparator func(a, b string) int

// Make - возвращает функцию сравнения
func Make(opts options.Options) Comparator {
	return func(a, b string) int {
		if opts.IgnoreB {
			a, b = strings.TrimSpace(a), strings.TrimSpace(b)
		}

		switch {
		case opts.Month:
			return compareInts(monthMap[a], monthMap[b])
		case opts.Human:
			return compareFloats(parseHumanSize(a), parseHumanSize(b))
		case opts.Numeric:
			fa, _ := strconv.ParseFloat(a, 64)
			fb, _ := strconv.ParseFloat(b, 64)
			return compareFloats(fa, fb)
		default:
			return strings.Compare(a, b)
		}
	}
}

func compareInts(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func compareFloats(a, b float64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func parseHumanSize(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	last := strings.ToUpper(string(s[len(s)-1]))
	if mult, ok := humanSuffixes[last]; ok {
		val, err := strconv.ParseFloat(s[:len(s)-1], 64)
		if err == nil {
			return val * mult
		}
	}
	val, _ := strconv.ParseFloat(s, 64)
	return val
}
