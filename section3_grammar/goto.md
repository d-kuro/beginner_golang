# goto

## goto 文

関数内の任意の位置へジャンプするための文。C と同様に Go の `goto` 文はラベルと合わせて使用する。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	fmt.Println("A")
	goto GOTO
	fmt.Println("B")
GOTO:
	fmt.Println("C")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
A
C
```

`goto` 文でジャンプできるのは関数内のみ、関数の間をジャンプすることはできない。また、`for` などが構成するブロックの「内側」にジャンプすることはできない。

変数定義を飛び越える `goto` 文もエラーとなる。

<!-- markdownlint-disable MD010 -->

```go
	goto GOTO
	n := 1 // エラー
GOTO:
	fmt.Println("hoge")
```

<!-- markdownlint-enable MD010 -->

`goto` 文が活かせる場面は、ネストしたループの内部から一気に抜けるなどが考えられるが、`goto` 文を使用せずに同様の処理を記述できるため殆どの場面で `goto` 文は不要。

## ラベル付き文

ラベルと `for` などの構造を持った文と組み合わせることで、複雑な処理フローをわかりやすくすることができる。

ラベル指定の `break` を使用するとラベルの処理を丸ごと抜けることができる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	fmt.Println("start")
LOOP:
	for {
		for {
			fmt.Println("hoge")
			break LOOP
		}
		fmt.Println("not processed")
	}
	fmt.Println("end")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
start
hoge
end
```

`continue` にもラベルを指定することができる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
LOOP:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue LOOP
			}
			fmt.Printf("i : %d, j : %d\n", i, j)
		}
	}
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
i : 1, j : 1
i : 2, j : 1
i : 3, j : 1
```

この場合だと、`continue LOOP` で外側のループの次の処理に移る。

***

* [Back to previous "8. switch"](./switch.md)
* [Back to index page](../README.md)
