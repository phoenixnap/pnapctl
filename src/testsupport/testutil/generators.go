package testutil

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandSeqPointer(n int) *string {
	random := RandSeq(n)
	return &random
}

func RandListStringPointer(n int) []string {
	b := make([]string, n)
	for i := range b {
		b[i] = RandSeq(10)
	}
	return b
}

func RanNumberPointer() *int32 {
	i := rand.Int31()
	return &i
}

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
