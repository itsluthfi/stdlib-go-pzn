package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// kalo mau pake sort harus implement interfacenya sort dulu, ada: Len, Less, Swap
func (userSlice UserSlice) Len() int { // gaperlu pointer karena slice itu defaultnya udah pointer
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	// conventional way
	// temp := userSlice[i]
	// userSlice[i] = userSlice[j]
	// userSlice[j] = temp
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := UserSlice{ // langsung pake type interfacenya
		{"Luthfi", 22},
		{"Izzuddin", 23},
		{"Hanif", 21},
		{"Lupi", 25},
		{"Hanip", 20},
	}

	sort.Sort(users)
	fmt.Println(users)

	// atau
	users1 := []User{ // ini pake struct dulu
		{"Luthfi", 22},
		{"Izzuddin", 23},
		{"Hanif", 21},
		{"Lupi", 25},
		{"Hanip", 20},
	}

	sort.Sort(UserSlice(users1))
	fmt.Println(users1)
}
