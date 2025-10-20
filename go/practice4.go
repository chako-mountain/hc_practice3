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
