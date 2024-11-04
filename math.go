package main

import "errors"

func multiplication(a int, b int) int {
	return a * b
}
func division(a float64, b float64) (float64, error) {
	if b == 0 {
		err := errors.New("division par 0 impossible")
		return 0, err
	}
	return (a / b), nil

}
