package functions

import "fmt"

func SimpleFunction() string {
	return "This is a simple function from the functions package."
}

func SayHello(name string) string {
	return "Hello, " + name + "!"
}

func IsGoodNumber(num int) (bool, error) {
	if num > 0 {
		return true, nil
	} else {
		return false, fmt.Errorf("number must be greater than zero, got %d", num)
	}
}
