package main

import (
	"encoding/json"
	"os"
)

type PersonalDataEncode struct {
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
	writer, _ := os.Create("data/CustomerOut.json")
	encoder := json.NewEncoder(writer)

	p1 := PersonalDataEncode{
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

	encoder.Encode(p1)

}
