package main

import "fmt"

// 定数
// 定数は基本的にグローバルに書かれることが多い
// 最初の文字を大文字にするとほかのパッケージから呼び出せる。(public)
// 最初の文字が小文字ならprivate
const Pi = 3.14

const (
	URL      = "http://google.com"
	SiteName = "google"
)

const (
	A = 1
	B
	C
	D
	E = "A"
	F
)

const (
	c0 = iota
	c1
	c2
)

var Big int = 9223372036854775807

const Big2 = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)
	fmt.Println(URL)
	fmt.Println(SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(Big + 1)
	fmt.Println(Big2 - 1)
	fmt.Println(c0, c1, c2)
}
