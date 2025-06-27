package basics

import (
	"errors"
	"fmt"
)

func main() {

	greet := func() {
		fmt.Println("Hello new function in GO!")
	}

	greet()

	result := addtwo(10, 20, add)

	fmt.Println(result)

	fmt.Println(" new function which return function", multiplier(5)(2))

	resut, error := addition(10, 10)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(resut)
	}

}

func add(a, b int) int {
	return a + b
}

func addtwo(a int, b int, operation func(int, int) int) int {
	return operation(a, b) // No change needed, this line is correct.
}

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func addition(a, b int) (string, error) {
	if a > b {
		return "a is greater thaan b", nil
	} else if a < b {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Can't determine which is greate")
	}
}
