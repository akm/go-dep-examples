package multiply

import (
	"fmt"
	"testing"
)

func TestMultiplyString(t *testing.T) {
	for _, p := range Patterns {
		t.Run(p.A+"*"+p.B, func(t *testing.T) {
			actual, err := Multiply(p.A, p.B)
			if err != nil {
				t.Errorf("Multiply(%s, %s) got error: %v", p.A, p.B, err)
			}
			if actual != p.Expected {
				t.Errorf("Multiply(%s, %s) = %s; want %s", p.A, p.B, actual, p.Expected)
			}
		})
	}
}

func init() {
	fmt.Println("app1/logics/multiply/multiply_string_test.go init")
}
