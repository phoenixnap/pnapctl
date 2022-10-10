package generators

import (
	"github.com/go-faker/faker/v4"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func Generate[T any]() (t T) {
	faker.FakeData(&t)
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