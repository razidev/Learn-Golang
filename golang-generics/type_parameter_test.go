package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Generic with constraints (constraints is must)
func Hello[T any](param T) T {
	return param
}

func TestHello(t *testing.T) {
	var char string = Hello[string]("razi")
	assert.Equal(t, "razi", char)

	var number int = Hello(123)
	assert.Equal(t, 123, number)
}

func MultipleTypeParamaeter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleTypeParamaeter(t *testing.T) {
	MultipleTypeParamaeter("Hello", 123)
	MultipleTypeParamaeter[int, string](123, "Razi")
}
