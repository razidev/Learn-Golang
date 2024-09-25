package main

import "fmt"

func sayHellow(name string) string {
	if name == "" {
		return "Hello Bang jago"
	} else {
		return "Hello bang " + name
	}
}
func main() {
	result := sayHellow("razi")
	fmt.Println(result)
	fmt.Println(sayHellow(""))
}
