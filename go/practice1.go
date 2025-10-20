package main

import (
	"fmt"
	"strconv"
)

func main() {
	// スライスを直接リテラルで作成
	// fruits := []string{"apple", "banana", "cherry"}

	var unmber = []interface{}{1, "2", 10, "11"}
	var ints []int

	// unmber_int = int(unmber)
	// fmt.Println(unmber)

	// // 要素を追加（appendを使う）
	// fruits = append(fruits, "orange")

	// // スライスの内容を表示
	// fmt.Println(fruits)

	// // 長さと容量を確認
	// fmt.Println("len:", len(fruits))
	// fmt.Println("cap:", cap(fruits))

	for _, content := range unmber {
		// fmt.Println(content)

		// int_number := strconv.Atoi(content)
		// fmt.Print(int_number)

		switch i := content.(type) {

		case int:
			ints = append(ints, i)

		case string:
			num, err := strconv.Atoi(i)

			if err != nil {
				fmt.Println("変換に失敗しました", i)
			} else {
				ints = append(ints, num)
			}

		}

	}

	for _, n := range ints {
		fmt.Printf("%02d\n", n)
	}

	// var s string = "100"
	// fmt.Print(s)

	// i, _ := strconv.Atoi(s)
	// fmt.Println(i)
	// fmt.Print(i)
}
