package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/axetroy/fslint"
	"github.com/fatih/color"
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
  --config      Specify ‘.fslint.json’. defaults to '.fslint.json'.
  --json        Output the lint result as JSON

SOURCE CODE:
  https://github.com/axetroy/fslint`)
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
		configPath  string
		outputJSON  bool
	)

	flag.StringVar(&configPath, "config", ".fslint.json", "The config of fslint")
	flag.BoolVar(&outputJSON, "json", false, "Output the lint result as JSON")
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
		if outputJSON {
			b, err := json.Marshal(results)

			if err != nil {
				return errors.WithStack(err)
			}

			if _, err = os.Stdout.Write(b); err != nil {
				return errors.WithStack(err)
			}
		} else {
			for _, result := range results {
				info := fmt.Sprintf("[fslint]: '%s' not match '%v'\n", color.YellowString(result.FilePath), color.GreenString(string(result.Expect)))

				_, err := os.Stderr.WriteString(info)

				if err != nil {
					return errors.WithStack(err)
				}
			}
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(255)
	}
}
