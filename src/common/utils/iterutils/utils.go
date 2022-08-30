package iterutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Mapper[T, O any] func(T) O

// Applies the function `mapper` to each element in the `slice`.
// Returns the resulting mapped array.
func Map[T, O any](slice []T, mapper Mapper[T, O]) (sliceOut []O) {
	sliceOut = make([]O, len(slice))

	for i, v := range slice {
		sliceOut[i] = mapper(v)
	}

	return
}

// Checks whether the `slice` contains the `item` passed.
func Contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if item == s {
			return true
		}
	}
	return false
}

type BiMapper[T, U, O any] func(T, U) O

// Same as Map, except it accepts 2 parameter in the mapping function.
func BiMap[T, U, O any](items []T, param U, mapper BiMapper[T, U, O]) []O {
	preparer := func(item T) O {
		return mapper(item, param)
	}

	return Map(items, preparer)
}

type AssertFn[T, U any] func(*testing.T, T, U)

// Asserts that elements on both 'expected' and 'actual' both pass the assertions in 'asserter'
// Also asserts that 'expected' and 'actual' are both of the same length.
func AssertOnListElements[T, U any](t *testing.T, expected []T, actual []U, asserter AssertFn[T, U]) {
	assert.Len(t, expected, len(actual))

	for i := range expected {
		asserter(t, expected[i], actual[i])
	}
}
