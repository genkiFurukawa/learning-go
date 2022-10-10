package main

import (
	"errors"
	"fmt"
	"os"
)

// 被除数と除数を受け取って、商と余りとエラーを返す
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator id 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

// 偶数なら2倍して返す
func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%dは偶数ではありません", i)
	}
	return i * 2, nil
}

func main() {
	numerator, denominator := 20, 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)

	i := 19
	double, err := doubleEven(i)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(i, "の2倍: ", double)
}
