package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	if x > 5 {

		fmt.Println(x) // 10
		// 変数のシャドーイング
		a, x := 20, 5
		fmt.Println(a, x) // 20 5
	}
	fmt.Println(x) // 10

	// fmt := "おお"
	// 局所変数にfmtに備わっていないメソッドが使われようとしているので怒られる
	// fmt.Println(fmt)

	n := rand.Intn(10)
	if n == 0 {
		fmt.Println(n)
	} else if n > 5 {
		fmt.Println("5より大きい", n)
	} else {
		fmt.Println("それ以外", n)
	}

	fmt.Println("==")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("==")
	i := 1
	for i < 5 {
		fmt.Println(i)
		i = i * 2
	}

	fmt.Println("==")
	// for-rangeループ
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	fmt.Println("==")
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// マップのイテレーション
	fmt.Println("==")
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	// 毎回順番が異なる
	for i := 0; i < 3; i++ {
		fmt.Println("==")
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	// ブランクswitch文
	fmt.Println("==")
	words := []string{"hi", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 3:
			fmt.Println(word, "3文字以下")
		default:
			fmt.Println(word, "その他")
		}
	}
}
