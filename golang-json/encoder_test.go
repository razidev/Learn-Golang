package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Razi",
		LastName:  "Syahputro",
		Married:   false,
	}

	encoder.Encode(customer)
}
