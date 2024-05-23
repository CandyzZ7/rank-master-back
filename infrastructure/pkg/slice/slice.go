package slice

func IsInSlice(str string, strList []string) bool {
	for _, s := range strList {
		if s == str {
			return true
		}
	}
	return false
}
