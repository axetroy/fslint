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

```json
{
  "exclude": ["vendor", "node_modules", "bin"],
  "include": [
    {
      "file": "./src/**/*.vue", // lint for file
      "level": "error",
      "pattern": "CamelCase",
      "ignore": ["index.vue"] // ignore index.vue in this rule
    },
    {
      "folder": "./src/**/*", // lint for folder
      "level": "error",
      "pattern": "kebab-kebab"
    }
  ]
}
```

| Pattern         | Description                                                        | Recommend |
| --------------- | ------------------------------------------------------------------ | --------- |
| **CamelCase**   | Camel case with initial capital                                    | Yes       |
| **camelCase**   | Camel case with lowercase initials                                 |           |
| **Kebab-Kebab** | Uppercase letters and concatenated by symbols `-`                  |           |
| **kebab-kebab** | Lowercase letters and concatenated by symbols `-`                  | Yes       |
| **KEBAB-KEBAB** | TODO                                                               | Yes       |
| **Snake_Case**  | Uppercase letters snake case style and concatenated by symbols `_` |           |
| **snake_case**  | Lowercase letters snake case style and concatenated by symbols `_` | Yes       |
| **SNAKE_CASE**  | Lowercase letters snake case style and concatenated by symbols `_` | Yes       |
| **dot.dot**     | TODO                                                               |           |
| **Dot.Dot**     | TODO                                                               |           |
| **DOT.DOT**     | TODO                                                               |           |

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
curl -fsSL -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/axetroy/fslint/master/install.sh | bash -s v0.1.0
# or install from gobinaries.com
curl -sf https://gobinaries.com/axetroy/fslint@v0.1.0 | sh
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/fslint/releases)

### Build from source code

Make sure you have `Golang@v1.15.x` installed.

```shell
$ git clone https://github.com/axetroy/fslint.git $GOPATH/src/github.com/axetroy/fslint
$ cd $GOPATH/src/github.com/axetroy/fslint
$ make build
```

### Test

```bash
$ make test
```

### License

The [Anti-996 License](LICENSE)
