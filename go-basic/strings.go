package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Razi Syahputro", "putro"))
	fmt.Println(strings.Split("Razi Syahputro", " "))
	fmt.Println(strings.ToLower("Razi Syahputro"))
	fmt.Println(strings.ToUpper("Razi Syahputro"))
	fmt.Println(strings.Trim("      Razi Syahputro      ", " "))
	fmt.Println(strings.ReplaceAll("Razi Syahputro Razi Aziz", "Razi", "Budi"))
}
