# 変数

## はじめに

Go 言語における全ての変数は型を持っており、大きく分けると以下の 3 種類に分かれる。

* 値型
  * 整数や実数など `値` そのものを格納する変数
* 参照型
  * `スライス`, `マップ`, `チャネル` という 3 種類のデータ構造のいずれかを示す。
* ポインタ型
  * 値や関数といったメモリ上の実態をアドレス値によって間接的に表現する `ポインタ` を表す変数

## 変数の定義

Go の変数を定義する方法には明示的な書き方と暗黙的な書き方の 2 種類が存在する。

### 明示的な書き方

```go
var [変数名] [型]
```

```go
// int 型の変数 n を定義する
var n int
```

同じ型の変数であれば、複数の変数をまとめて定義することもできる。

```go
// int 型の変数 x, y. z を定義する
var x, y, z int
```

`var` の内容を `()` で囲うことで、異なる型の変数をまとめて定義することも可能。

<!-- markdownlint-disable MD010 -->

```go
var (
	x, y int
	name string
)
```

<!-- markdownlint-enable MD010 -->

定義した変数には `=` を使用して値を代入できる。再代入の制限はないが、異なる型の値を代入しようとするとコンパイルエラーが発生する。

以下のように複数の変数にまとめて代入を行うこともできる。

```go
var x, y int
x, y = 1, 2
x, y = 1, 2, 3 // 個数が異なるためコンパイルエラー
```

### 暗黙的な書き方

暗黙的な定義は型を指定する必要がない。

```go
i := 1
```

`:=` を使用することで変数の型定義と代入をまとめることができる。型指定を行なっていないが今回のように整数 `1` を代入すると暗黙的に変数 `i` は `int` 型であると決定される。(型推論)

関数の戻り値を利用することで暗黙的に変数を定義することも可能。

<!-- markdownlint-disable MD010 -->

```go
func getAge() int {
	return 20
}

// int 型の変数 age が定義され 20 が代入される
age := getAge
```

<!-- markdownlint-enable MD010 -->

暗黙的な変数の定義を利用した代入は 1 度しか許されないので注意する。

```go
i := 1 // int 型の変数 i を定義して 1 を代入
i := 2 // コンパイルエラー
```

同じ変数を複数回定義するとコンパイルエラーが発生する。

```go
var i int
var i int // コンパイルエラー
```

以下のような書き方も許されているが基本的には冗長で `:=` を使用した方がスマートである。

```go
var i = 1 // int 型の変数 i に 1 を代入

i := 1 // こう書いた方がスマート
```

しかし、`var` で変数定義をまとめる場合などは有用である。複数の変数を定義する場合は可能な限り `var` にまとめることでわかりやすくなる。

<!-- markdownlint-disable MD010 -->

```go
// var で変数定義をまとめる
var (
	age  = 20
	name = "hoge"
)

// 暗黙的な定義をまとめる
age := 20
name := "hoge"
```

<!-- markdownlint-enable MD010 -->

## パッケージ変数

Go の変数は定義される場所によって 2 種類に分かる。関数内に定義された変数は `ローカル変数` 関数の外に定義された変数は `パッケージ変数` になる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

var n = 100 // パッケージ変数

func main() {
	n = n + 1
	fmt.Printf("n = %d\n", n)
}
```

<!-- markdownlint-enable MD010 -->

パッケージ変数 `n` は `main` パッケージの中であればどこからでも参照することができる。

## 参照されない変数

Go では参照されない変数を定義したまま実行しようとするとコンパイルエラーが発生する。

<!-- markdownlint-disable MD010 -->

```go
func main() {
	age := 20
	name := "hoge" // 使用されない

	fmt.Println(age)
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
# command-line-arguments
./main.go:7:2: name declared and not used
```

***

* [Go to next "2. 配列"](./array.md)
* [Back to index page](../README.md)
