package multiply

import "fmt"

func Multiply(aStr, bStr string) (string, error) {
	// Convert string to float64, then multiply
	a, b := 0.0, 0.0
	if _, err := fmt.Sscanf(aStr, "%f", &a); err != nil {
		return "", err
	}
	if _, err := fmt.Sscanf(bStr, "%f", &b); err != nil {
		return "", err
	}
	res := a * b
	return fmt.Sprintf("%f", res), nil
}

func init() {
	fmt.Println("app1/logics/multiply/multiply_string.go init")
}
