package main

import "fmt"

func main() {
	// ポインタのサンプル
	x := 10
	pointerToX := &x
	fmt.Println("アドレス: ", pointerToX)
	fmt.Println("デリファレンス: ", *pointerToX)
	z := *pointerToX * 5
	fmt.Println(z)

	var y *int
	fmt.Println("y == nil: ", y == nil)
	// nilの値をデリファレンスしようとするとpanicになる
	fmt.Println(*y)
}
