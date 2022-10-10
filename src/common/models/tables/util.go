package tables

import (
	"fmt"
	"time"
)

func Deref[T any](item *T) string {
	if item == nil {
		return ""
	}
	return fmt.Sprintf("%v", *item)
}

type actuallyString interface {
	~string
}

func DerefString[T actuallyString](item *T) string {
	if item == nil {
		return ""
	}
	return string(*item)
}

func DerefStringable[T fmt.Stringer](item *T) string {
	if item == nil {
		return ""
	}

	return (*item).String()
}

func DerefTimeAsString(time *time.Time) string {
	if time == nil {
		return ""
	}

	return time.String()
}
