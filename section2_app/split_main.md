# main パッケージの分割

## 構成

```text
.
├── language
│   ├── english.go
│   └── japanese.go
├── app.go
└── main.go
```

## main.go

<!-- markdownlint-disable MD010 -->

```go
package main

import (
	"fmt"

	"./language"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(language.English())
	fmt.Println(language.Japanese())
}
```

<!-- markdownlint-enble MD010 -->

## app.go

<!-- markdownlint-disable MD010 -->

```go
package main

func AppName() string {
	return "language application"
}
```

<!-- markdownlint-enble MD010 -->

## 実行結果

```bash
$ go run main.go app.go
language application
English
Japanese
```

## 解説

`main` パッケージに属している `app.go` を作成、`main.go` 内で呼び出すようにした。

これをそのまま `go run` コマンドで実行しようとするとコンパイルエラーが発生する。

```bash
$ go run main.go
# command-line-arguments
./main.go:10:14: undefined: AppName
```

`go run` コマンドは明示的に指定された `main.go` のみを対象とし、 `app.go` は無視する。そのため、`app.go` も明示的に指定をするか、ワイルドカードを使用する必要がある。

```bash
# app.go も指定する
$ go run main.go app.go
language application
English
Japanese

# ワイルドカードを使用する
$ go run *.go
language application
English
Japanese
```

`go build` コマンドではファイルの指定がない場合は `go build *.go` と同じ振る舞いをする。

***

* [Go to next "3. テストをする" page](./test.md)
* [Back to previous "1. 簡単なアプリケーションの作成" page](./application.md)
* [Back to index page](../README.md)