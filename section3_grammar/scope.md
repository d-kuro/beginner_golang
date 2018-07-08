# スコープ

## 識別子の命名規則

パッケージに定義された定数、変数、関数などが他のパッケージから参照可能であるかは、識別子の 1 文字目が大文字であるかどうかで決定される。

<!-- markdownlint-disable MD010 -->

```go
// 公開される
func DoSomething() {
	// 略
}

// パッケージ内のみで参照できる
func doSomething() {
	// 略
}
```

<!-- markdownlint-enable MD010 -->

## import

`import` したパッケージに任意のパッケージ名を付与することができる。パッケージ名は上書きされるので、付与した場合は元のパッケージ名で参照することはできない。

<!-- markdownlint-disable MD010 -->

```go
import (
	f "fmt" // パッケージ名 f を指定
)
```

<!-- markdownlint-enable MD010 -->

`.` を使用することでパッケージ名を省略することも可能だが、識別子の重複に注意する必要がある。

<!-- markdownlint-disable MD010 -->

```go
package main

import (
	. "fmt"
)

func main() {
	Println("hoge")
}
```

<!-- markdownlint-enable MD010 -->

## ファイルのスコープ

Go では 1 つのパッケージを定義するために複数の Go ファイルを使用できる。パッケージが複数のファイルによって構成されている場合、`import` 宣言は各ファイル内でのみ有効になる。

## 関数のスコープ

関数の中で定義された変数や定数は、定義された関数の中でのみ参照が可能。

関数のスコープでは、引数の変数と戻り値の変数も含まれる。

<!-- markdownlint-disable MD010 -->

```go
func doSomething(a int) (b string) {
	var a int    // 定義済み
	var b string // 定義済み
}
```

<!-- markdownlint-enable MD010 -->

関数の中では `{}` を使用して、明示的に別のブロックを定義することができる。ブロックを分けることで関数スコープにある識別子と重複する識別子を使用することが可能。

<!-- markdownlint-disable MD010 -->

```go
func doSomething(a int) (b string) {
	{
		var a int
		var b string
	}
	return b
}
```

<!-- markdownlint-enable MD010 -->

***

* [Back to previous "4. 定数"](./const.md)
* [Back to index page](../README.md)
