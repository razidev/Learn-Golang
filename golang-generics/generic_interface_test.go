package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](gs GetterSetter[T], value T) T {
	gs.SetValue(value)

	return gs.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "Razi")

	assert.Equal(t, "Razi", result)
	assert.Equal(t, "Razi", myData.Value)
}
