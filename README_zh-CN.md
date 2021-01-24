[English](README.md) | 中文简体

[![Build Status](https://github.com/axetroy/fslint/workflows/ci/badge.svg)](https://github.com/axetroy/fslint/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/fslint)](https://goreportcard.com/report/github.com/axetroy/fslint)
![Latest Version](https://img.shields.io/github/v/release/axetroy/fslint.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/fslint.svg)

## fslint

这是用于检测文件系统命名风格的工具。 很难想象在同一应用程序中有几种不同的文件命名风格。

### Usage

```bash
$ fslint --config=.fslintrc.json
```

`.fslintrc.json` 是一个配置文件

```jsonc
{
  "exclude": ["vendor", "node_modules", "bin"],
  "include": [
    {
      "file": "./src/**/*.vue", // 检测 *.vue 文件
      "level": "error",
      "pattern": "CamelCase",
      "ignore": ["**/index.vue"] // 在这条规则中忽略 index.vue 文件
    },
    {
      "folder": "./src/**/*", // 检测文件夹
      "level": "error",
      "pattern": "kebab-kebab"
    }
  ]
}
```

| Pattern          | 描述                                     | 推荐 |
| ---------------- | ---------------------------------------- | ---- |
| **CamelCase**    | 大写开头的驼峰式风格                     | Yes  |
| **camelCase**    | 小写开头的驼峰式风格                     |      |
| **Kebab-Kebab**  | 使用 破折号(`-`) 连接的大写开头风格      |      |
| **kebab-kebab**  | 使用 破折号(`-`) 连接的小写风格          | Yes  |
| **KEBAB-KEBAB**  | 使用 破折号(`-`) 连接的全全写风格        |      |
| **Snake_Case**   | 大写开头的蛇式风格                       |      |
| **snake_case**   | 小写开头的蛇式风格                       | Yes  |
| **SNAKE_CASE**   | 使用 点号(`_`) 连接的全大写蛇式风格      |      |
| **Dot.Dot**      | 使用 点号(`.`) 连接的大写开头风格        |      |
| **dot.dot**      | 使用 点号(`.`) 连接的小写开头风格        |      |
| **DOT.DOT**      | 使用 点号(`.`) 连接的全大写风格          |      |
| **/\<regexp\>/** | 以 `/`开头和结尾的字符串被视为正则表达式 |      |

### 安装

如果你已安装 nodejs，则可以通过 npm 进行安装

```bash
npm install @axetroy/fslint -g
```

如果你使用的是 Linux/macOS，你可以通过以下命令进行安装

```shell
# 安装最新版本
curl -fsSL -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/axetroy/fslint/master/install.sh | bash
# 或者安装指定版本
curl -fsSL -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/axetroy/fslint/master/install.sh | bash -s v0.1.0
# 或者通过 gobinaries.com 安装
curl -sf https://gobinaries.com/axetroy/fslint@v0.1.0 | sh
```

从[release page](https://github.com/axetroy/fslint/releases)页面下载对应平台的可执行文件，并且把它加入到 `$PATH` 环境变量中，并尝试以下命令

### 测试

```bash
$ make test
```

### 开源许可

The [Anti-996 License](LICENSE_zh-CN)
