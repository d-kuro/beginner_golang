# init

## init とは

パッケージの初期化を目的とした特殊な関数。引数をとらず、戻り値も返さない。制御構文でも定義済み関数でもなく、開発者が任意で定義することができる。

`init` はプログラムのメインルーチンが実行される前の段階で、初期化処理を実行するために利用できる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
init
main
```

`init` はパッケージ内に同名の関数を複数定義することができる。それぞれの実行順序はソースコードに出現した順序となる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("main")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
init 1
init 2
main
```

***

* [Go to next "13. 参照型"](./reference.md)
* [Back to previous "11. panic"](./panic.md)
* [Back to index page](../README.md)
