package main

import "fmt"

func main() {
	var nilMap map[string]int  // 初期値nil
	fmt.Println(len(nilMap))   // 0
	fmt.Println(nilMap["abc"]) // 0(intのゼロ値)
	//nilMap["abc"] = 3        // panic: nilマップへの書き込みはできない

	wins := map[string]int{} // string型のキーとint型の値を持つ空のマップ
	fmt.Println(wins == nil) // false
	fmt.Println(wins["abc"]) // 0(intのゼロ値)
	wins["abc"] = 3          // 書き込み可能
	fmt.Println(wins["abc"]) // 3

	teams := map[string][]string{
		"giants": {"坂本", "岡本"},
		"tigers": {"糸井", "大山"},
		"carp":   {"鈴木", "菊池"},
	}
	fmt.Println(teams["giants"][0]) // 坂本

	ages := make(map[string]int, 10) // サイズを指定してmapを作成
	fmt.Println(len(ages))           // 0

	user := map[string]any{
		"name": "Taro",
		"age":  25,
	}
	fmt.Println(user["name"]) // Taro
	delete(user, "age")
	fmt.Println(user)

	// マップのキーの存在チェック
	m := map[string]int{
		"apple":  100,
		"banana": 200,
	}

	v, ok := m["apple"]
	fmt.Println(v, ok) // 100 true

	v, ok = m["banana"]
	fmt.Println(v, ok) // 200 true

	v, ok = m["nothing"]
	fmt.Println(v, ok) // 0 false
}
