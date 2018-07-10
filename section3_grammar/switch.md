# switch

## 式による switch

`if` と同様に式の前に任意で簡易文を書くことができる。`switch` のブロックの内部には任意個の `case` を書くことができ、全ての `case` に分岐しなかった場合に実行される `default` を任意で 1 つだけ書くことができる。

<!-- markdownlint-disable MD010 -->

```go
switch [簡易文;] [式] {

}
```

<!-- markdownlint-enable MD010 -->

`case` には複数の値をカンマで区切って並べることができる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	switch n := "foo"; n {
	case "foo", "bar":
		fmt.Println("foo or bar")
	case "baz":
		fmt.Println("baz")
	default:
		fmt.Println("unknown")
	}
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
foo or bar
```

他の言語における `switch` 文では各 `case` の処理に明示的に `break` が置かれない限り次の `case` に処理が継続されていた(フォークスルー) が、Go の `switch` では `case` 実行が終了すると同時に `switch` の処理が完了する。

意図的にフォークスルーしたい場合は `fallthrough` を使用する。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	switch n := "A"; n {
	case "A":
		n += "B"
		fallthrough
	case "B":
		n += "C"
		fallthrough
	default:
		n += "D"
		fmt.Println(n)
	}
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
ABCD
```

## 式を伴う case

`case` に `bool` を返す任意の式を書くことができる。この場合 `switch` に与える式は省略することができる。

省略された `switch` の式は内部的に `true` が置かれているものとみなす。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	n := 4
	switch {
	case 0 < n && n < 3:
		fmt.Println("0 < n < 3")
	case 3 < n && n < 6:
		fmt.Println("3 < n < 6")
	}
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
3 < n < 6
```

定数を列挙する `case` と、式による `case` の混在はエラーになるので注意する。

<!-- markdownlint-disable MD010 -->

```go
// コンパイルエラー
switch {
case 0, 1, 2, 3:
	fmt.Println("0 < n < 3")
case 3 < n && n < 6:
	fmt.Println("3 < n < 6")
}
```

<!-- markdownlint-enable MD010 -->

***

* [Back to previous "7. for"](./for.md)
* [Back to index page](../README.md)
