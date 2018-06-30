# 環境構築

## はじめに

このページでは Go 言語が動作する環境を構築する。

以下の環境を使用して構築した際の手順を記載する。

```text
OS : macOS High Sierra version 10.13.5
```

## goenv

goのバージョン管理ツール

Homebrew でインストールする。

<!-- markdownlint-disable MD014 -->

```bash
$ brew install goenv
```

<!-- markdownlint-enable MD014 -->

`~/.bash_profile` に以下を記載する。

```~/.bash_profile
export PATH="$HOME/.goenv/bin:$PATH"
eval "$(goenv init -)"
```

`source` コマンドで反映させる。

<!-- markdownlint-disable MD014 -->

```bash
$ source ~/.bash_profile
```

<!-- markdownlint-enable MD014 -->

確認。

```bash
$ goenv -v
goenv 1.0.0
```

## Golang のインストール

ダウンロード可能な Golang バージョンを見る。

```bash
$ goenv install -l
Available versions:
  1.2.2
  1.3.0
  :
  1.10.0
  1.10beta2
  1.10rc1
  1.10rc2
  1.10.1
```

ここでは `1.10.1` をインストールする。

<!-- markdownlint-disable MD014 -->

```bash
$ goenv install 1.10.1
```

<!-- markdownlint-enable MD014 -->

バージョンを切り替える。

<!-- markdownlint-disable MD014 -->

```bash
$ goenv global 1.10.1
$ goenv rehash
```

<!-- markdownlint-enable MD014 -->

Golang のバージョンを確認。

```bash
$ go version
go version go1.10.1 darwin/amd64
```

## GOPATHの設定

Go は外部のライブラリが格納されているディレクトリを知るために、環境変数 `GOPATH` を利用する。

`GOPATH` に設定するディレクトリを作成。

<!-- markdownlint-disable MD014 -->

```bash
$ mkdir ~/go
```

<!-- markdownlint-enable MD014 -->

`~/.bash_profile` に以下を記載する。

```~/.bash_profile
export GOPATH=$HOME/go
```

`source` コマンドで反映させる。

<!-- markdownlint-disable MD014 -->

```bash
$ source ~/.bash_profile
```

<!-- markdownlint-enable MD014 -->

確認。

```bash
$ echo $GOPATH
/Users/hoge.user/go
```
