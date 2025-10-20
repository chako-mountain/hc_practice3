package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	userPtrs := make([]*User, len(users))
	for i := range users {
		userPtrs[i] = &users[i]
	}

	for _, user := range userPtrs {
		user.Age = 44
	}

	fmt.Printf("%v\n", users)
}

// 説明
// 全ての人間の年齢が14にならなかったのは、forで使用しているuserがusersの単なるコピーであるためである。
// コピーをいくら書き換えても元のusersスライスには影響しない。ので、全ての人間の年齢は元のままである。
// 上記コードでは、修正案として、usersスライスの各要素のポインタを格納するuserPtrsスライスを作成し、forループ内でそのポインタを使って元のusersスライスの要素を書き換えている。
