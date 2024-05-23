package slice

import "testing"

func TestIsInSlice(t *testing.T) {
	strList := []string{"A", "b"}
	if IsInSlice("A", strList) {
		t.Logf("should be true")
	}
	if IsInSlice("B", strList) {
		t.Error("should be false")
	}
}
