package main

import (
	"fmt"
	"package_example/formatter"
	"package_example/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
