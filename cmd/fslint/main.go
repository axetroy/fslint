package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

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

	startAt = time.Now()
)

func printHelp() {
	println(`fslint - a file system naming style lint

USAGE:
  fslint [OPTIONS]

OPTIONS:
  --help                 Print help information
  --version              Print version information
  --config <filepath>    Specify fslint config. defaults to '.fslintrc.jsonc'
  --json                 Output the lint result as JSON
  --no-color             Disabled color for printing
  --max-error <number>   Set max error. defaults to zero

SOURCE CODE:
  https://github.com/axetroy/fslint`)
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
		configPath  string
		outputJSON  bool
		noColor     bool
		maxError    int
	)

	flag.StringVar(&configPath, "config", ".fslintrc.jsonc", "The config of fslint")
	flag.BoolVar(&outputJSON, "json", false, "Output the lint result as JSON")
	flag.BoolVar(&noColor, "no-color", false, "disabled color for printing")
	flag.IntVar(&maxError, "max-error", 0, "Set max error")
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

	if color.SupportColor() {
		color.Enable = !noColor
	} else {
		color.Enable = false
	}

	results, err := fslint.Lint(configPath, &fslint.LintOptions{MaxError: &maxError})

	if err != nil {
		return errors.WithStack(err)
	}

	if outputJSON {
		b, err := json.Marshal(results.Values())

		if err != nil {
			return errors.WithStack(err)
		}

		if _, err = os.Stdout.Write(b); err != nil {
			return errors.WithStack(err)
		}
	} else {
		for _, result := range results.Values() {
			level := "warn"

			switch result.Level {
			case fslint.LevelWarn:
				level = yellow("warn")
			case fslint.LevelError:
				level = red("error")
			}

			color.Fprintf(os.Stderr, "[fslint]: %s '%s' not match with pattern '%v'\n", level, blue(result.FilePath), green(result.Expect))

			if err != nil {
				return errors.WithStack(err)
			}
		}

		elapsed := time.Since(startAt)

		color.Fprintf(os.Stderr, "[fslint]: finish with %s warning and %s error in %s.\n", yellow(results.WarnCount()), red(results.ErrorCount()), green(elapsed))

		if results.ErrorCount() > 0 {
			os.Exit(1)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(255)
	}
}
