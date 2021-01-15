package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/axetroy/fslint"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printHelp() {
	println(`fslint - a cli tool for lint file system style

USAGE:
  fslint [OPTIONS] [version...]

ARGUMENTS:
  [version...]  Optional version or version range.

OPTIONS:
  --help        Print help information.
  --version     Print version information.

SOURCE CODE:
  https://github.com/axetroy/fslint`)
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
		configPath  string
	)

	flag.StringVar(&configPath, "config", ".fslint.json", "The config of fslint")
	flag.BoolVar(&showHelp, "help", false, "Print help information")
	flag.BoolVar(&showVersion, "version", false, "Print version information")

	flag.Usage = printHelp

	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	if showVersion {
		println(fmt.Sprintf("%s %s %s", version, commit, date))
		os.Exit(0)
	}

	if results, err := fslint.Lint(configPath); err != nil {
		return errors.WithStack(err)
	} else {
		for _, result := range results {
			fmt.Printf("fslint: '%s' not match '%v'\n", result.FilePath, result.Expect)
		}
	}

	return nil
}

func main() {
	var (
		err error
	)

	defer func() {
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(255)
		}
	}()

	err = run()
}
