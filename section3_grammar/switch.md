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

## 型による switch

### 型アサーション

`interface{}` 型の変数について、実態の型がどのようなものなのかを動的にチェックする。

```go
var x interface{} = 3
i := x.(int)     // 変数 x は int
f := x.(float64) // 実行時にエラー
```

上記の書き方だと、アサーションに失敗した場合エラーが発生しプログラムが停止するため明確に型を推測できる場面以外で使用しない。

以下の書き方の場合だと、1 番目の戻り値には値、2 番目の戻り値にはアサーションの成功可否が `bool` で代入される。アサーションに失敗した場合の 1 番目の戻り値は初期値になる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	var y interface{} = 3.14
	i, isInt := y.(int)
	f, isFloat := y.(float64)

	fmt.Println(i, isInt)
	fmt.Println(f, isFloat)
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
0 false
3.14 true
```

### switch で型アサーションを使用する

`switch [変数(interface{})].(type)` という記述をする。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	var x interface{} = 3
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("int or uint")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
int or uint
```

値を取得する必要があるならば以下のような書き方もできる。

<!-- markdownlint-disable MD010 -->

```go
var x interface{} = 3
switch i := x.(type) {
case bool:
	fmt.Println("bool:", i)
case int, uint:
	fmt.Println("int or uint:", i)
case string:
	fmt.Println("string:", i)
default:
	fmt.Println("unknown:", i)
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
int or uint: 3
```

`case int, uint:` のように `case` の中で型を複数列挙した場合は `interface{}` 型の変数として `case` の中で振る舞うので注意する。

***

* [Go to next "9. goto"](./goto.md)
* [Back to previous "7. for"](./for.md)
* [Back to index page](../README.md)
