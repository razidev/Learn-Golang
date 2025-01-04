package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, int64(100), FindMin[int64](100, 200))
	assert.Equal(t, 100.5, FindMin(200, 100.5))
	assert.Equal(t, float64(100.5), FindMin(200, 100.5))
}

func GetFirst[T []E, E any](slice T) E {
	return slice[0]
}

func TestGetFirst(t *testing.T) {
	assert.Equal(t, 1, GetFirst([]int{1, 2, 3}))
	assert.Equal(t, "Razi", GetFirst([]string{"Razi", "Aziz"}))
}
