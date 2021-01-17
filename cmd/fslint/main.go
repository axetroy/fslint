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
				level := color.YellowString("warn ")

				switch result.Level {
				case fslint.LevelWarn:
					level = color.YellowString("warn ")
					warnNum = warnNum + 1
				case fslint.LevelError:
					level = color.RedString("error")
					errorNum = errorNum + 1
				}

				info := fmt.Sprintf("[fslint]: %s '%s' not match '%v'\n", level, color.BlueString(result.FilePath), color.GreenString(string(result.Expect)))

				_, err := os.Stderr.WriteString(info)

				if err != nil {
					return errors.WithStack(err)
				}
			}

			msg := fmt.Sprintf("[fslint]: finish with %s warning and %s error.\n", color.YellowString(fmt.Sprint(warnNum)), color.RedString(fmt.Sprint(errorNum)))
			os.Stderr.WriteString(msg)

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
