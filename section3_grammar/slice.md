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
	slice := make([]int, 5)
	fmt.Println(slice)
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
slice := make([]int, 5)
fmt.Println(slice)      // [0 0 0 0 0]
slice[0] = 1
fmt.Println(slice[0])   // 1
fmt.Println(slice[5])   // ランタイムパニック
```

## len

スライスは可変長配列なので、動的に要素数が変化する。現在の要素数を調べるためには組み込み関数 `len` を使用する。

```go
slice := make([]int, 5)
length := len(slice)
fmt.Println(length) // 5
```

`len` はスライスだけではなく配列型に対しても使用することができる。

## cap

スライスは要素数のほかに容量という属性を持っている。スライスの容量を調べるためには組み込み関数 `cap` を使用する。

```go
slice1 := make([]int, 5)
capacity1 := cap(slice1)
fmt.Println(capacity1)   // 5

slice2 := make([]int, 5, 10)
capacity2 := cap(slice2)
fmt.Println(capacity2)   // 10
```

## 簡易スライス式

配列やスライスの一部を抜き出し、新しいスライスを生成することができる。

```go
array := []int{1, 2, 3, 4, 5}
slice := array[0:2]
fmt.Println(slice) // [1 2]
```

### 簡易スライス式のパターン

| array := []int{1, 2, 3, 4, 5} | result |
| - | - |
| array[0:2] | [1, 2] |
| array[2:] | [3, 4, 5] |
| array[:4] | [1, 2, 3, 4] |
| array[:] | [1, 2, 3, 4, 5] |

### 文字列と簡易スライス式

簡易スライス式は文字列にも適用することができる。

```go
slice := "ABCDE"[1:3]
fmt.Println(slice)    // BC
```

ただし、簡易スライス式の単位は文字単位ではなくバイト単位であるため、マルチバイト文字を使用する場合には適していない。

***

* [Back to previous "13. 参照型"](./reference.md)
* [Back to index page](../README.md)
