# 関数

## 関数定義

<!-- markdownlint-disable MD010 -->

```go
func addition(x, y int) int {
	return x + y
}
```

<!-- markdownlint-enable MD010 -->

関数は以下のフォーマットで定義を行う。

<!-- markdownlint-disable MD010 -->

```go
func [関数名]([引数]) [戻り値の型] {
	[処理]
}
```

<!-- markdownlint-enable MD010 -->

## 引数定義

```go
func addition(x, y int) int
```

上記の関数には `x` と `y` の 2 つの引数が定義されているが、型である `int` は 1 つだけ指定されている。このように複数の同じ型の引数を定義する場合は、型指定を 1 箇所でまとめることができる。

## 戻り値のない関数

戻り値の型定義を省略すると戻り値を持たない関数を定義することができる。Go には `void` 型は存在しない。

<!-- markdownlint-disable MD010 -->

```go
func hello() {
	fmt.Println("Hello!")
}
```

<!-- markdownlint-enable MD010 -->

## 複数の戻り値

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	q, r := division(19, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", q, r)
}

func division(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}
```

<!-- markdownlint-enable MD010 -->

Go の関数は複数の戻り値を返すことができる。`(int, int)` のように `()` で囲んで列挙を行う。`()` の省略は不可能。

## 戻り値の破棄

Go では `_` を使用して戻り値の一部を破棄することができる。また、関数の引数も `_` を使用して無視することが可能。

```go
q, _ := division(15, 7)
```

## エラー処理

Go には例外機構が存在しないため、任意の関数を呼び出した後にその処理が成功したかどうかを戻り値の一部で示す。

<!-- markdownlint-disable MD010 -->

```go
result, err := doSomething()
if (err != nil) {
	// エラー処理
}
```

<!-- markdownlint-enable MD010 -->

## 戻り値を表す変数

戻り値の型に変数名を書くことで関数内の変数定義を短縮することができる。

<!-- markdownlint-disable MD010 -->

```go
func division(x, y int) (q, r int) {
	q = x / y
	r = x % y
	return
}
```

<!-- markdownlint-enable MD010 -->

## 無名関数

関数を値として表現する。Go では関数の引数に関数をとったり、関数を返す関数を記述することができる。

```go
f := func(x, y int) int { return x + y }
fmt.Println(f(2, 3)) // 5
```

この場合の変数 `f` の型は `func(int, int) int` 型となる。

### 関数を返す関数

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func returnFunc() func() {
	return func() {
		fmt.Println("Hello!")
	}
}

func main() {
	returnFunc()()
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
Hello!
```

関数 `returnFunc` は戻り値に「引数も戻り値もない関数」を返す関数として定義されている。`returnFunc` が返す関数を暗黙的に代入して呼び出しているが、以下のように記述することで直接呼び出すこともできる。

```go
returnFunc()() // Hello!
```

### 関数を引数にとる関数

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func callFunc(f func()) {
	f()
}

func main() {
	callFunc(func() {
		fmt.Println("Hello!")
	})
}
```

<!-- markdownlint-enable MD010 -->

無名関数を引数として渡すことで、`callFunc` 内で関数が呼び出される。

## クロージャ

Go の無名関数はクロージャ。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func later() func(string) string {
	// 1 つ前の文字列を保持する
	var store string
	// 引数に文字列を取り文字列を返す関数を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()

	fmt.Println(f("foo"))
	fmt.Println(f("bar"))
	fmt.Println(f("baz"))
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
    # 空文字
foo
bar
```

関数 `later` の中で定義されている変数 `store` はローカル変数。通常、関数のローカル変数は関数の実行が完了した後に破棄される。

だが、Go のコンパイラはクロージャ内からのローカル変数の参照を検出すると、ローカル変数とは別にクロージャと紐づいた変数として処理を行う。

<!-- markdownlint-disable MD010 -->

```go
var store string
return func(next string) string {
	s := store
	store = next
	return s
}
```

<!-- markdownlint-enable MD010 -->

クロージャに捕捉された変数の領域はクロージャが参照され続ける限り破棄されることはない。

***

* [Go to next "4. 定数"](./const.md)
* [Back to previous "2. 配列"](./array.md)
* [Back to index page](../README.md)
