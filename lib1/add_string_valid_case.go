package lib1

import "testing"

func AddStringValidCase(t *testing.T, a, b, expected string) {
	ret, err := AddString(a, b)
	if err != nil {
		t.Error(err)
	}
	if ret != expected {
		t.Errorf("%s + %s != %s", a, b, expected)
	}
}

func init() {
	println("lib1/add_string_test_helper.go init")
}
