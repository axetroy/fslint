package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/axetroy/fslint"
	"github.com/gookit/color"
	"github.com/pkg/errors"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	red    = color.FgRed.Render
	green  = color.FgGreen.Render
	blue   = color.FgLightMagenta.Render
	yellow = color.FgYellow.Render
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
  --config      Specify ‘.fslintrc.json’. defaults to '.fslintrc.json'.
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

	flag.StringVar(&configPath, "config", ".fslintrc.json", "The config of fslint")
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
			b, err := json.Marshal(results.Values())

			if err != nil {
				return errors.WithStack(err)
			}

			if _, err = os.Stdout.Write(b); err != nil {
				return errors.WithStack(err)
			}
		} else {
			var (
				errorNum = 0
				warnNum  = 0
			)

			for _, result := range results.Values() {
				level := "warn "

				switch result.Level {
				case fslint.LevelWarn:
					level = yellow("warn ")
					warnNum = warnNum + 1
				case fslint.LevelError:
					level = red("error")
					errorNum = errorNum + 1
				}

				color.Fprintf(os.Stderr, "[fslint]: %s '%s' not match with pattern '%v'\n", level, blue(result.FilePath), green(result.Expect))

				if err != nil {
					return errors.WithStack(err)
				}
			}

			color.Fprintf(os.Stderr, "[fslint]: finish with %s warning and %s error.\n", yellow(warnNum), red(errorNum))

			if results.ErrorCount() > 0 {
				os.Exit(1)
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
