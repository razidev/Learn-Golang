package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`az([a-z])z`)

	fmt.Println(regex.MatchString("aziz"))
	fmt.Println(regex.MatchString("hiyah"))

	search := regex.FindAllString("razi aziz syahputro", 1)
	fmt.Println(search)
}

//github.com/google/re2/wiki/Syntax
