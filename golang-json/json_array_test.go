package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Nurul",
		MiddleName: "Iza",
		LastName:   "Hiyah",
		Hobbies:    []string{"Masak", "baca", "traveling"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Nurul","MiddleName":"Iza","LastName":"Hiyah","Age":0,"Married":false,"Hobbies":["Masak","baca","traveling"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Razi",
		Addresses: []Address{
			{
				Street:     "Jl. STM Walang Jaya",
				Country:    "Indonesia",
				PostalCode: "14260",
			},
			{
				Street:     "GG. Teladan V",
				Country:    "Indonesia",
				PostalCode: "14260",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Razi","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl. STM Walang Jaya","Country":"Indonesia","PostalCode":"14260"},{"Street":"GG. Teladan V","Country":"Indonesia","PostalCode":"14260"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl. STM Walang Jaya","Country":"Indonesia","PostalCode":"14260"},{"Street":"GG. Teladan V","Country":"Indonesia","PostalCode":"14260"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	json.Unmarshal(jsonBytes, addresses)

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jl. STM Walang Jaya",
			Country:    "Indonesia",
			PostalCode: "14260",
		},
		{
			Street:     "GG. Teladan V",
			Country:    "Indonesia",
			PostalCode: "14260",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println("=>>>", string(bytes))
}
