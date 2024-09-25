package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Razi Aziz Syahputro"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decoded))
	}
}
