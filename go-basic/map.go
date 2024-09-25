package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "razi",
		"age":  "23",
	}
	person["job"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["name"])

	m := map[string]int{
		"Razi":  25,
		"Hiyah": 23,
	}
	fmt.Println(m)
	v, ok := m["Aziz"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["Razi"]; ok {
		fmt.Println("the output is", v)
	}
	m["Aziz"] = 24
	for k, v := range m {
		fmt.Println(k, v)
	}

	var book map[string]string = make(map[string]string)
	book["title"] = "learn go"
	book["author"] = "eko"
	book["delete"] = "please"
	fmt.Println(book)
	delete(book, "delete")
	fmt.Println(book)
}
