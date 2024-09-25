package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"razi", "aziz", "syahputro"})
	_ = writer.Write([]string{"aziz", "razi", "syahputro"})
	_ = writer.Write([]string{"syahputro", "aziz", "razi"})

	writer.Flush()
}
