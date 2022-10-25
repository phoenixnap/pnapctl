package testutil

// Takes a pointer of whatever is being passed.
func AsPointer[T any](item T) *T {
	return &item
}

// Defines a function that can generate a specific type.
//
//	var generateInt Generator[int]
//
//	generateInt = func() int {
//		return rand.Int32()
//	}
type Generator[T any] func() T

// Generates n items using the generator passed.
//
//	generateInt := func() int {
//		return rand.Int32()
//	}
//
//	nums := GenN(10, generateInt) // 10 integers
func GenN[T any](n int, gen Generator[T]) (out []T) {
	out = make([]T, n)

	for i := range out {
		out[i] = gen()
	}

	return
}

// Generates n items using the generator passed, and dereferences them.
//
//	generateInt := func() *int {
//		num := rand.Int32()
//		return &num
//	}
//
//	nums := GenN(10, generateInt) // 10 integers (not pointers)
func GenNDeref[T any](n int, gen Generator[*T]) (out []T) {
	out = make([]T, n)

	for i := range out {
		out[i] = *gen()
	}

	return
}
