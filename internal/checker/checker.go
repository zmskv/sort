package checker

import (
	"github.com/zmskv/sort/internal/comparator"
	"github.com/zmskv/sort/internal/options"
	"github.com/zmskv/sort/internal/sorter"
)

// CheckSorted проверяет, что данные отсортированы
func CheckSorted(lines []sorter.LineData, opts options.Options) bool {
	cmp := comparator.Make(opts)
	for i := 1; i < len(lines); i++ {
		res := cmp(lines[i-1].Key, lines[i].Key)
		if opts.Reverse {
			if res < 0 {
				return false
			}
		} else {
			if res > 0 {
				return false
			}
		}
	}
	return true
}
