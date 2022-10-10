package main

import "fmt"

// 埋め込み
type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

func main() {
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println("o.X:", o.X)
	fmt.Println("o.Inner.X:", o.Inner.X)
}
