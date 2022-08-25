package iterutils

func Map[T, O any](slice []T, mapper func(T) O) (sliceOut []O) {
	sliceOut = make([]O, len(slice))

	for i, v := range slice {
		sliceOut[i] = mapper(v)
	}

	return
}

func Contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if item == s {
			return true
		}
	}
	return false
}
