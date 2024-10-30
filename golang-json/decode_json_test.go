package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonstring := `{"FirstName":"Razi","MiddleName":"Aziz","LastName":"Syahputro","Age":25,"Married":false}`
	jsonBytes := []byte(jsonstring)

	customer := &Customer{}

	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
}
