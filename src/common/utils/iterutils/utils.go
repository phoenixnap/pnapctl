package iterutils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* TYPES */

// Defines a function that maps from type In to type Out.
//
//	var s Mapper[[]byte, string]
//
//	s := func(bytes []byte) string {
//		return string(bytes)
//	}
type Mapper[In, Out any] func(In) Out

// Defines a function that can be 'Curried' into a Mapper.
//
//	var mul Currier[int, int, int]
//
//	mul := func(int a, int b) int {
//		return a*b
//	}
//
//	var doubler Mapper[int, int]
//	doubler := Curry(mul, 2)
type Currier[In, Const, Out any] func(In, Const) Out

// Defines a function that performs an assertion on T and U.
//
//	var mustBeEqual AssertFn[int, int]
//
//	mustBeEqual := func(t *testing.T, x int, y int) {
//		assert.Equal(t, x, y)
//	}
type AssertFn[T, U any] func(*testing.T, T, U)

// Defines a function that performs a check on the item passed.
//
//	var isEven Predicate[int]
//
//	isEven := func(a int) bool {
//		return a % 2 == 0
//	}
type Predicate[T any] func(T) bool

// Applies the function `mapper` to each element in the `slice`.
// Returns the resulting mapped array.
//
//	nums := []int{1, 2, 3}
//	double := Map(nums, func(n int) int {
//		return n * 2
//	})
func Map[T, O any](slice []T, mapper Mapper[T, O]) (sliceOut []O) {
	sliceOut = make([]O, len(slice))

	for i, v := range slice {
		sliceOut[i] = mapper(v)
	}

	return
}

// Curries the function passed, setting it up with the constant
//
//	func Multiply(item int, multiplier int) int {
//		return item * multiplier
//	}
//
//	doubler := Curry(Multiply, 2)
//	nums := []int{1, 2, 3}
//
//	doubled := Map(nums, doubler) // [2, 4, 6]
func Curry[In, Const, Out any](mapper Currier[In, Const, Out], constant Const) Mapper[In, Out] {
	return func(item In) Out {
		return mapper(item, constant)
	}
}

// Returns a new list that only contains elements which satisfy the predicate.
//
//	nums := []int{1, 2, 3, 4, 5}
//	even := Filter(nums, IsEven) // [2, 4]
func Filter[T any](slice []T, predicate Predicate[T]) (out []T) {
	for _, v := range slice {
		if predicate(v) {
			out = append(out, v)
		}
	}
	return
}

// Checks whether the `slice` contains the `item` passed.
//
//	nums := []int{1, 2}
//	hasTwo := Contains(nums, 2) // true
func Contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// Asserts that elements on both 'expected' and 'actual' both pass the assertions in 'asserter'
// Also asserts that 'expected' and 'actual' are both of the same length.
//
//	nums1 := []int{1, 2, 3}
//	nums2 := []int{1, 2, 3}
//
//	isEqual := func(t *testing.T, a int, b int) {
//		assert.IsEqual(t, a, b)
//	}
//
//	AssertOnListElements(t, nums1, nums2, isEqual) // passes
func AssertOnListElements[T, U any](t *testing.T, expected []T, actual []U, asserter AssertFn[T, U]) {
	assert.Len(t, expected, len(actual))

	for i := range expected {
		asserter(t, expected[i], actual[i])
	}
}

// Returns the first element in a list that satisfies the predicate passed.
// Returns nil if no element matched.
//
//	nums := []int{1, 2, 3, 4, 5}
//	num := FindElementThat(nums, IsEven) // &2
//	num := FindElementThat(nums, IsNil) // nil
func FindElementThat[T any](slice []T, predicate Predicate[T]) *T {
	if slice == nil {
		return nil
	}

	for _, v := range slice {
		if predicate(v) {
			return &v
		}
	}

	return nil
}

/* SPECIFIC MAPPERS */

// Dereferences each element in a list if they are a pointer.
// Otherwise it keeps them as is.
//
//	num := 5
//	ptrs := []*int{&num, &num} // [&5, &5]
//	vals := Deref(ptrs) // [5, 5]
func Deref(slice []interface{}) []interface{} {
	deref := func(item interface{}) interface{} {
		val := reflect.ValueOf(item)
		if val.Kind() == reflect.Ptr {
			return val.Elem().Interface()
		} else {
			return val.Interface()
		}
	}

	return Map(slice, deref)
}

/* PREDICATES */

// Flips boolean returned by predicate.
//
//	items := []string{"item", ""}
//
//	// This will only keep 'Zero' items.
//	zero := Filter(items, IsZero)
//
//	// This will do the opposite.
//	item := Filter(items, Not(IsZero))
func Not[T any](p Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !p(t)
	}
}

// Evaluates whether the item passed is 'nil' or not.
// Verifies whether the underlying value is 'nil' as well.
//
//	var x interface{}
//	IsNil(x) // true
//
//	var s *string
//	IsNil(s) // true
//
//	n := 3
//	IsNil(n) // false
func IsNil(item interface{}) bool {
	val := reflect.ValueOf(item)
	return item == nil || val.IsNil()
}
