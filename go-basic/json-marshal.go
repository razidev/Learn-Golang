package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string         `json:"name"`
	Hobbies   []string       `json:"hobbies"`
	Books     map[int]string `json:"books"` //for unpredictable the key's objects
	Addresses []Address      `json:"address"`
}

type Address struct {
	Name       string `json:"name"`
	Street     string `json:"street"`
	Country    string `json:"country"`
	PostalCode string `json:"postalcode"`
}

func main() {
	p1 := Person{
		Name:    "Razi",
		Hobbies: []string{"Badminton", "Games"},
		Books: map[int]string{
			1: "Atomic Habits",
			2: "The Intelligent Investor",
			3: "Supernova",
		},
		Addresses: []Address{
			{
				Name:       "Home",
				Street:     "Jl. Nusantara",
				Country:    "Indonesia",
				PostalCode: "14444",
			},
			{
				Name:       "Office",
				Street:     "Jl. Megapolitan",
				Country:    "Indonesia",
				PostalCode: "14440",
			},
		},
	}
	p2 := Person{
		Name:    "Syahputro",
		Hobbies: []string{"Football", "Boxing"},
		Books: map[int]string{
			1: "The Lord of The Ring",
			2: "Harry Potter",
			3: "The Great Gatsby",
		},
	}

	bytes1, _ := json.Marshal(p1)
	bytes2, _ := json.Marshal(p2)
	fmt.Println(string(bytes1))

	slice := []string{string(bytes1), string(bytes2)}
	fmt.Println(slice)
}
