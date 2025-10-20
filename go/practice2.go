package main

import (
	"fmt"
)

func findKeyByValue(m map[int]string, targetValue string) (int, error) {
	for key, value := range m {
		if value == targetValue {
			return key, nil
		}
	}
	return 0, fmt.Errorf("値が見つかりませんでした")
}

func main() {
	m := map[int]string{
		1: "01",
		2: "02",
		3: "03",
	}
	key, err := findKeyByValue(m, "03") // key→3, err→nil
	fmt.Print(key, err)
}
