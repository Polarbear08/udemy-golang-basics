package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 型変換
	var ii int = 1
	f641 := float64(ii)
	fmt.Println(f641)
	fmt.Printf("ii = %T\n", ii)
	fmt.Printf("f641 = %T\n", f641)

	ii2 := int(f641)
	fmt.Printf("i2 = %T", ii2)

	var ss string = "100"
	fmt.Printf("s = %T\n", ss)

	ii3, _ := strconv.Atoi(ss)
	fmt.Println(ii3)
	fmt.Printf("i = %T\n", ii3)

	var ii4 int = 200
	ss2 := strconv.Itoa(ii4)
	fmt.Println(ss2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(h)
	fmt.Println(h2)

	// interface と　nil
	var x interface{} // interfaceはあらゆる型と互換性がある
	fmt.Println(x)    // nilは何もない

	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3.14
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// fmt.Println(x + 3)	interfaceは演算の対象にはできない

	// 配列型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)

	arr3 := [3]int{1}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"} // 要素数を自動でカウント
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr1[0])
	fmt.Println(arr4[1])
	fmt.Println(arr4[2-1])

	fmt.Println(arr2[2])
	arr2[2] = "c"
	fmt.Println(arr2[2])

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	fmt.Println(len(arr1))

	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	// 文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`
	test
	test
	test
	`)

	fmt.Println("\"\"")
	fmt.Println(`""`)
	fmt.Println(string(s[0]))
	fmt.Println(s[0])

	// 論理値型
	var tr, fa bool = true, false
	fmt.Println(tr, fa)
	fmt.Printf("%T\n", tr)

	// uint, complex型
	var u8 uint = 255
	var c64 complex64 = 5 + 12i
	fmt.Println(u8, c64)
	fmt.Printf("%T\n", u8)
	fmt.Printf("%T\n", c64)

	// 浮動小数点型
	var f64 float64 = 2.4
	fmt.Println(f64)

	f := 3.2 // 暗黙的な定義なら、環境によらずfloat64
	fmt.Println(f)

	fmt.Printf("%T, %T\n", f, f64)

	var f32 float32 = 1.2
	fmt.Println(f32)

	fmt.Printf("%T\n", float64(f32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	// 整数型
	var i int = 100 // 環境依存。64bitPCならばint64
	var i2 int64 = 200

	fmt.Println(i, i2)
	fmt.Println(i + 50)
	fmt.Println(i2 + 50)

	fmt.Printf("%T\n", i2) // %Tは型を表示するフォーマット指定子

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

}
