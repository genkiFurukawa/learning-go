package main

import "fmt"

// % go run ex10_03.go
// 24 24 24 24 28 32 36 40 40 40
// % go run ex10_03.go
// 24 24 24 24 32 36 36 40 40 40 %
func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch1 := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch1 <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch1, " ")
	}
	fmt.Println("")

	// ゴルーチン起動時に決めた値を渡す
	// ①シャドーイング
	ch2 := make(chan int, len(a))
	for _, v := range a {
		v := v
		go func() {
			ch2 <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch2, " ")
	}
	fmt.Println("")

	// ②ゴルーチンに引数を渡す
	ch3 := make(chan int, len(a))
	for _, v := range a {
		go func(val int) {
			ch3 <- val * 2
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch3, " ")
	}
	fmt.Println("")
}
