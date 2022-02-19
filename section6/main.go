package main

import "fmt"

func main() {
	// 論理演算子
	fmt.Println("論理演算子")
	fmt.Println(true && false == true)
	fmt.Println(true && false == false)

	fmt.Println(true || false == true)
	fmt.Println(!true)
	fmt.Println("-------------------------------")

	// 比較演算子
	fmt.Println("比較演算子")
	fmt.Println(1 == 1)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(true == false)
	fmt.Println(true != false)
	fmt.Println("-------------------------------")

	// 算術演算子
	fmt.Println("算術演算子")
	fmt.Println(1 + 2)
	fmt.Println(2 - 1)
	fmt.Println(3 * 4)
	fmt.Println(60 / 12)
	fmt.Println(60 / 7)
	fmt.Println(60 % 7)

	fmt.Println("a" + "b")

	n := 0
	n += 4
	n++
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)
}
