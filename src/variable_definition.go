package main

import "fmt"

func main() {
	var name string = "Taro"                 // 型を指定して宣言
	age := 25                                // 「:=」を使用した変数宣言 ※関数内のみで使用可能
	var isAdult = true                       // 型を省略して宣言
	var height, nationality = 170.5, "Japan" // 多重代入 ※関数の戻り値を受け取る時だけ使用するのが一般的

	fmt.Println(name, age, isAdult, height, nationality)
}
