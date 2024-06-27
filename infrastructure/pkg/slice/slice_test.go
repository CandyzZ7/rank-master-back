package slice

import "testing"

func TestIsContains(t *testing.T) {
	strList := []string{"A", "b"}
	if IsContains("A", strList) {
		t.Logf("should be true")
	}
	if IsContains("B", strList) {
		t.Error("should be false")
	}
}

func TestFilterSlice(t *testing.T) {
	strList := []string{"A", "a", "b", "a", "b"}
	filterSlice := FilterSlice(strList)
	t.Logf("filterSlice: %v", filterSlice)
}
