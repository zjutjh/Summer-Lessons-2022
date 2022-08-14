package p1_assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertCase3(t *testing.T) {
	assertion := assert.New(t)

	assertion.Contains("hello world", "world")
	assertion.NotContains("hello world", "helllo")

	s := make([]int, 4)
	assertion.Len(s, 4)

	assertion.FileExists("/Users/patrick_star/go/Summer-Lessons-2022/testify/p1_assert/case3_test.go")
	assertion.DirExists("/Users/patrick_star/go/Summer-Lessons-2022/testify/p1_assert")
}
