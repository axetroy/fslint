[中文简体](README.md) | English

[![Build Status](https://github.com/axetroy/fslint/workflows/ci/badge.svg)](https://github.com/axetroy/fslint/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/fslint)](https://goreportcard.com/report/github.com/axetroy/fslint)
![Latest Version](https://img.shields.io/github/v/release/axetroy/fslint.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/fslint.svg)

## fslint

This is a tool for detecting file system naming styles. It is hard to imagine that there are several different styles of file naming in the same application.

### Usage

```bash
$ fslint --config=.fslintrc.jsonc
```

`.fslintrc.jsonc` is a JSON file look like this:

```jsonc
{
  "maxError": 0, // Max amount of error
  "exclude": ["vendor", "node_modules", "bin", ".git"], // The folder name will be ignore
  "include": [
    {
      "file": "./src/**/*.vue", // lint for file, support Glob grammar
      "level": "error",
      "pattern": "PascalCase",
      "ignore": ["**/index.vue"] // ignore index.vue in this rule, support Glob grammar
    },
    {
      "folder": "./src/**/*", // lint for folder
      "level": "warn",
      "pattern": "kebab-case"
    }
  ]
}
```

| Pattern          | Description                                                        | Recommend |
| ---------------- | ------------------------------------------------------------------ | --------- |
| **PascalCase**   | Pascal Case style                                                  | Yes       |
| **camelCase**    | Camel Case style                                                   | Yes       |
| **kebab-case**   | Lowercase letters and concatenated by symbols `-`                  | Yes       |
| **snake_case**   | Lowercase letters snake case style and concatenated by symbols `_` | Yes       |
| **/\<regexp\>/** | regular expression start with `/` and end with `/`                 |           |

### Install

1. Shell (Mac/Linux)

```bash
curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=axetroy/fslint
```

2. PowerShell (Windows):

```bash
$r="axetroy/fslint";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/axetroy/fslint/releases) (All platforms)

> download the executable file and put the executable file to `$PATH`

4. Build and install from source using [Golang](https://golang.org) (All platforms)

```bash
go install github.com/axetroy/fslint/cmd/fslint
```

5. Install via npm

```sh
npm install @axetroy/fslint -g
```

### Test

```bash
$ make test
```

### License

The [Anti-996 License](LICENSE)
