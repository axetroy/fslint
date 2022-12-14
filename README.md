中文简体 | [English](README_en-US.md)

[![Build Status](https://github.com/axetroy/fslint/workflows/ci/badge.svg)](https://github.com/axetroy/fslint/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/fslint)](https://goreportcard.com/report/github.com/axetroy/fslint)
![Latest Version](https://img.shields.io/github/v/release/axetroy/fslint.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/fslint.svg)

## fslint

这是用于检测文件系统命名风格的工具。 很难想象在同一应用程序中有几种不同的文件命名风格。

### Usage

```bash
$ fslint --config=.fslintrc.jsonc
```

`.fslintrc.jsonc` 是一个配置文件

```jsonc
{
  "exclude": ["vendor", "node_modules", "bin", ".git"], // 忽略的目录名称
  "include": [
    {
      "file": "./src/**/*.vue", // 检测 *.vue 文件，支持 Glob 语法
      "level": "error",
      "pattern": "PascalCase",
      "ignore": ["**/index.vue"] // 在这条规则中忽略 index.vue 文件，支持 Glob 语法
    },
    {
      "folder": "./src/**/*", // 检测文件夹
      "level": "warn",
      "pattern": "kebab-case"
    }
  ]
}
```

| Pattern          | 描述                             | 推荐 |
| ---------------- | -------------------------------- | ---- |
| **PascalCase**   | 大写的驼峰式风格                 | Yes  |
| **camelCase**    | 小写的驼峰式风格                 | Yes  |
| **kebab-case**   | 使用 破折号(`-`) 连接的小写风格  | Yes  |
| **snake_case**   | 小写开头的蛇式风格               | Yes  |
| **/\<regexp\>/** | 以 `/` 开头 `/` 结尾的正则表达式 |      |

### 安装

1. Shell (Mac/Linux)

```bash
curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=axetroy/fslint
```

2. PowerShell (Windows):

```bash
$r="axetroy/fslint";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/axetroy/fslint/releases) (全平台支持))

> 下载可执行文件，并且把它加入到`$PATH` 环境变量中

4. 使用 [Golang](https://golang.org) 从源码中构建并安装 (全平台支持)

```bash
go install github.com/axetroy/fslint/cmd/fslint
```

5. 通过 npm 安装

```sh
npm install @axetroy/fslint -g
```

### 测试

```bash
$ make test
```

### 开源许可

The [Anti-996 License](LICENSE_zh-CN)
