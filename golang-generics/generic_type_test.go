package golanggenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Razi", "Aziz", "Syahputro"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	names := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(names)
}
