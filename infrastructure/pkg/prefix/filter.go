package prefix

import "sort"

// FilterLongestPrefix 过滤最长前缀
// Example:
// Code:
// path := []string{"image/1/", "image/1/2/"}
// filterLongestPrefix([]string{"image/1/", "image/1/22/"})
// Output: {  "image/1/2/" }
func FilterLongestPrefix(paths []string) []string {
	// 排序以确保相同前缀的路径在一起
	sort.Strings(paths)

	// 存储结果的切片
	filteredPaths := make([]string, 0)

	// 遍历所有路径
	for i := 0; i < len(paths); i++ {
		// 如果已经是最后一个路径或者当前路径和下一个路径不具有相同的前缀
		if i == len(paths)-1 || !hasCommonPrefix(paths[i], paths[i+1]) {
			filteredPaths = append(filteredPaths, paths[i])
		}
	}

	return filteredPaths
}

// 判断两个路径是否具有相同的前缀
func hasCommonPrefix(a, b string) bool {
	minLen := minInt(len(a), len(b))
	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 返回两个数中的最小值
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
