package iterutils

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
