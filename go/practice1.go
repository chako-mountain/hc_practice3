package main

import (
	"fmt"
	"strconv"
)

func main() {
	var unmber = []interface{}{1, "2", 10, "11"}
	var ints []int

	for _, content := range unmber {
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
}
