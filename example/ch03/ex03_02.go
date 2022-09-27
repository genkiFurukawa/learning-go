package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	x[1] = 20
	y[0] = 10 // xの値も書き変わる
	z[1] = 30 //  xの値も書き変わる
	fmt.Println("==")
	fmt.Println("x: ", x) // x:  [10 20 30 4]
	fmt.Println("y: ", y) // y:  [10 20]
	fmt.Println("z: ", z) // z:  [20 30 4]

	// キャパシティが共有される
	a := []int{1, 2, 3, 4}
	b := a[:2]
	fmt.Println(cap(a), cap(b))
	y = append(y, 30)
	fmt.Println(cap(a), cap(b))

	//
	d := make([]int, 0, 5)
	d = append(d, 1, 2, 3, 4)
	e := d[:2]                // [1, 2]
	f := d[2:]                // [3, 4]
	e = append(e, 30, 40, 50) // [1, 2, 30 ,40, 50]
	d = append(d, 60)         // [1, 2, 30, 40, 60]
	f = append(f, 70)         // [30, 40, 70]
	fmt.Println("==")
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("f", f)

	// フルスライス式
	// アペンドで上書きが生じないようにする
	g := make([]int, 0, 5)
	g = append(g, 1, 2, 3, 4)
	h := g[:2:2]              // [1, 2]
	i := g[2:4:4]             // [3, 4]
	h = append(h, 30, 40, 50) // [1, 2, 30 ,40, 50]
	g = append(g, 60)         // [1, 2, 3, 4, 60]
	i = append(i, 70)         // [3, 4, 70]
	fmt.Println("==")
	fmt.Println("g", g)
	fmt.Println("h", h)
	fmt.Println("i", i)

	// 配列からスライスの変換
	j := [...]int{5, 6, 7, 8} //配列を生成
	k := j[:2]
	l := j[2:]
	m := k[:] // 配列 -> スライスの変換
	j[0] = 10
	fmt.Println("==")
	fmt.Println("j", j)
	fmt.Println("k", k)
	fmt.Println("l", l)
	fmt.Println("m", m)
}
