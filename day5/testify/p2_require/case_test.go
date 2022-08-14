package p2_require

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequireCase(t *testing.T) {
	name := "Bob"
	age := 10

	require.Equal(t, "Bob", name, "they should be equal")
	require.NotEqual(t, 10, age, "they should not be equal")
}
