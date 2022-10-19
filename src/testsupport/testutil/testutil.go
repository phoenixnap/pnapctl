package testutil

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// A generic test error.
var TestError = errors.New("TEST ERROR")
var TestKeycloakError = ctlerrors.CLIError{Message: "Failed to resolve provided credentials", Cause: TestError}

// A fake Error response from the server.
var GenericBMCError = ctlerrors.BMCError{
	Message:          "Something went wrong!",
	ValidationErrors: []string{},
}

// Utility struct used for the assertion pattern.
//
// Wraps a list of items.
type assertList[T any] struct {
	items []T
}

// Utility struct used for the assertion pattern.
//
// Wraps two lists of equally sized items.
type assertPairs[T, U any] struct {
	expected []T
	actual   []U
}

// Sets up the data to run an assertion on each of its items.
//
//	list := []int{2,4,6,8,10}
//	assertIsEven := func(t *testing.T, num int) {
//		assert.IsTrue(num % 2 == 0)
//	}
//
//	ForEach(list).Do(t, assertIsEven)
func ForEach[T any](slice []T) assertList[T] {
	return assertList[T]{
		items: slice,
	}
}

// Sets up the data to run an assertion using both lists' items.
//
//	list1 := []int{2,4,6,8,10}
//	list2 := []int{2,4,6,8,10}
//
//	assertEqual := func(t *testing.T, num1 int, num2 int) {
//		assert.Equal(t, num1, num2)
//	}
//
//	ForEachPair(list1, list2).Do(t, assertEqual)
func ForEachPair[T, U any](exp []T, act []U) assertPairs[T, U] {
	return assertPairs[T, U]{
		expected: exp,
		actual:   act,
	}
}

// Runs the assertion function passed for each element in the list.
//
//	list := []int{2,4,6,8,10}
//	assertIsEven := func(t *testing.T, num int) {
//		assert.IsTrue(num % 2 == 0)
//	}
//
//	ForEach(list).Do(t, assertIsEven)
func (a assertList[T]) Do(t *testing.T, asserter func(*testing.T, T)) {
	for i := range a.items {
		asserter(t, a.items[i])
	}
}

// Runs the assertion function passed for each pair of elements in both lists.
//
// Sets up the data to run an assertion using both lists' items.
//
//	list1 := []int{2,4,6,8,10}
//	list2 := []int{2,4,6,8,10}
//
//	assertEqual := func(t *testing.T, num1 int, num2 int) {
//		assert.Equal(t, num1, num2)
//	}
//
//	ForEachPair(list1, list2).Do(t, assertEqual)
func (a assertPairs[T, U]) Do(t *testing.T, asserter func(*testing.T, T, U)) {
	for i := range a.expected {
		asserter(t, a.expected[i], a.actual[i])
	}
}

// Generic asserters

// Used as shorthand to assert two things are equal.
func AssertEqual[T any](t *testing.T, a, b T) {
	assert.Equal(t, a, b)
}

type Stringlike interface {
	~string
}

func AsStrings[S Stringlike](enums []S) []string {
	return iterutils.Map(enums, func(s S) string { return string(s) })
}

func AssertIsType[T any](test_framework *testing.T, item interface{}) {
	var t T
	assert.Equal(
		test_framework,
		fmt.Sprintf("%T", item),
		fmt.Sprintf("%T", t),
	)
}
