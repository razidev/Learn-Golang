package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	return value1 == value2
}

func TestComparable(t *testing.T) {
	assert.False(t, IsSame("Razi", "Aziz"))
	assert.True(t, IsSame[int](123, 123))
}
