package utils

func Is2xxSuccessful(statusCode int) bool {
	if statusCode >= 200 && statusCode < 300 {
		return true
	} else {
		return false
	}
}
