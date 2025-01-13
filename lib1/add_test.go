package lib1

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 != 3")
	}
}

func init() {
	println("lib1/add_test.go init")
}
