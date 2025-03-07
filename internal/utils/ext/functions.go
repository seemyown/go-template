package ext

func Contains[T comparable](slice []T, item T) bool {
	if slice == nil {
		return false
	}
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func Ternary[T any](condition bool, a T, b T) T {
	if condition {
		return a
	} else {
		return b
	}
}
