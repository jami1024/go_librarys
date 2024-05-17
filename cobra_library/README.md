## 如何构建Go命令行工具
> 本文引用于[孔令飞的云原生实战营](https://articles.zsxq.com/id_8r01wr1mrfdp.html)


在 Go 项目开发中，我们需要编写 main 函数，并编译成为二进制文件， 部署启动服务。有多种方式可以开发一个 main 函数。
例如你可以手撸一个 main 函数，并在 main 函数中处理命令行参数，配置文件解析，应用初始化等操作。
如下所示：

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 解析命令行参数
	option1 := flag.String("option1", "default_value", "Description of option 1")
	option2 := flag.Int("option2", 0, "Description of option 2")
	flag.Parse()

	// 执行简单的业务逻辑
	fmt.Println("Option 1 value:", *option1)
	fmt.Println("Option 2 value:", *option2)

	// 在这里添加您的业务逻辑代码
}
```
虽然可以手撸一个应用，但是开发效率低下，要处理各种场景，开发出来的应用还不怎么优雅。为了，解决这些问题，社区涌现出了大批优秀的应用开发框架，
例如：kingpin、cli、cobra 等。开发者可以直接复用这些应用开发框架，来构建优秀的命令行工具，这也是当前开发应用时，采用最多的方式。
当前社区最受欢迎的应用开发框架是 cobra。本文就来详细介绍下 cobra 框架的功能及使用方式。

## Cobra 包介绍
Cobra 是一个用于创建现代命令行界面（CLI）应用程序的库，同时它也提供了一个名为 cobra-cli 的命令行工具，
用于生成应用程序和命令文件。许多大型项目，如 Kubernetes、Docker、Etcd、Rkt 和 Hugo，都使用 Cobra 来构建它们的 CLI。
Cobra 提供了多种特性，包括：

- 支持创建具有嵌套子命令的 CLI，例如 app server 和 app fetch。
- 可以通过 cobra-cli init appname 和 cobra-cli add cmdname 快速创建应用和子命令。
- 提供智能命令建议，例如在输入错误时提示“app srver… did you mean app server?”。
- 自动生成命令和标志的帮助文本，并能识别 -h、--help 等帮助标志。
- 自动为应用程序生成 bash、zsh、fish 和 powershell 的自动补全脚本。
- 支持命令别名、自定义帮助信息和自定义用法等功能。
- 可以与 viper（配置库）和 pflag（命令行标志库）紧密集成，用于构建遵循十二因素应用原则的应用程序。

Cobra 是基于命令（commands）、参数（arguments）和标志（flags）的结构来构建的。 命令是指操作的类型，参数是指非选项参数，标志是指选项参数。
一个优秀的应用程序应该具有清晰的用法，让用户容易理解如何操作。
应用程序的命令行通常遵循“APPNAME VERB NOUN --ADJECTIVE”或“APPNAME COMMAND ARG --FLAG”的模式，例如
```shell
git clone URL --bare # clone 是一个命令，URL 是一个非选项参数，bare 是一个选项参数
```
这里，VERB 代表动词，NOUN 代码名词，ADJECTIVE 代表形容词。

## cobra-cli 命令安装

```shell
go install github.com/spf13/cobra-cli@latest
```
cobra-cli 命令提供了 4 个子命令：
- init：初始化一个 cobra 应用程序；
- add：给通过 cobra init 创建的应用程序添加子命令；
- completion：为指定的 shell 生成命令自动补全脚本；
- help：打印任意命令的帮助信息。

cobra-cli 命令还提供了一些全局的参数：
- -a, --author：指定 Copyright 版权声明中的作者；
- --config：指定 cobra 配置文件的路径；
- -l, --license：指定生成的应用程序所使用的开源协议，内置的有：GPLv2, GPLv3, LGPL, AGPL, MIT, 2-Clause BSD or 3-Clause BSD；
- --viper：使用 viper 作为命令行参数解析工具，默认为 true。

## Cobra 使用方法
在构建 cobra 应用时，我们可以自行组织代码目录结构，但 cobra 建议如下目录结构：

```shell
▾ appName/
    ▾ cmd/
        add.go
        your.go
        commands.go
        here.go
      main.go
```
main.go 文件目的只有一个，初始化 cobra 应用：
```go
package main

import (
  "{pathToYourApp}/cmd"
)

func main() {
  cmd.Execute()
}
```

## Cobra 包常用功能
Cobra 包含了许多功能，帮助你创建高质量的应用程序。
在这一部分，我们将详细介绍 Cobra 包提供的核心功能和使用方法。
你可以使用 cobra-cli 命令行工具来快速创建一个应用程序并添加子命令，然后在这些生成的代码基础上进行进一步开发，这样可以提高开发效率。
具体操作步骤如下：

### 1. 生成应用程序

```shell
$ mkdir -p cobrademo && cd cobrademo && go mod init
$ cobra-cli init --license=MIT --viper
$ ls
cmd  go.mod  go.sum  LICENSE  main.go
```
> 提示：如果遇到错误 Error: invalid character '{' after top-level value)'}'，
> 可参考：https://github.com/spf13/cobra-cli/issues/26。

当一个应用程序被初始化之后，就可以给这个应用程序添加一些命令：

```shell
$ cobra-cli add serve
$ cobra-cli add config
$ cobra-cli add create -p 'configCmd' # 此命令的父命令的变量名（默认为 'rootCmd'）
$ ls cmd/
config.go  create.go  root.go  serve.go
```
执行`cobra-cli add` 之后，会在 cmd 目录下生成命令源码文件。
cobra-cli add 不仅可以添加命令，也可以添加子命令，
例如上面的例子，通过 cobra-cli add create -p 'configCmd' 给 config 命令添加了 create 子命令，-p 指定子命令的父命令：<父命令>Cmd。

### 2. 编译并执行

在生成完命令后，可以直接执行 go build 命令编译应用程序：


```shell
cd cobrademo
go build -v .
./main -h                                                                                 ✔ | 15:16:57 
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  main [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      A brief description of your command
  help        Help about any command
  serve       A brief description of your command

Flags:
      --config string   config file (default is $HOME/.main.yaml)
  -h, --help            help for main
  -t, --toggle          Help message for toggle

Use "main [command] --help" for more information about a command.

```
这里需要注意：命令名称要是 camelCase 格式，而不是 snake_case / snake-case 格式，如果不是驼峰格式，cobra 会报错。

### 3. 配置cobra
当你使用 Cobra 生成应用程序时，它会自动在当前目录创建一个 LICENSE 文件，并在生成的 Go 源代码文件中添加 LICENSE 头部（Header）。
LICENSE 文件和 LICENSE 头部的内容可以通过 Cobra 的配置文件进行自定义，默认的配置文件路径是 ~/.cobra.yaml。
例如，
```yaml
author: Steve Francia <spf@spf13.com>
year: 2020
license:
  header: This file is part of CLI application foo.
  text: |
    {{ .copyright }}

    This is my license. There are many like it, but this one is mine.
    My license is my best friend. It is my life. I must master it as I must
    master my life.
```
配置文件中可以设置作者、年份以及许可证的文本。
在提供的 YAML 配置文件示例中，{{ .copyright }} 标签的内容会根据作者（author）和年份（year）自动生成。
根据这个配置，生成的 LICENSE 文件将包含版权声明和自定义的许可证文本。

此外，Cobra 还支持使用内置的许可证，包括 GPLv2、GPLv3、LGPL、AGPL、MIT、2-Clause BSD 和 3-Clause BSD。
如果你想要使用 MIT 许可证，可以在初始化应用程序时使用 `cobra-cli init --license=MIT` 命令。
