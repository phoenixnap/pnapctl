package iterutils

import (
	"reflect"
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

// Defines a function that performs a check on the item passed.
//
//	var isEven Predicate[int]
//
//	isEven := func(a int) bool {
//		return a % 2 == 0
//	}
type Predicate[T any] func(T) bool

// Runs the function `action` on each element in the `slice`.
//
//	nums := []int{1, 2, 3}
//	Each(nums, func(n int) { fmt.Println(n) })
func Each[T any](slice []T, action func(T)) {
	for _, item := range slice {
		action(item)
	}
}

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

// Same as Map, except it starts by referencing each element in a list.
// Useful for passing a list of non-pointers into a function that only
// accepts pointers.
//
//	nums := []int{1, 2, 3}
//	double := MapRef(nums, func(n *int) int {
//		return *n * 2
//	})
func MapRef[T, O any](slice []T, mapper Mapper[*T, O]) (sliceOut []O) {
	sliceOut = make([]O, len(slice))

	for i, v := range slice {
		sliceOut[i] = mapper(&v)
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
// For a list of interface{} - use DerefInterface.
//
//	num := 5
//	ptrs := []int{}{&num, &num} // [&5, &5]
//	vals := Deref(ptrs) // [5, 5]
func Deref[T any](slice []*T) []T {
	deref := func(item *T) T {
		return *item
	}

	return Map(slice, deref)
}

// Dereferences each element in a list if they are a pointer.
// Otherwise it keeps them as is.
//
// Only works for a list of interface{} - for concrete types,
// use Deref.
//
//	num := 5
//	ptrs := []interface{}{&num, &num} // [&5, &5]
//	vals := DerefInterface(ptrs) // [5, 5]
func DerefInterface(slice []interface{}) []interface{} {
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

// Same as map, except it implicitly casts the output to an interface{}.
//
//	nums := []int{1, 2, 3, 4} // [&5, &5]
//	doubled := MapInterface(nums, func(n int) int {
//		return n * 2
//	})
//	doubled[0] // 2 as interface{}
func MapInterface[T, U any](slice []T, mapper Mapper[T, U]) []interface{} {
	inter := func(item T) interface{} {
		var elem interface{}
		elem = mapper(item)
		return elem
	}

	return Map(slice, inter)
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
	return item == nil || (val.Kind() == reflect.Pointer && val.IsNil())
}
