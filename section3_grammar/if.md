# if

## 基本形

<!-- markdownlint-disable MD010 -->

```go
if x == 1 {
	// 略
} else if x == 2 {
	// 略
} else {
	// 略
}
```

<!-- markdownlint-enable MD010 -->

`if` や `else` が構成するブロックは常に `{}` で囲む必要があり、省略は不可能。

`if` の条件式は論理値を返す式である必要があり、0 は偽、0以外は真の値として利用する他の言語のようなことはできない。

## 簡易文付き if

Go では簡易文を伴う条件文を書くことができる。

簡易文とは式、代入文、暗黙の変数定義など複雑な構造を持たない単一の文のこと。

<!-- markdownlint-disable MD010 -->

```go
if x, y := 1, 2; x < y {
	// 略
}
```

<!-- markdownlint-enable MD010 -->

Go では `if` は暗黙的なブロックを構成する。`if` の外側で定義された変数と、内側で定義された変数ではスコープが異なるので、同一の名称だったとしても別々の変数となる。

<!-- markdownlint-disable MD010 -->

```go
x := 5
if x := 2; true {
	// x == 2
}
// x == 5
```

<!-- markdownlint-enable MD010 -->

簡易文はエラー処理に用いることが多い。

<!-- markdownlint-disable MD010 -->

```go
if result, err := doSomething(); err != nil {
	// エラー処理
}
```

<!-- markdownlint-enable MD010 -->

***

* [Go to next "7. for"](./for.md)
* [Back to previous "5. スコープ"](./scope.md)
* [Back to index page](../README.md)
