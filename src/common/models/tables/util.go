package tables

import (
	"fmt"
	"time"
)

func DerefStringList(list *[]string) []string {
	if list == nil {
		return []string{}
	}

	return *list
}

// func DerefString(str *string) string {
// 	if str == nil {
// 		return ""
// 	}

// 	return *str
// }

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
