# 定数

## 定数の定義

```go
const X = 1
```

Go では `const` を使用して定数を定義できる。

`var` による変数定義と同様に `()` で囲むことでまとめて定義をすることが可能。

<!-- markdownlint-disable MD010 -->

```go
const (
	FOO = 1
	BAR = 2
	BAZ = 3
)
```

<!-- markdownlint-enable MD010 -->

## 値の省略

定数の値を省略することができる。省略した場合ははじめに定義された定数の値が、そのまま以降の定数に割り当てられる。

<!-- markdownlint-disable MD010 -->

```go
const (
	A = 1
	B         // 1
	C         // 1
	D = "hoge"
	F         // "hoge"
	E         // "hoge"
)
```

<!-- markdownlint-enable MD010 -->

## 定数の型

Go の定数では以下の値を定義できる。

* 論理値
* 整数
* 浮動小数店
* 複素数
* ルーン
* 文字列

Go の定数は `型なし定数` と `型あり定数` の 2 つに別れており、定数の値に型を与える場合は型を明示する。

<!-- markdownlint-disable MD010 -->

```go
const (
	// どちらでも書くことができるが後者のが見やすい
	FOO int64 = 1
	BAR = int64(1)
)
```

<!-- markdownlint-enable MD010 -->

## iota

Go には列挙型(enum) の機能はないが定義済み識別子 `iota` を定数定義と合わせて使用することで近い振る舞いを実現できる。

<!-- markdownlint-disable MD010 -->

```go
const (
	A = iota  // 0
	B         // 1
	C         // 2
)
```

<!-- markdownlint-enable MD010 -->

`iota` は参照の有無に関わらず `const` ブロックの中で定数が定義されるたびに 1 ずつ増分する。

<!-- markdownlint-disable MD010 -->

```go
const (
	A = 999
	B = iota  // 1
	C         // 2
	D = 998
	E = iota  // 4
)
```

<!-- markdownlint-enable MD010 -->

***

* [Back to previous "3. 関数"](./function.md)
* [Back to index page](../README.md)