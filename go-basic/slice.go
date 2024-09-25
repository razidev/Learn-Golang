package main

import "fmt"

func main() {

	//ini array
	var month = [...]string{
		"Januari",   // 0
		"Februari",  // 1
		"Maret",     // 2
		"April",     // 3
		"Mei",       // 4
		"Juni",      // 5
		"Juli",      // 6
		"Agustus",   // 7
		"September", // 8
		"Oktober",   // 9
		"November",  // 10
		"Desember",  // 11
	}

	var slice1 = month[4:7]
	fmt.Println("slice1", slice1)
	fmt.Println("slice1-len", len(slice1))
	fmt.Println("slice1-cap", cap(slice1))

	// month[5] = "ubah"
	// fmt.Println(slice1)

	// slice1[0] = "mai-ubah"
	// fmt.Println(month)

	var slice2 = month[10:]
	fmt.Println("slice2", slice2)

	var slice3 = append(slice2, "razi", "aziz")
	fmt.Println("slice3", slice3)
	slice3[1] = "ubah-december"
	fmt.Println("ubah slice3", slice3)
	fmt.Println("slice2", slice2)
	fmt.Println("month", month)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "razi"
	newSlice[1] = "hiyah"
	// newSlice[2] = "aziz" : akan error, harus gunakan append
	newSlice = append(newSlice, "aziz")
	fmt.Println("newSlice:", newSlice)
	fmt.Println("newSlice-len:", len(newSlice))
	fmt.Println("newSlice-cap:", cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copySlice:", copySlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	iniSlice2 := []int{6, 7, 8, 9, 10}
	fmt.Println("iniarray:", iniArray)
	fmt.Println("inislice:", iniSlice)

	iniSlice2 = append(iniSlice2, iniSlice...)
	fmt.Println("iniSlice2", iniSlice2)
	for i, v := range iniSlice2 {
		fmt.Printf("index ke %v = %v\n", i, v)
	}
	fmt.Printf("%T\n", iniSlice2)

	deleteSomeininiSlice2 := append(iniSlice2[:2], iniSlice2[4:]...)
	fmt.Println("deleteSomeininiSlice2", deleteSomeininiSlice2)
}

// hati2 dalam menggunakan slice jika capasitas array masih dalam kapasitasnnya maka variabel
// data array inisialisasi awalnya akan ikut keubah walau disimpan dalam variabel baru
