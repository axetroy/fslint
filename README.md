[![Build Status](https://github.com/axetroy/fslint/workflows/ci/badge.svg)](https://github.com/axetroy/fslint/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/fslint)](https://goreportcard.com/report/github.com/axetroy/fslint)
![Latest Version](https://img.shields.io/github/v/release/axetroy/fslint.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/fslint.svg)

## fslint

a cli tool for lint File System

### Usage

```bash

```

### Installation

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL https://raw.githubusercontent.com/axetroy/fslint/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/axetroy/fslint/master/install.sh | bash -s v0.1.0
# or install from gobinaries.com
curl -sf https://gobinaries.com/axetroy/fslint@v0.1.0 | sh
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/fslint/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/bin"
```

then, try it out.

```bash
fslint --help
```

Finally, to use Deno correctly, you also need to set environment variables

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/.deno/bin"
```

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
