# 参照型

## 参照型とは

Go には特殊なデータ構造を備えた「参照型」という型が定義されている。標準で `slice`、`map`、`channel` の 3 つが定義されている。

## make

参照型の生成には、組み込み関数 `make` を使用する。スライス、マップ、チャネルのデータ構造は `make` によって生成されるため、関数の呼び出し方にもバリエーションが存在する.

| make | 型 `T` | 意味 |
| - | - | - |
| `make(T, n)` | スライス | 要素数と容量が `n` である `T` 型のスライスを生成 |
| `make(T, n, m)` | スライス | 要素数が `n` で容量が `m` である `T` 型のスライスを生成 |
| `make(T)` | マップ | `T` 型のマップを生成 |
| `make(T, n)` | マップ | `T` 型のマップを要素数 `n` をヒントにして生成 |
| `make(T)` | チャネル | バッファのない `T` 型のチャネルを生成 |
| `make(T, n)` | チャネル | バッファサイズ `n` の `T` 型のチャネルを生成 |

***

* [Go to next "14. スライス"](./slice.md)
* [Back to previous "12. init"](./init.md)
* [Back to index page](../README.md)
