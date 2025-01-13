package lib1

import "testing"

func AddStringInvalidCase(t *testing.T, a, b string) {
	ret, err := AddString(a, b)
	if err == nil {
		t.Error("AddString should return an error")
	}
	if ret != "" {
		t.Error("AddString should return an empty string")
	}
}
