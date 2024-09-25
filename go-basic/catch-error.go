package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("data/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	nf, err := os.Open("data/no-file.txt")
	if err != nil {
		// fmt.Println("err occured", err) //it shown in terminal
		// log.Println("err occured", err) //it will log to data/log.txt
		// log.Fatalln(err) //it will stop the program, and shown the number's error
		panic(err) //it will stop the program, should recover to execute after the panic
	}
	defer nf.Close()
	fmt.Print("End of File")
}
