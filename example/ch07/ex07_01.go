package main

import (
	"fmt"
	"time"
)

// 型の定義
type Score int                    // intを基底型としてScoreを定義
type Cpnverter func(string) Score // stringを引数、Scoreを返り値とする関数型Converterを定義
type TeamScores map[string]Score  // stringをScoreにマップするmap、TeamScoresを定義

// 型Personを定義
type Person struct {
	LastName  string // 姓
	FirstName string // 名
	Age       int    // 年齢
}

// 型Personにメソッドを定義
// funcとメソッド名の間にレシーバが追加されている
// レシーバは型の先頭の小文字がよく使われる
func (p Person) String() string {
	return fmt.Sprintln(p.LastName + " " + p.FirstName)
}

type Counter struct {
	total       int       // 合計
	lastUpdated time.Time // 最終更新時刻
}

// ポイントレシーバ
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

// 値がコピーして渡されるので、引数で渡したcは更新されない
func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("NG: ", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("OK: ", c.String())
}

func main() {
	p := Person{
		LastName:  "LastName",
		FirstName: "FirstName",
		Age:       3,
	}
	fmt.Print(p.String())
	//
	fmt.Println("==")
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	//
	var c2 Counter
	fmt.Println("==")
	fmt.Println("funcで値を渡したときとポインタを渡した時の挙動")
	fmt.Println(c2.String())
	doUpdateWrong(c2)
	fmt.Println("main", c2.String())
	doUpdateRight(&c2)
	fmt.Println("main", c2.String())
}
