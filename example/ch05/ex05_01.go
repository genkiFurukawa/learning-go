package main

import (
	"errors"
	"fmt"
)

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("0での除算はできません")
	}
	return numerator / denominator, numerator % denominator, nil
}

// 名前付き引数を実現するのであれば、構造体を使う
type MyFuncOps struct {
	FirstName string
	LatsName  string
	Age       int
}

func MyFunc(myFuncOps MyFuncOps) error {
	fmt.Println(myFuncOps)
	return nil
}

// 可変長引数とスライス
func addTo(base int, vals ...int) []int {
	// 通常のスライスを同じ扱いができる
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	result := div(5, 2)
	fmt.Println(result)
	//
	MyFunc(MyFuncOps{})
	MyFunc(MyFuncOps{
		FirstName: "FirstName",
		LatsName:  "LatsName",
		Age:       10,
	})
	//
	vals := addTo(1, 2, 3, 4, 5)
	fmt.Println(vals)
	//
	res, remainder, err := divAndRemainder(0, 0)
	fmt.Println(res, remainder, err)
	// 戻り値を無視することができる
	res, _, err = divAndRemainder(5, 3)
	fmt.Println(res, err)
}
