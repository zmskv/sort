package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/zmskv/sort/internal/checker"
	"github.com/zmskv/sort/internal/comparator"
	"github.com/zmskv/sort/internal/options"
	"github.com/zmskv/sort/internal/reader"
	"github.com/zmskv/sort/internal/sorter"
)

func main() {
	opts := options.Parse()

	r, closer, err := reader.Open(opts.Input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error input data: %v\n", err)
		os.Exit(1)
	}
	defer closer()

	lines, err := reader.LoadLines(r, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading lines: %v\n", err)
		os.Exit(1)
	}

	if opts.Check {
		if checker.CheckSorted(lines, opts) {
			fmt.Println("OK")
		} else {
			fmt.Println("Data is not sorted")
			os.Exit(1)
		}
		return
	}

	cmp := comparator.Make(opts)
	sort.Sort(sorter.New(lines, cmp, opts.Reverse))

	if opts.Unique {
		lines = sorter.RemoveDuplicates(lines)
	}

	w := bufio.NewWriter(os.Stdout)
	for _, l := range lines {
		fmt.Fprintln(w, l.Original)
	}
	w.Flush()
}
