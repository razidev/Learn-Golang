package main

import (
	"fmt"
	helper "learn-go-basic/helper"
)

func main() {
	helper.SayHello("Razi")
	// helper.sayGoodBye("Razi") //not public
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //not public
}
