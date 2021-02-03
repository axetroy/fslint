English | [中文简体](README_zh-CN.md)

[![Build Status](https://github.com/axetroy/fslint/workflows/ci/badge.svg)](https://github.com/axetroy/fslint/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/fslint)](https://goreportcard.com/report/github.com/axetroy/fslint)
![Latest Version](https://img.shields.io/github/v/release/axetroy/fslint.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/fslint.svg)

## fslint

This is a tool for detecting file system naming styles. It is hard to imagine that there are several different styles of file naming in the same application.

### Usage

```bash
$ fslint --config=.fslintrc.json
```

`.fslintrc.json` is a JSON file look like this:

```jsonc
{
  "exclude": ["vendor", "node_modules", "bin"],
  "include": [
    {
      "file": "./src/**/*.vue", // lint for file
      "level": "error",
      "pattern": "PascalCase",
      "ignore": ["**/index.vue"] // ignore index.vue in this rule
    },
    {
      "folder": "./src/**/*", // lint for folder
      "level": "error",
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
| **KEBAB-KEBAB**  | All characters are uppercase and concatenated by symbols `-`       |           |
| **Snake_Case**   | Uppercase letters snake case style and concatenated by symbols `_` |           |
| **snake_case**   | Lowercase letters snake case style and concatenated by symbols `_` | Yes       |
| **SNAKE_CASE**   | All characters are uppercase and concatenated by symbols `_`       |           |
| **Dot.Dot**      | Start with lowercase, with `.` as the delimiter style              |           |
| **dot.dot**      | Start with uppercase, with `.` as the delimiter style              |           |
| **DOT.DOT**      | All characters are uppercase and concatenated by symbols `.`       |           |
| **/\<regexp\>/** | regular expression start with `/` and end with `/`                 |           |

### Installation

If you have installed nodejs, you can install it via npm

```bash
npm install @axetroy/fslint -g
```

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/axetroy/fslint/master/install.sh | bash
# or install specified version
curl -fsSL -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/axetroy/fslint/master/install.sh | bash -s v0.2.17
# or install from gobinaries.com
curl -sf https://gobinaries.com/axetroy/fslint@v0.2.17 | sh
```

Or Download the executable file for your platform at [release page](https://github.com/axetroy/fslint/releases)

### Test

```bash
$ make test
```

### License

The [Anti-996 License](LICENSE)
