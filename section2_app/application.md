# 簡単なアプリケーションの作成

## プログラムの構造

```text
.
├── language
│   ├── english.go
│   └── japanese.go
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
	fmt.Printf(language.English())
	fmt.Printf(language.Japanese())
}
```

<!-- markdownlint-enable MD010 -->

## english.go

<!-- markdownlint-disable MD010 -->

```go
package language

func English() string {
	return "English"
}
```

<!-- markdownlint-enable MD010 -->

## japanese.go

<!-- markdownlint-disable MD010 -->

```go
package language

func Japanese() string {
	return "Japanese"
}
```

<!-- markdownlint-enable MD010 -->

## 実行結果

```bash
$ go run main.go
English
Japanese
```

## 解説

`english.go` と `japanese.go` は同じ `language` に属している。Go では 1 つのパッケージを複数のソースコードを使って定義することができる。

Go で import に指定するパッケージは通常は `GOPATH` に指定されたディレクトリから探索されるが、`main.go` 内の `./language` のように相対パスで記述すると、import が書かれているファイルからの相対位置に置かれているディレクトリを指定することができる。

***

* [Go to next "2. main パッケージの分割"](./split_main.md)
* [Back to index page](../README.md)
