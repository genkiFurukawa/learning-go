package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("無名関数の中で", j, "を出力")
		}(i)
	}
	// 関数引数のサンプル
	fmt.Println("==")
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("●初期データ")
	fmt.Println(people)
	// LastNameでソート
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("●LastNameでソート")
	fmt.Println(people)
	// Ageでソート
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("●Ageでソート")
	fmt.Println(people)
	fmt.Println("●ソート後")
	fmt.Println(people)

	// 関数から関数を返す
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i <= 5; i++ {
		fmt.Println(i, ":", twoBase(i), ",", threeBase(i))
	}
}
