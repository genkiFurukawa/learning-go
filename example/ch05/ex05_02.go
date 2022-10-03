package main

import (
	"fmt"
	"strconv"
)

// MOTE:
// 関数は値
// 関数の型はキーワードfunc、および引数と戻り値の型によって決まる
// この組み合わせを関数のシグネチャと呼ぶ
// 2つの関数の引数と戻り値が同じであれば、両者のシグネチャは同じであると言える

// 同じシグネチャを持った関数
func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

//
var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}

	// fmt.Println(expressions)

	for _, exexpression := range expressions {
		if len(exexpression) != 3 {
			fmt.Println(exexpression, " -- 不正な式です")
			continue

		}
		// 文字列を数値に変換するメソッド
		p1, err := strconv.Atoi(exexpression[0]) // 1番目の非演算子のチェック
		if err != nil {
			fmt.Println(exexpression, " -- ", err)
			continue
		}

		op := exexpression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println(exexpression, " -- ", "定義されていない演算子です", op)
			continue
		}

		p2, err := strconv.Atoi(exexpression[2]) // 3番目の非演算子のチェック
		if err != nil {
			fmt.Println(exexpression, " -- ", err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(exexpression, " -> ", result)
	}
}
