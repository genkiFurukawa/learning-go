package main

import "fmt"

func main() {
	fmt.Println("==")
	var s string = "Hello ๐"
	var s2 string = s[8:]
	fmt.Println(len(s)) // 7ใงใฏใชใ12ใๅบใ
	fmt.Println(s)
	fmt.Println(s2) // ๆๅญใฎใใคใใฎ้ไธญใใณใใผใใใ ใใชใฎใงใๆๅญๅใใใ

	// ใใใ
	totalWins := map[string]int{} // string -> int ใใใใ่ฆ็ด ใชใใงๅๆๅ
	fmt.Println("==")
	fmt.Println("abc:", totalWins["abc"])
	totalWins["abc"] = 3
	fmt.Println(totalWins)
	fmt.Println(len(totalWins))
	fmt.Println(totalWins == nil)

	// ๅๆๅคใใ
	teams := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("==")
	fmt.Println(teams)

	// ใซใณใokใคใใฃใชใ 
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println("==")
	fmt.Println(v, ok)
	v, ok = m["hoge"]
	fmt.Println(v, ok)
	// ๅ้ค
	delete(m, "hello")
	v, ok = m["hello"]
	fmt.Println(v, ok)

	var nilMap map[string]int //nilใงๅๆๅ
	fmt.Println("==")
	fmt.Println(nilMap)
	fmt.Println(len(nilMap))
	fmt.Println(nilMap == nil)
	// nilMap["abc"] = 3 // ใใใใฏใซใชใ

	// ๆง้ ไฝ
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

	// ็กๅๆง้ ไฝ
	fmt.Println("==")
	pet := struct {
		name string
		kind string
	}{
		name: "ใฝใก",
		kind: "dog",
	}
	fmt.Println(pet)
}
