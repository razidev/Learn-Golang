package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 Address = address1
	var address3 *Address = &address2

	address3.City = "Bandung"
	// address3 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} //create new memory for address3
	*address3 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} //create new same memory for address3 & address2

	fmt.Println("address1", address1)
	fmt.Println("address2", address2)
	fmt.Println("address3", address3)

	//create a empty pointer
	var address4 *Address = new(Address)
	address4.Province = "DIY Jogjajakarta"
	fmt.Println("address4", address4)
}
