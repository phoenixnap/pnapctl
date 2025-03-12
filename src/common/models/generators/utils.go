package generators

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// Randomly generate data of type T.
func Generate[T any]() (t T) {
	// using WithIgnoreFields sets AdditionalProperties to nil which causes tests
	// to break, as upon unmarshalling it instead becomes an empty map.
	//
	// so instead, we use a custom provider to auto-set this property as empty.
	emptyAdditionalProperties := options.WithCustomFieldProvider(
		"AdditionalProperties",
		func() (interface{}, error) {
			return map[string]interface{}{}, nil
		})

	faker.FakeData(&t, emptyAdditionalProperties)
	return
}

// Randomly generate data of type T. Accepts options to customize generation.
func GenerateWithOpts[T any](opt ...options.OptionFunc) (t T) {
	faker.FakeData(&t, opt...)
	return
}

// Constructs and returns a generation function.
// Passed methods can be used to update the data post-generation to ensure validity.
func Generator[T any](updates ...func(*T)) func() T {
	return func() T {
		initial := Generate[T]()
		iterutils.Each(updates, func(u func(*T)) {
			u(&initial)
		})
		return initial
	}
}

// Similar to Generator(...) except it works for OneOf types, which usually have a wrapper.
// An example of a wrapper is RatedUsageGet200ResponseInner.
func OneOfGenerator[T any, Inner any](wrapper func(T) Inner, updates ...func(*T)) func() Inner {
	return func() Inner {
		raw := Generator(updates...)
		return wrapper(raw())
	}
}
