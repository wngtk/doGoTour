https://go.dev/tour/basics/1
https://go.dev/doc/tutorial/getting-started

## Go 语言的包的机制

> Every Go program is made up of packages.
> Programs start running in package main.

a package is a way to group functions, and it's made up of all the files in the same directory.

Go 语言包的机制和 Java 很像，Go 的源代码文件必须声明包，而且在目录层级上有要求。同一个目录里的 .go 文件都是必须是同一个包下面的，因为一个包是由同一个目录下的所有的文件组成的，那要是有文件是异类说我不是那个包的，那岂不是矛盾了。


## Go 语言有自己的一套目录层级规范

Go 语言有 Dependency tracking (所有的依赖都在 GOPATH 下，而不会污染系统的库的目录)。Go 语言提供的工具链也都是在 go/bin 这样的目录下，也不污染系统安装程序的目录。

```sh
go init mod example/hello # Enable dependency tracking for your code.
go mod tidy # Add new module requirements and sums.
```

go get 可以下载包。go mod tidy 会生成一个 go.sum 文件，用于模块认证。go init mod 会生成一个 go.mod 用来管理依赖的模块。

## Go 语言使用第三方库

import 包名的前面加上包的位置