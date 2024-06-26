package slice

func IsInSlice(str string, strList []string) bool {
	for _, s := range strList {
		if s == str {
			return true
		}
	}
	return false
}

func IsContains[T comparable](element T, slice []T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
