package generators

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func Generate[T any]() (t T) {
	ignoreAdditionalProperties := options.WithFieldsToIgnore("AdditionalProperties")
	faker.FakeData(&t, ignoreAdditionalProperties)
	return
}

func GenerateWithOpts[T any](opt ...options.OptionFunc) (t T) {
	faker.FakeData(&t, opt...)
	return
}

func Generator[T any](updates ...func(*T)) func() T {
	return func() T {
		initial := Generate[T]()
		iterutils.Each(updates, func(u func(*T)) {
			u(&initial)
		})
		return initial
	}
}

func OneOfGenerator[T any, Inner any](wrapper func(T) Inner, updates ...func(*T)) func() Inner {
	return func() Inner {
		raw := Generator(updates...)
		return wrapper(raw())
	}
}
