package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// Options args
type Options struct {
	Target, Saveto             string
	Concurrency                int
	Silent, Verbose, UpdateTld bool
	Output                     *os.File
	Domains                    []string
}

var o *Options

func init() {
	o = &Options{}
}

func init() {
	o = &Options{}

	flag.StringVar(&o.Target, "t", "", "")
	flag.StringVar(&o.Target, "target", "", "")

	flag.StringVar(&o.Saveto, "o", "", "")
	flag.StringVar(&o.Saveto, "output", "", "")

	flag.IntVar(&o.Concurrency, "c", 20, "")
	flag.IntVar(&o.Concurrency, "concurrent", 20, "")

	flag.BoolVar(&o.Silent, "s", false, "")
	flag.BoolVar(&o.Silent, "silent", false, "")

	flag.BoolVar(&o.Verbose, "v", false, "")
	flag.BoolVar(&o.Verbose, "verbose", false, "")

	flag.BoolVar(&o.UpdateTld, "u", false, "")
	flag.BoolVar(&o.UpdateTld, "update-domain", false, "")

	// Override help text
	flag.Usage = func() {
		banner()
		h := []string{
			"",
			"Usage:" + usage,
			"",
			"Options:" + options,
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()
}

// ParseOptions will parse given args
func ParseOptions() *Options {
	// Show banner to user
	if !o.Silent {
		banner()
	}

	// Validate input options
	o.validate()

	return o
}

func banner() {
	gologger.Print().Msgf("%s\n\n", header)
}
