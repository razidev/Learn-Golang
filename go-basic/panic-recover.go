package main

import "fmt"

func endApp() {
	message := recover() //catch the panic
	if message != nil {
		fmt.Println("error message", message)
	}
	fmt.Println("Aplikasi done")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("App Error") //stop after the code if error
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Razi") //if not using recover, Razi doesn't show
}
