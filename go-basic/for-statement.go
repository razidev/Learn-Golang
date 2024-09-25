package main

import "fmt"

func main() {
	slice := []string{"razi", "aziz", "syahputro"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println("index", i, "=", v)
	}

	person := make(map[string]string)
	person["name"] = "razi"
	person["job"] = "engineer"
	person["salary"] = "$10000"

	for key, v := range person {
		fmt.Println("key", key, "=", v)
	}
}
