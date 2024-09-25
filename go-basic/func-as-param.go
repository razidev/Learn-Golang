package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	return strings.ToUpper(name)
}

func main() {
	sayHelloFilter("World", spamFilter)
	sayHelloFilter("Go", spamFilter)
}
