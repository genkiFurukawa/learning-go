package main

import "fmt"

func main() {
	fmt.Println("==")
	var s string = "Hello 🌄"
	var s2 string = s[8:]
	fmt.Println(len(s)) // 7ではなく12が出る
	fmt.Println(s)
	fmt.Println(s2) // 文字のバイトの途中をコピーしただけなので、文字化けする

	// マップ
	totalWins := map[string]int{} // string -> int マップを要素なしで初期化
	fmt.Println("==")
	fmt.Println("abc:", totalWins["abc"])
	totalWins["abc"] = 3
	fmt.Println(totalWins)
	fmt.Println(len(totalWins))
	fmt.Println(totalWins == nil)

	// 初期値あり
	teams := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("==")
	fmt.Println(teams)

	// カンマokイディオム
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println("==")
	fmt.Println(v, ok)
	v, ok = m["hoge"]
	fmt.Println(v, ok)
	// 削除
	delete(m, "hello")
	v, ok = m["hello"]
	fmt.Println(v, ok)

	var nilMap map[string]int //nilで初期化
	fmt.Println("==")
	fmt.Println(nilMap)
	fmt.Println(len(nilMap))
	fmt.Println(nilMap == nil)
	// nilMap["abc"] = 3 // パニックになる

	// 構造体
	type person struct {
		name string
	}

	bob := person{}
	fmt.Println("==")
	fmt.Println(bob)

	alice := person{
		"alice",
	}
	fmt.Println(alice)

	julia := person{
		name: "julia",
	}
	fmt.Println(alice)
	fmt.Println(julia)

	// 無名構造体
	fmt.Println("==")
	pet := struct {
		name string
		kind string
	}{
		name: "ぽち",
		kind: "dog",
	}
	fmt.Println(pet)
}
