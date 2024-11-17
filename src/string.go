package main

import "fmt"

func main() {
	var a string
	fmt.Println(a) // ゼロ値は""

	// 文字列の連結
	b := "Hello"
	c := " World"
	fmt.Println(b + c) // Hello World

	var d string = "Hello World!"
	var e byte = d[4]
	fmt.Printf("%s\n", string(e)) // o

	// スライス式で部分文字列を取得
	f := d[6:11]
	fmt.Println(f) // World

	// 絵文字や日本語は2バイト以上の文字列となるため、スライス式で部分文字列を取得すると文字化けする場合がある
	var g = "Hello 😎"
	fmt.Println(g[6:7])   // �
	fmt.Println(g[6:8])   // �
	fmt.Println(g[6:9])   // �
	fmt.Println(g[6:10])  // 😎
	fmt.Println(len(g))   // 10
	fmt.Println(len("あ")) // 3
}
