package p1_assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetValue(t *testing.T) {

	assert.Equal(t, "bob", getValue(1))
	assert.NotEqual(t, 20, getValue(2), "they should not be equal")
}
