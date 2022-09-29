package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(100)
	for a < 100 {
		fmt.Println(a)
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}

done:
	fmt.Println("ループを抜けた時に実行")
	fmt.Println(a)
}
