package main

import "fmt"

// selectを使ってデッドロックを防ぐサンプル
// ※mainもゴルーチンとして起動されている
func main() {
	// デッドロックする2つのゴルーチン
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println(">> 無名関数内: ")
		v := 1
		ch1 <- v // ch1に書き込めない限りここで待たされる
		v2 := <-ch2
		fmt.Println("<< 無名関数内: ", v, " ", v2)
	}()

	fmt.Println(">> mainの最後: ")
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(">> mainの最後: ", v, " ", v2)
}
