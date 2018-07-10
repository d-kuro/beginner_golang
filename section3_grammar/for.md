# for

## 何も書かない for

何も書かない `for` は無限ループを構成する。

<!-- markdownlint-disable MD010 -->

```go
for {
	fmt.Println("infinite loop")
}
```

<!-- markdownlint-enable MD010 -->

## break

ループを中断させるための機能として `break` 文が用意されている。`for` 文が複数ネストしている場合は最も近い位置の `for` ループを中断させる。

## 条件式付き for

Java などにおける `while` に相当する。

<!-- markdownlint-disable MD010 -->

```go
i := 0
for i < 100 {
	i++
}
```

<!-- markdownlint-enable MD010 -->

## 古典的 for

一般的な `for`。

<!-- markdownlint-disable MD010 -->

```go
for i := 0; i < 100; i++ {
	fmt.Println(i)
}
```

<!-- markdownlint-enable MD010 -->

## continue

ループ内で残処理をスキップして次のループ処理へ継続する。`for` 文が複数ネストしている場合は最も近い位置の `for` ループが対象になる。

## 範囲節による for

予約語 `range` と任意の式を組み合わせて定義する。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	array := [3]string{"foo", "bar", "baz"}
	for i, s := range array {
		fmt.Printf("index = %d, value = %s\n", i, s)
	}
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
index = 0, value = foo
index = 1, value = bar
index = 2, value = baz
```

### range に適用できる型

| 型 | 反復値1 | 反復値2 |
| - | - | - |
| 配列型 | index | value |
| 配列へのポインタ | index | value |
| スライス | index | value |
| 文字列 | index | rune |
| マップ | key | value |
| チャネル | チャネルの値 | - |

文字列に対する `range` は UTF-8 でエンコードされた文字列のコードポイントごとに反復されるので、「 n 番目」 の文字という意味ではない点に注意する。

***

* [Go to next "8. switch"](./switch.md)
* [Back to previous "6. if"](./if.md)
* [Back to index page](../README.md)
