# テストを書く

## 構成

```test
.
├── app.go
├── app_test.go
├── language
│   ├── language.go
│   └── language_test.go
└── main.go
```

## language_test.go

<!-- markdownlint-disable MD010 -->

```go
package language

import "testing"

func TestEnglish(t *testing.T) {
	except := "English"
	actual := English()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}

func TestJapanese(t *testing.T) {
	except := "Japanese"
	actual := Japanese()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}
```

<!-- markdownlint-enble MD010 -->

## app_test.go

<!-- markdownlint-disable MD010 -->

```go
package main

import "testing"

func TestAppName(t *testing.T) {
	except := "language application"
	actual := AppName()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}
```

<!-- markdownlint-enble MD010 -->

## 実行結果

```bash
$ go test -v
=== RUN   TestAppName
--- PASS: TestAppName (0.00s)
PASS
ok      .../beginner_golang/section2_app/test       0.007s

$ go test -v ./language
=== RUN   TestEnglish
--- PASS: TestEnglish (0.00s)
=== RUN   TestJapanese
--- PASS: TestJapanese (0.00s)
PASS
ok      .../beginner_golang/section2_app/test/language      0.055s
```

## 解説

テストには標準パッケージの `testing` をインポートする。関数 `TestEnglish` のように名前が `Test` で始まる関数がテストの単位を表している。

テストの実行には `go test` コマンドを使用する。

```bash
# go test [package dir]
$ go test ./language
ok      .../beginner_golang/section2_app/test/language      0.008s
```

`-v` オプションを使用すると個々のテストに対しての詳細が確認できる。

```bash
$ go test -v ./language
=== RUN   TestEnglish
--- PASS: TestEnglish (0.00s)
=== RUN   TestJapanese
--- PASS: TestJapanese (0.00s)
PASS
ok      .../beginner_golang/section2_app/test/language      0.055s
```

テストが失敗した時は以下のように出力される。

```bash
$ go test -v ./language
=== RUN   TestEnglish
--- PASS: TestEnglish (0.00s)
=== RUN   TestJapanese
--- FAIL: TestJapanese (0.00s)
        language_test.go:19: Japanese != Japanese_faild
FAIL
FAIL    .../beginner_golang/section2_app/test/language      0.007s
```

***

* [Back to previous "2. main パッケージの分割"](./split_main.md)
* [Back to index page](../README.md)
