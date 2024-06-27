package slice

func IsContains[T comparable](element T, slice []T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func FilterSlice[T comparable](slice []T) []T {
	filteredMap := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		filteredMap[v] = struct{}{}
	}

	filteredSlice := make([]T, len(filteredMap))
	i := 0
	for v := range filteredMap {
		filteredSlice[i] = v
		i++
	}
	return filteredSlice
}
