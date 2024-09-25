package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type userSlice []User

func (value userSlice) Len() int {
	return len(value)
}

func (value userSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value userSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	i := []int{5, 2, 6, 3, 1, 4} // unsorted number
	sort.Ints(i)
	fmt.Println(i)

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"} // unsorted string
	sort.Strings(s)
	fmt.Println(s)

	users := []User{
		{"Leonardo da Vinci", 25},
		{"Michelangelo", 30},
		{"Rembrandt", 40},
		{"Vermeer", 23},
	}

	fmt.Println(users)
	sort.Sort(userSlice(users))
	fmt.Println(users)
}
