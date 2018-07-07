# Hello World

## プログラム

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

<!-- markdownlint-enable MD010 -->

## go run コマンド

Go は基本的にコンパイルを前提とした言語だが、`go run` コマンドを使用するとビルドプロセスを隠蔽しつつ直接プログラムを実行することができる。

```bash
# go run [file path]
$ go run main.go
hello world
```

## go build コマンド

コンパイルを行い、実行ファイルを作成するコマンド。

```bash
# go build [file path]
$ go build main.go
$ ls
main    main.go
$ ./main
hello world
```

## package

```go
package main
```

変数や関数などプログラムの全ての要素は何らかのパッケージに属する。Go のプログラムはパッケージの宣言から始める。

`package main` によってこのプログラムが `main` パッケージであることが示されている。Go では 1 ファイルに記述できるのは 1 パッケージのみという制約があり、満たさないとコンパイルエラーが発生する。

## import

```go
import "fmt"
```

プログラムで使用するパッケージを宣言する。参照されないパッケージを使用するとコンパイルエラーが発生する。

`Println` は `fmt` パッケージに含まれている関数。

## main 関数

```go
func main() {
	fmt.Println("hello world")
}
```

Go のプログラムが実行される開始場所は `main` パッケージ内の関数 `main` と定めされている。

***

* [Back to previous "2. 環境構築"](./install.md)
* [Back to index page](../README.md)
