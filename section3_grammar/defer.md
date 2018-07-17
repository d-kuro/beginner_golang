# defer

## defer とは

関数の終了時に実行される式を登録することができる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	defer fmt.Println("defer")
	fmt.Println("hoge")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
hoge
defer
```

`defer` に登録できる文は関数呼び出しの形式の式に限られる。

`defer` はいくつでも指定可能だが、LIFO で後に登録された式が先に評価される。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("hoge")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
hoge
3
2
1
```

`defer` はリソースの解放処理によく用いられる。

<!-- markdownlint-disable MD010 -->

```go
file, err := os.Open("/path/to/file")
if err != nil {
	return
}
defer file.Close()
```

<!-- markdownlint-enable MD010 -->

`defer` に複数の処理と登録する場合には無名関数を利用することも可能だが、関数呼び出しの形式にしなければならないので `()` が必要になる。

<!-- markdownlint-disable MD010 -->

```go
defer func() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}()
```

<!-- markdownlint-enable MD010 -->

***

* [Back to previous "9. goto"](./goto.md)
* [Back to index page](../README.md)
