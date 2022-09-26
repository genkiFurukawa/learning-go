package main

import (
	"fmt"
)

func main() {
	// 配列の宣言方法
	x := [3]int{10, 20, 30}
	var y [3]int

	fmt.Println(x)
	fmt.Println(y)

	// 代入の仕方
	y = [...]int{10, 20, 30}
	fmt.Println(y)

	// スライス
	z := []int{10, 20, 30}
	fmt.Println(z)
	fmt.Println(z[0])
	fmt.Println(z[1])

	// 再編を書かないとコンパイルエラーとなる
	z = append(z, 3)
	fmt.Println(z[3])

	// capで連続して確保されている領域のサイズを確認ができる
	a := append(z, 5)
	fmt.Println(a, len(a), cap(a))

	// makeを使うとあらかじめ領域を確保できる
	// 型、長さ、キャパシティ(cap)の順で指定される
	b := make([]int, 128, 1024)
	fmt.Println(len(b), cap(b)) // 128 1024

	// 宣言方法でnilスライスに変わる
	var data01 []int // nilスライスを生成
	fmt.Println(data01)
	fmt.Println(data01 == nil)
	data02 := []int{} //長さ0でスライスを生成
	fmt.Println(data02)
	fmt.Println(data02 == nil)

}
