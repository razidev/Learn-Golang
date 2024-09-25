package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PersonalData struct {
	Name      string         `json:"name"`
	Hobbies   []string       `json:"hobbies"`
	Books     map[int]string `json:"books"` //for unpredictable the key's objects
	Addresses []struct {
		Name       string `json:"name"`
		Street     string `json:"street"`
		Country    string `json:"country"`
		PostalCode string `json:"postalcode"`
	} `json:"address"`
}

func main() {
	reader, _ := os.Open("data/customer.json")
	decoder := json.NewDecoder(reader)

	humanData := &[]PersonalData{}
	decoder.Decode(humanData)

	fmt.Println(humanData)
}
