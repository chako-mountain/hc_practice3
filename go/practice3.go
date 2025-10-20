package main

import (
	"fmt"
)

type MyIntSlice []int

func (sl MyIntSlice) Unique() MyIntSlice {
	check := map[int]bool{}
	var result MyIntSlice

	for _, v := range sl {
		if !check[v] {
			check[v] = true

			result = append(result, v)
		}
	}

	return result
}

func main() {

	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(m.Unique()) // [1, 2, 3, 4, 5]

}
