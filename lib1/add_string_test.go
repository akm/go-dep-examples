package lib1

import "testing"

func TestAddString(t *testing.T) {
	AddStringValidCase(t, "1", "2", "3")
	AddStringInvalidCase(t, "1", "a")
	AddStringInvalidCase(t, "1", "")
	AddStringInvalidCase(t, "", "1")
}
