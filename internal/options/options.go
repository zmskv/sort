package options

import "flag"

// Options - опции
type Options struct {
	Column  int
	Numeric bool
	Reverse bool
	Unique  bool
	Month   bool
	IgnoreB bool
	Check   bool
	Human   bool
	Input   string
}

// Parse обрабатывает командную строку
func Parse() Options {
	var opts Options

	flag.IntVar(&opts.Column, "k", 0, "Sort by column N")
	flag.BoolVar(&opts.Numeric, "n", false, "Sort numerically")
	flag.BoolVar(&opts.Reverse, "r", false, "Reverse order")
	flag.BoolVar(&opts.Unique, "u", false, "Output only unique lines")
	flag.BoolVar(&opts.Month, "M", false, "Sort by month name (Jan, Feb, ... Dec)")
	flag.BoolVar(&opts.IgnoreB, "b", false, "Ignore trailing blanks")
	flag.BoolVar(&opts.Check, "c", false, "Check whether data is sorted")
	flag.BoolVar(&opts.Human, "h", false, "Sort human-readable numbers (1K, 2M, ...)")
	flag.Parse()

	if flag.NArg() > 0 {
		opts.Input = flag.Arg(0)
	}

	return opts
}
