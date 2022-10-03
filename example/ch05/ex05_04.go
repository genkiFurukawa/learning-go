package main

import (
	"io"
	"log"
	"os"
)

// deferのサンプル
// catの簡易版
func main() {
	if len(os.Args) < 2 { // ファイル名が指定されているかチェック
		log.Fatal("ファイルが指定されていません")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// deferを使うと関数の終了時まで実行が延期される
	defer f.Close() // 後始末のコード

	data := make([]byte, 2048) // バイトのスライスを作成
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
