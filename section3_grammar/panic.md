# panic

## panic とは

`panic` は定義済み関数として定義されており、実行するとランタイムパニック (ランタイムエラー) が発生し、実行中の関数は中断される。

`panic` はプログラムにおいてこれ以上処理を継続しようがない状態を表すために使用され、一般的なエラー処理などで使用するものではない。`panic` を使用する場合は、どのような原因でランタイムが停止したのかについての情報を付与すると良い。

`panic` はランタイムをエラー終了させるが、中断時までに登録された `defer` は全て実行される。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	defer fmt.Println("run defer")
	panic("runtime error")
	fmt.Println("hoge")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
run defer
panic: runtime error

goroutine 1 [running]:
main.main()
        .../beginner_golang/main.go:7 +0x95
exit status 2
```

## recover

`panic` によって発生したランタイムパニックによるプログラムの中断を回復させる定義済み関数。

`recover` は `defer` と組み合わせて使うのが原則。`recover` は `interface{}` 型の値を戻し、その値が `nil` でなければ `panic` が実行されたと判断することができる。

<!-- markdownlint-disable MD010 -->

```go
package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("panic")
}
```

<!-- markdownlint-enable MD010 -->

```bash
$ go run main.go
panic
```

変数 `x` には `panic` に渡された `interface{}` が格納される。

`panic` と `recover` を組み合わせることによって、例外処理のような動作を実現することができるが、関数をまたぐ `goto` のように強力すぎる機能であるため使用は原則控えた方が良い。Go における任意の関数が `panic` を起こす可能性があることを前提にデザインするのは避ける。

***

* [Go to next "12. init"](./init.md)
* [Back to previous "10. defer"](./defer.md)
* [Back to index page](../README.md)
