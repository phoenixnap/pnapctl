package client

import (
	"time"
)

func IsZero[T comparable](item T) bool {
	var t T
	return t == item
}

func ParseDate(item string) *time.Time {
	ft, err := time.Parse(time.RFC3339, item)
	if err != nil {
		return nil
	}
	return &ft
}
