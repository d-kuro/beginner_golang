# 配列

## 配列の定義

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
[1 2 3 4 5]
```

この場合の型名は `[5]int` になる。`{}` で囲ったブロックで配列の初期化処理を行なっている。

配列の要素を参照するには `[n]` の形式でインデックスを指定する。インデックスは 0 から始まり負の数は指定できない。

```go
a := [3]int{1, 2, 3}
a[0] // 1
a[1] // 2
a[2] // 3
```

`{}` で囲ったブロック内では、配列の要素数と同じ数の初期値を記述する必要はない。

```go
a := [5]int{}        // [0 0 0 0 0]
b := [5]int{1, 2, 3} // [1 2 3 0 0]

// 明示的に配列の変数を定義した場合は初期値を与えなかった場合と同じになる
var a [5]int         // [0 0 0 0 0]
```

Go では定義された型すべてで配列型を定義することができる。明示的な初期値が与えられない場合はゼロの意味を表す値で初期化される。`bool` 型の場合は `false`、`string` 型の場合は `""(空文字)` となる。

## 要素数の省略

要素数は `...` を使用して省略することが可能。この場合は与えた初期値の数が配列の要素数になる。

```go
a := [...]int{1, 2, 3}   // [3]int{1, 2, 3}
b := [...]int{}          // [0]int{}
```

## 要素への代入

`a[index] = v` の形式で配列の指定した要素のに値を代入することができる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	fmt.Printf("%v\n", a) // [1 2 3]
	a[0] = 0
	a[2] = 0
	fmt.Printf("%v\n", a) // [0 2 0]
}
```

<!-- markdownlint-enable MD010 -->

## 配列型の代入

配列型の要素数と要素の型が同一の変数であれば相互に代入を行うことが可能。

配列型の代入では全ての要素がコピーされるため、代入後に代入元の配列の要素を書き換えても代入先の要素には影響がない。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}

	fmt.Printf("%v\n", a) // [1 2 3]

	a = b

	fmt.Printf("%v\n", a) // [4 5 6]
	fmt.Printf("%v\n", b) // [4 5 6]

	a[0] = 0
	b[2] = 0

	fmt.Printf("%v\n", a) // [0 5 6]
	fmt.Printf("%v\n", b) // [4 5 0]
}
```

<!-- markdownlint-enable MD010 -->

## 配列の拡張について

Go では配列型の拡張や縮小は不可能でサイズは常に固定。`Slice` というデータ構造が Go では可変長配列に相当する。

***

* [Go to next "3. 関数"](./function.md)
* [Back to previous "1. 変数"](./variable.md)
* [Back to index page](../README.md)