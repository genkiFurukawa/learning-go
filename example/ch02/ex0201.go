package main

import (
	"fmt"
	"math/cmplx"
)

// 定数宣言
const hoge int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)

func main() {
	fmt.Println(hoge)
	fmt.Println(idKey, nameKey)

	// 複素数
	x := complex(2.3, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(real(y))
	fmt.Println(cmplx.Abs(x))

	// 変数の宣言方法
	var x1 = 10
	var x2 int = 20
	var x3, x4 = 10, "hello" // カンマ区切りで変数宣言できる
	fmt.Println(x1, x2, x3, x4)

	// = で代入できる
	x3 = 0o3

	// まとめて宣言もできる
	var (
		x5     int
		x6     = 20
		x7, x8 string
	)
	fmt.Println(x5, x6, x7, x8)

	// 関数内では:= を使って変数宣言もできる
	x09 := 20 // var x09 = 20 と等価
	fmt.Println(x09)

	// GOの言語仕様
	// := の左辺に新しい変数が一つでもあれば、既存の変数にも値を代入できる
	x09, x10 := 30, "hoge"
	fmt.Println(x09, x10)
}
