package p1_assert

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type user struct {
	id string
}

func TestAssertCase2(t *testing.T) {
	assertion := assert.New(t)

	name := "Bob"
	age := 10

	assertion.Equalf("Bob", name, "they should be equal")
	assertion.NotEqualf(20, age, "they should not be equal")

	assertion.Nil(nil, "they should be nil")
	assertion.NotNil("nil", "they should not be nil")

	var i = user{}
	assertion.Empty(i, "they should be empty") //可以是0, "", false, 0.0, 不能是" "和' '
	assertion.NotEmpty("/", "they should not be empty")

	err := errors.New("new error")
	var err2 error
	assertion.Error(err, "they should be a error")
	assertion.NoError(err2, "they should not be a error")

	var flag bool
	assertion.True(true, "they should be true")
	assertion.False(flag, "they should be false")
}
