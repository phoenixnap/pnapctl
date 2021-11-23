package tables

import "time"

func DerefStringList(list *[]string) []string {
	if list == nil {
		return []string{}
	}

	return *list
}

func DerefString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func DerefTimeAsString(time *time.Time) string {
	if time == nil {
		return ""
	}

	return time.String()
}
