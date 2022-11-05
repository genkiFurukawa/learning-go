package main

import "fmt"

// デッドロック状態になった時にGoのランタイムがプログラムを強制終了するサンプル
// ※mainもゴルーチンとして起動されている
func main() {
	// デッドロックする2つのゴルーチン
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println(">> 軽量スレッド")
		v := 1
		fmt.Println(">> ch1に書き込み")
		ch1 <- v // ch1に書き込めない限りここで待たされる
		fmt.Println("<< ch1に書き込み")
		v2 := <-ch2
		fmt.Println(v, v2)
		fmt.Println("<< 軽量スレッド")
	}()

	fmt.Println(">> 処理実行")
	v := 2
	fmt.Println(">> ch2に書き込み")
	ch2 <- v // ch2に書き込めない限りここで待たされる
	fmt.Println("<< ch2に書き込み")
	v2 := <-ch1
	fmt.Println(v, v2)
	fmt.Println("<< 処理実行")
}
