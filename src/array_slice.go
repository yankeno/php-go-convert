package main

import "fmt"

func main() {
	/**************************************
	 * 配列
	 *************************************/
	var x [3]int   // 長さ3のint型配列
	fmt.Println(x) // [0 0 0]

	var y = [3]int{10, 20, 30} // 配列の初期化
	fmt.Println(y[1])

	fruits := []string{"apple", "banana", "cherry"}
	for _, fruit := range fruits {
		fmt.Println(fruit) // apple, banana, cherry
	}

	var a = [...]string{"a", "b"} // 配列の長さを省略
	var b = [2]string{"a", "b"}
	fmt.Println(a == b)

	/**************************************
	 * スライス
	 *************************************/
	var c = []int{10, 20, 30}
	fmt.Println(c) // [10 20 30]

	var d []int           // nilで初期化される
	fmt.Println(d == nil) // true

	var e []int
	e = append(e, 10, 20, 30) // 要素の追加
	// fmt.Println(c == e)            // invalid operation: c == e (slice can only be compared to nil) ※スライスと比較可能なのはnilのみ

	f := make([]int, 5, 10)     // 長さ5、キャパシティ10のスライス
	fmt.Println(len(f), cap(f)) // 5 10

	g := []int{1, 2, 3, 4}
	// slice[n:m] でサブスライスを作成(n:開始オフセット, m:終了オフセット)
	h := g[:2]
	i := g[1:]
	j := g[1:3]
	k := g[:]
	fmt.Println("g:", g) // [1 2 3 4]
	fmt.Println("h:", h) // [1 2]
	fmt.Println("i:", i) // [2 3 4]
	fmt.Println("j:", j) // [2 3]
	fmt.Println("k:", k) // [1 2 3 4]

	// メモリを共有しているため、スライスの要素を変更するとその要素を共有している他のスライスにも影響が出る
	h[1] = 100
	fmt.Println("g:", g) // [1 100 3 4]

	h = append(h, 200)
	fmt.Println("h:", h) // [1 100 200]
	fmt.Println("g:", g) // [1 100 200 4]

	l := []int{1, 2, 3, 4}
	m := l[:2:2]
	n := l[2:4:4]
	fmt.Println("l:", l) // [1 2 3 4]
	fmt.Println("m:", m) // [1 2]
	fmt.Println("n:", n) // [3 4]

	// フルスライス式を使用してスライスを作成すると、元のスライスとはメモリを共有しない
	m = append(m, 100)
	n = append(n, 200)
	fmt.Println("l:", l) // [1 2 3 4]
	fmt.Println("m:", m) // [1 2 100]
	fmt.Println("n:", n) // [3 4 200]
}
