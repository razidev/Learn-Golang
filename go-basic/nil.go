package main

import "fmt"

func NewMap(name string) map[string]string {
	if name != "" {
		return map[string]string{
			"name": name,
		}
	} else {
		return nil
	}
}

func main() {
	person := NewMap("razi")
	if person != nil {
		fmt.Println(person)
	} else {
		fmt.Println("data kosong")
	}
}

//nil only use in interface, function, map, slice, pointer, and channel. beside that defaul value
