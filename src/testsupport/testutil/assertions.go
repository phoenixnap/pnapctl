package testutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertIsType[T any](test_framework *testing.T, item interface{}) {
	var t T
	assert.Equal(
		test_framework,
		fmt.Sprintf("%T", item),
		fmt.Sprintf("%T", t),
	)
}
