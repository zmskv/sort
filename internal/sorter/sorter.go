package sorter

import "github.com/zmskv/sort/internal/comparator"

// LineData хранит строку и ключ
type LineData struct {
	Original string
	Key      string
}

// Sorter - сортировщик
type Sorter struct {
	lines      []LineData
	comparator comparator.Comparator
	reverse    bool
}

// New - конструктор сортировщика
func New(lines []LineData, cmp comparator.Comparator, reverse bool) Sorter {
	return Sorter{
		lines:      lines,
		comparator: cmp,
		reverse:    reverse,
	}
}

func (s Sorter) Len() int      { return len(s.lines) }
func (s Sorter) Swap(i, j int) { s.lines[i], s.lines[j] = s.lines[j], s.lines[i] }

func (s Sorter) Less(i, j int) bool {
	res := s.comparator(s.lines[i].Key, s.lines[j].Key)
	if s.reverse {
		return res > 0
	}
	return res < 0
}

// RemoveDuplicates удаляет дубликаты
func RemoveDuplicates(lines []LineData) []LineData {
	seen := make(map[string]struct{}, len(lines))
	var result []LineData
	for _, l := range lines {
		if _, ok := seen[l.Original]; !ok {
			seen[l.Original] = struct{}{}
			result = append(result, l)
		}
	}
	return result
}
