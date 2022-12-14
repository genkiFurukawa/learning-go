# 初めての Go 言語

- [初めての Go 言語](https://www.oreilly.co.jp/books/9784814400041/)の写経リポジトリ

## 環境

```sh
$ go version
go version go1.19 darwin/arm64
```

## 各チャプターでの学び

### ch01

- `go run`を使うと Go のプログラムをインタプリタ言語のスクリプトのように扱うことができる
- `go mod init {モジュール名}`で`go.mod`が生成される
  - `go.mod`があると、ファイルを指定しなくてもそのディレクトリに置かれた.go ファイルを解析してコンパイルしてビルドしてくれる
  - `go mod tidy`は必要なサードパーティのダウンロードや不要になったファイルの削除を行ってくれる
- `go install`したものは`$GOPTH/bin`配下にダウンロードされる
- `go install github.com/rakyll/hey@latest`で補完が効かなかったので、パスを通した

  ```sh
  $ cat ~/.zshrc
  eval "$(anyenv init -)"
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:$GOBIN
  ```

### ch02

- 宣言されたが値が割り振られていない変数に対してはゼロ値が割り当てられる
- リテラル

  1. 整数リテラル

  - デフォルトは 10 進数
  - 0b をつけると二進数(**b**inaly)、0o をつけると 8 進数(**o**ctal)、0x をつけると 16 進数(he**x**adecimal)

  1. 不動小数点リテラル
  1. rune リテラル

  - 先頭と最後に`'`を置いたもの

  1. 文字列リテラル

- コードのフォーマットには`go fmt`か`goimports`を使う
  - [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- GO では変数名はキャメルケースを使う
- 名前の先頭文字が大文字か小文字がどうかでパッケージの外からアクセスできるかどうか決定される

### ch03

- Go で配列が直接使われることはない
  - `[3]int`とした時と`[4]int`と宣言した時では型が異なる
  - 可変長の配列であるスライスを使う
- Go 言語は「値呼び出し」の言語
- スライスが大きくならない可能性があるときは nil スライスでスライスを生成する
  - どうすればスライスを大きくする回数を減らせるのか考える
  ```go
  var data01 []int
  fmt.Println(data01 == nil) //trueになる
  ```

### ch04

- for-range ループではインデックスと値が取得できる。インデックスを使わないときは`_`を使う
- マップのイテレーションはマップの定義順となるので、注意
- for-range は値コピー
- goto 文がある。標準ライブラリのパッケージ strconv などで使われている
  - 稀に goto 文を使った方がわかりやすいときがあるので、そのときは goto を使う

## 参考

- [『Learning Go』を読んで、Go に入門している](https://blog.magnolia.tech/entry/2022/06/25/161525)
