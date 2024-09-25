package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCounteryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func changeProvince(address Address) {
	address.Province = "DKI Jakarta"
	fmt.Println(address)
}

func main() {
	var alamat = Address{
		City:     "Ambarawa",
		Province: "Jawa Tengah",
		Country:  "",
	}

	var alamat2 = Address{
		City:     "Jakarta",
		Province: "",
		Country:  "Indonesia",
	}

	ChangeCounteryToIndonesia(&alamat)
	changeProvince(alamat2) // it won't change the Province
	fmt.Println(alamat)
	fmt.Print(alamat2)
}
