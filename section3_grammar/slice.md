# スライス

## スライスとは

可変長配列を表現する型。スライスを表す型は `[]` の後ろに任意の型を置いて表現する。

```go
var s []int
```

スライスを生成するには組み込み関数の `make` を使用する。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Println(s)
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
[0 0 0 0 0]
```

## 代入と参照

スライスに格納されている各要素への代入、参照は、配列と同様に `[n]` のように要素のインデックスを指定する。代入も配列と同様に操作することができる。

スライスは可変長配列だが、スライスの要素数を超えた要素へのアクセスはランタイムパニックを発生させる。

```go
s := make([]int, 5)
fmt.Println(s)      // [0 0 0 0 0]
s[0] = 1
fmt.Println(s[0])   // 1
fmt.Println(s[5])   // ランタイムパニック
```

## len

スライスは可変長配列なので、動的に要素数が変化する。現在の要素数を調べるためには組み込み関数 `len` を使用する。

```go
s := make([]int, 5)
length := len(s)
fmt.Println(length) // 5
```

`len` はスライスだけではなく配列型に対しても使用することができる。

***

* [Back to previous "13. 参照型"](./reference.md)
* [Back to index page](../README.md)
