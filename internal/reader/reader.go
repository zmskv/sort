package reader

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/zmskv/sort/internal/options"
	"github.com/zmskv/sort/internal/sorter"
)

func trimTrailing(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func splitColumns(line string) []string {
	return strings.FieldsFunc(line, unicode.IsSpace)
}

// Open - открывает файл
func Open(path string) (io.Reader, func(), error) {
	if path == "" {
		return os.Stdin, func() {}, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, func() {}, err
	}
	return file, func() { file.Close() }, nil
}

// LoadLines - загружает строки
func LoadLines(r io.Reader, opts options.Options) ([]sorter.LineData, error) {
	scanner := bufio.NewScanner(r)
	var lines []sorter.LineData
	for scanner.Scan() {
		line := scanner.Text()
		key := parseColumnKey(line, opts)
		lines = append(lines, sorter.LineData{
			Original: line,
			Key:      key,
		})
	}
	return lines, scanner.Err()
}

func parseColumnKey(line string, opts options.Options) string {
	if opts.IgnoreB {
		line = trimTrailing(line)
	}
	cols := splitColumns(line)
	if opts.Column > 0 && opts.Column <= len(cols) {
		return cols[opts.Column-1]
	}
	return line
}
